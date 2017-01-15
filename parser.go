package mini

import "io"

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

func (p *Parser) Parse() (Expression, error) {
	ast := &TreeExpr{}
	return ast, nil
}
