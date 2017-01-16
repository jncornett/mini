package mini

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Parser struct {
	s        *Scanner
	last     Token
	haveLast bool
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) scan() Token {
	if p.haveLast {
		p.haveLast = false
		return p.last
	}

	p.last = p.s.Scan()
	return p.last
}

func (p *Parser) unscan() {
	p.haveLast = true
}

func (p *Parser) scanIgnoreWhitespace() Token {
	tok := p.scan()
	if tok.Type == WS {
		return p.scan()
	}
	return tok
}

func (p *Parser) accept(tt TokenType) bool {
	tok := p.scanIgnoreWhitespace()
	if tok.Type == tt {
		return true
	}
	p.unscan()
	return false
}

func (p *Parser) Parse() (Expression, error) {
	return p.parseExpressionBlock(false)
}

func (p *Parser) parseExpression(expect bool) (Expression, error) {
	tok := p.scanIgnoreWhitespace()
	var (
		expr Expression
		err  error
	)
	switch tok.Type {
	case STRING:
		expr = &ConstExpr{Value: String(tok.Value)}
	case NUMBER:
		// FIXME need to transparently support floating point
		v, errConv := strconv.Atoi(tok.Value)
		if errConv != nil {
			err = fmt.Errorf("Expected a number at %v: %v", tok.Start, errConv)
		} else {
			expr = &ConstExpr{Value: Number(v)}
		}
	case BOOL:
		v, errConv := strconv.ParseBool(tok.Value)
		if errConv != nil {
			err = fmt.Errorf("Expected a boolean at %v: %v", tok.Start, errConv)
		} else {
			expr = &ConstExpr{Value: Bool(v)}
		}
	case IDENT:
		if p.accept(ROUNDOPEN) {
			expr, err = p.parseFunctionCall(tok.Value)
		} else if p.accept(ASSIGN) {
			expr, err = p.parseAssignment(tok.Value)
		} else {
			expr = &LookupExpr{Symbol: tok.Value}
		}
	case NOT, SUBTRACT:
		expr, err = p.parseUnaryExpression(tok.Type)
	case ROUNDOPEN:
		expr, err = p.parseParenthesizedExpression()
	case IF:
		expr, err = p.parseIfExpression()
	case FOR:
		expr, err = p.parseForExpression()
	}
	// Short-circuit if we have an error at this point
	if err != nil {
		return nil, err
	}
	// Now we need to lookahead one token to check if this expression is part of a BinExpr
	next := p.scanIgnoreWhitespace()
	switch next.Type {
	case ADD, SUBTRACT, MULTIPLY, DIVIDE, LESS, LESSEQUAL, GREATER, GREATEREQUAL, EQUAL, NOTEQUAL, AND, OR:
		expr, err = p.parseBinaryExpression(expr, next.Type)
	default:
		p.unscan()
	}
	if expr == nil && expect {
		err = fmt.Errorf("Expected expression")
	}
	return expr, err
}

func (p *Parser) parseFunctionCall(sym string) (Expression, error) {
	args, err := p.parseExpressionList()
	if err != nil {
		return nil, err
	}
	return &CallExpr{Symbol: sym, Args: args}, nil
}

func (p *Parser) parseAssignment(sym string) (Expression, error) {
	rhs, err := p.parseExpression(true)
	if err != nil {
		return nil, err
	}
	return &AssignExpr{LHSSymbol: sym, RHS: rhs}, nil
}

func (p *Parser) parseUnaryExpression(tt TokenType) (Expression, error) {
	expr, err := p.parseExpression(true)
	if err != nil {
		return nil, err
	}
	sym, err := getUnaryFunctionName(tt)
	if err != nil {
		return nil, err
	}
	return &CallExpr{Symbol: sym, Args: []Expression{expr}}, nil
}

func (p *Parser) parseBinaryExpression(lhs Expression, tt TokenType) (Expression, error) {
	rhs, err := p.parseExpression(true)
	if err != nil {
		return nil, err
	}
	sym, err := getBinaryFunctionName(tt)
	if err != nil {
		return nil, err
	}
	return &CallExpr{Symbol: sym, Args: []Expression{lhs, rhs}}, nil
}

func (p *Parser) parseParenthesizedExpression() (Expression, error) {
	expressions, err := p.parseExpressionList()
	if err != nil {
		return nil, err
	}
	return &TreeExpr{Children: expressions}, nil
}

func (p *Parser) parseForExpression() (Expression, error) {
	forCond, forBlock, err := p.parseConditional()
	if err != nil {
		return nil, err
	}
	return &ForExpr{Condition: forCond, Block: forBlock}, nil
}

func (p *Parser) parseIfExpression() (Expression, error) {
	ifCond, ifBlock, err := p.parseConditional()
	if err != nil {
		return nil, err
	}
	expr := IfExpr{IfCond: ifCond, IfBlock: ifBlock}
	if p.accept(ELSE) {
		elseCond, elseBlock, err := p.parseConditional()
		if err != nil {
			return nil, err
		}
		expr.ElseCond = elseCond
		expr.ElseBlock = elseBlock
	}
	return &expr, nil
}

func (p *Parser) parseConditional() (cond Expression, block Expression, err error) {
	if p.accept(CURLYOPEN) {
		// skip the condition block
		cond = &ConstExpr{Value: Bool(true)}
	} else {
		cond, err = p.parseExpression(true)
		if err != nil {
			return
		}
		if !p.accept(CURLYOPEN) {
			err = errors.New("Expected block")
		}
	}
	block, err = p.parseExpressionBlock(true)
	return
}

func (p *Parser) parseExpressionBlock(enclosed bool) (Expression, error) {
	var expressions []Expression
	for {
		if enclosed && p.accept(CURLYCLOSE) {
			break
		}
		expr, err := p.parseExpression(false)
		if err != nil {
			return nil, err
		}
		if expr == nil {
			break
		}
		expressions = append(expressions, expr)
	}
	return &TreeExpr{Children: expressions}, nil
}

func (p *Parser) parseExpressionList() ([]Expression, error) {
	var expressions []Expression
	for {
		if p.accept(ROUNDCLOSE) {
			break
		}
		expr, err := p.parseExpression(false)
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, expr)
		p.accept(COMMA)
	}
	return expressions, nil
}

func getUnaryFunctionName(tt TokenType) (string, error) {
	switch tt {
	case NOT:
		return "__not", nil
	case SUBTRACT:
		return "__neg", nil
	}
	return "", fmt.Errorf("No unary function for token: %v", tt)
}

func getBinaryFunctionName(tt TokenType) (string, error) {
	switch tt {
	case ADD:
		return "__add", nil
	case SUBTRACT:
		return "__sub", nil
	case MULTIPLY:
		return "__mul", nil
	case DIVIDE:
		return "__div", nil
	case LESS:
		return "__lt", nil
	case LESSEQUAL:
		return "__le", nil
	case GREATER:
		return "__gt", nil
	case GREATEREQUAL:
		return "__ge", nil
	case EQUAL:
		return "__eq", nil
	case NOTEQUAL:
		return "__ne", nil
	case AND:
		return "__and", nil
	case OR:
		return "__or", nil
	}
	return "", fmt.Errorf("No binary function for token: %v", tt)
}
