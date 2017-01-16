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

func (p *Parser) Parse() (Expression, error) {
	return p.parseExpressionBlock(false)
}

func (p *Parser) scanToken() Token {
	if p.haveLast {
		p.haveLast = false
		return p.last
	}

	p.last = p.s.Scan()
	return p.last
}

func (p *Parser) unscanToken() {
	p.haveLast = true
}

func (p *Parser) scanIgnoreWhitespace() Token {
	tok := p.scanToken()
	if tok.Type == WS {
		return p.scanToken()
	}
	return tok
}

func (p *Parser) accept(tt TokenType) bool {
	tok := p.scanIgnoreWhitespace()
	if tok.Type == tt {
		return true
	}
	p.unscanToken()
	return false
}

func (p *Parser) parseExpression(expect bool) (Expression, error) {
	tok := p.scanIgnoreWhitespace()
	var (
		expr Expression
		err  error
	)
	switch tok.Type {
	case STRING:
		expr = NewStringFromString(tok.Value)
	case NUMBER:
		expr, err = convertTokenToNumber(tok)
	case BOOL:
		expr, err = convertTokenToBool(tok)
	case IDENT:
		if p.accept(ROUNDOPEN) {
			expr, err = p.parseFunctionCall(tok.Value)
		} else if p.accept(ASSIGN) {
			expr, err = p.parseAssignment(tok.Value)
		} else {
			expr = Symbol(tok.Value)
		}
	case NOT:
		expr, err = p.parseNotExpression()
	case SUBTRACT:
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
	case AND:
		expr, err = p.parseAndExpression(expr)
	case OR:
		expr, err = p.parseOrExpression(expr)
	case ADD, SUBTRACT, MULTIPLY, DIVIDE, LESS, LESSEQUAL, GREATER, GREATEREQUAL, EQUAL, NOTEQUAL:
		expr, err = p.parseBinaryExpression(expr, next.Type)
	default:
		p.unscanToken()
	}
	if expr == nil && expect {
		err = fmt.Errorf("Expected expression")
	}
	return expr, err
}

func (p *Parser) parseAndExpression(lhs Expression) (Expression, error) {
	rhs, err := p.parseExpression(true)
	if err != nil {
		return nil, err
	}
	return &AndExpr{LHS: lhs, RHS: rhs}, nil
}

func (p *Parser) parseOrExpression(lhs Expression) (Expression, error) {
	rhs, err := p.parseExpression(true)
	if err != nil {
		return nil, err
	}
	return &OrExpr{LHS: lhs, RHS: rhs}, nil
}

func (p *Parser) parseNotExpression() (Expression, error) {
	expr, err := p.parseExpression(true)
	if err != nil {
		return nil, err
	}
	return &NotExpr{Expr: expr}, nil
}

func (p *Parser) parseFunctionCall(sym string) (Expression, error) {
	args, err := p.parseExpressionList()
	if err != nil {
		return nil, err
	}
	return &CallExpr{Name: Symbol(sym), Args: args}, nil
}

func (p *Parser) parseAssignment(sym string) (Expression, error) {
	rhs, err := p.parseExpression(true)
	if err != nil {
		return nil, err
	}
	return &AssignExpr{Name: Symbol(sym), Expr: rhs}, nil
}

func (p *Parser) parseUnaryExpression(tt TokenType) (Expression, error) {
	expr, err := p.parseExpression(true)
	if err != nil {
		return nil, err
	}
	return &OpExpr{Base: expr, Op: getUnaryOp(tt)}, nil
}

func (p *Parser) parseBinaryExpression(lhs Expression, tt TokenType) (Expression, error) {
	rhs, err := p.parseExpression(true)
	if err != nil {
		return nil, err
	}
	return &OpExpr{Base: lhs, Args: []Expression{rhs}, Op: getBinaryOp(tt)}, nil
}

func (p *Parser) parseParenthesizedExpression() (Expression, error) {
	expressions, err := p.parseExpressionList()
	if err != nil {
		return nil, err
	}
	return &Tree{Children: expressions}, nil
}

func (p *Parser) parseForExpression() (Expression, error) {
	cb, err := p.parseConditional()
	if err != nil {
		return nil, err
	}
	return &ForExpr{For: cb}, nil
}

func (p *Parser) parseIfExpression() (Expression, error) {
	ifExpr := IfExpr{}
	cb, err := p.parseConditional()
	if err != nil {
		return nil, err
	}
	ifExpr.If = cb
	if p.accept(ELSE) {
		cb, err := p.parseConditional()
		if err != nil {
			return nil, err
		}
		ifExpr.Else = cb
	}
	return &ifExpr, nil
}

func (p *Parser) parseConditional() (ConditionalBlock, error) {
	cb := ConditionalBlock{}
	if p.accept(CURLYOPEN) {
		// skip the condition block
		cb.Condition = TRUE
	} else {
		cond, err := p.parseExpression(true)
		if err != nil {
			return cb, err
		}
		cb.Condition = cond
	}
	if !p.accept(CURLYOPEN) {
		return cb, errors.New("Expected block") // FIXME position info error
	}
	block, err := p.parseExpressionBlock(true)
	if err != nil {
		return cb, err
	}
	cb.Block = block
	return cb, nil
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
	return &Tree{Children: expressions}, nil
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

func getUnaryOp(tt TokenType) Op {
	switch tt {
	case SUBTRACT:
		return OpNeg
	}
	return OpNoop
}

func getBinaryOp(tt TokenType) Op {
	switch tt {
	case ADD:
		return OpAdd
	case SUBTRACT:
		return OpSub
	case MULTIPLY:
		return OpMul
	case DIVIDE:
		return OpDiv
	case LESS:
		return OpLt
	case LESSEQUAL:
		return OpLe
	case GREATER:
		return OpGt
	case GREATEREQUAL:
		return OpGe
	case EQUAL:
		return OpEq
	case NOTEQUAL:
		return OpNe
	}
	return OpNoop
}

// FIXME should catch parse errors at lexing
func convertTokenToNumber(t Token) (Number, error) {
	val, err := strconv.ParseFloat(t.Value, 64)
	if err != nil {
		return 0, fmt.Errorf("Expected a number at %v: %v", t.Start, err)
	}
	return NewNumberFromFloat(val), nil
}

func convertTokenToBool(t Token) (Bool, error) {
	val, err := strconv.ParseBool(t.Value)
	if err != nil {
		return false, fmt.Errorf("Expected a bool at %v: %v", t.Start, err)
	}
	return NewBoolFromBool(val), nil
}
