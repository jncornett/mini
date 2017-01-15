package mini

import (
	"bufio"
	"bytes"
	"io"
)

type TokenType int

type Position struct {
	Row int
	Col int
}

type Token struct {
	Type  TokenType
	Value string
	Start Position
	End   Position
}

const (
	// Special tokens
	ILLEGAL TokenType = iota
	EOF
	WS

	// Literals
	STRING
	IDENT
	NUMBER
	BOOL

	// Operators/Delimiters
	CURLYOPEN
	CURLYCLOSE
	ROUNDOPEN
	ROUNDCLOSE
	SQUAREOPEN
	SQUARECLOSE
	COMMA
	ASSIGN
	ADD
	SUBTRACT
	MULTIPLY
	DIVIDE
	NOT
	LESS
	GREATER
	LESSEQUAL
	GREATEREQUAL
	EQUAL
	NOTEQUAL

	// Keywords
	IF
	ELSE
	FOR
	FUNC
	BREAK
	CONTINUE
	RETURN
	YIELD
	AND
	OR
)

type Scanner struct {
	r       *bufio.Reader
	pos     Position
	lastPos Position
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() Token {
	start := s.pos
	ch := s.readRune()

	if isWhitespace(ch) {
		return s.scanWhitespace(ch, start)
	} else if isLetter(ch) || ch == '_' {
		return s.scanIdent(ch, start)
	} else if isNumber(ch) || ch == '.' {
		return s.scanNumberLiteral(ch, start)
	} else if ch == '"' {
		return s.scanStringLiteral(start)
	} else if ch == EOFCHAR {
		return Token{EOF, "", start, s.pos}
	}

	tt := ILLEGAL
	val := string(ch)
	switch ch {
	case '{':
		tt = CURLYOPEN
	case '}':
		tt = CURLYCLOSE
	case '(':
		tt = ROUNDOPEN
	case ')':
		tt = ROUNDCLOSE
	case '[':
		tt = SQUAREOPEN
	case ']':
		tt = SQUARECLOSE
	case ',':
		tt = COMMA
	case '+':
		tt = ADD
	case '-':
		tt = SUBTRACT
	case '*':
		tt = MULTIPLY
	case '/':
		tt = DIVIDE
	case '=':
		if s.readRune() == '=' {
			val += "="
			tt = EQUAL
		} else {
			s.unreadRune()
			tt = ASSIGN
		}
	case '!':
		if s.readRune() == '=' {
			val += "="
			tt = NOTEQUAL
		} else {
			s.unreadRune()
			tt = NOT
		}
	case '>': // GREATER | GREATEREQUAL
		if s.readRune() == '=' {
			val += "="
			tt = GREATEREQUAL
		} else {
			s.unreadRune()
			tt = GREATER
		}
	case '<': // LESS | LESSEQUAL
		if s.readRune() == '=' {
			val += "="
			tt = LESSEQUAL
		} else {
			s.unreadRune()
			tt = LESS
		}
	}

	return Token{tt, val, start, s.pos}
}

func (s *Scanner) readRune() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		ch = EOFCHAR
	}
	s.lastPos = s.pos
	if ch == '\n' {
		s.pos.Row++
		s.pos.Col = 0
	} else {
		s.pos.Col++
	}
	return ch
}

func (s *Scanner) unreadRune() {
	_ = s.r.UnreadRune()
	s.pos = s.lastPos
	s.lastPos = Position{}
}

func (s *Scanner) scanWhitespace(first rune, start Position) Token {
	var buf bytes.Buffer
	buf.WriteRune(first)
	for {
		ch := s.readRune()
		if isWhitespace(ch) {
			buf.WriteRune(ch)
		} else {
			s.unreadRune()
			break
		}
	}
	return Token{WS, buf.String(), start, s.pos}
}

func (s *Scanner) scanIdent(first rune, start Position) Token {
	var buf bytes.Buffer
	buf.WriteRune(first)
	for {
		ch := s.readRune()
		if isLetter(ch) || isNumber(ch) || ch == '_' {
			buf.WriteRune(ch)
		} else {
			s.unreadRune()
			break
		}
	}
	val := buf.String()
	var tok TokenType
	// Check to see if we have a keyword
	switch val {
	case "if":
		tok = IF
	case "else":
		tok = ELSE
	case "for":
		tok = FOR
	case "func":
		tok = FUNC
	case "break":
		tok = BREAK
	case "continue":
		tok = CONTINUE
	case "return":
		tok = RETURN
	case "yield":
		tok = YIELD
	case "and":
		tok = AND
	case "or":
		tok = OR
	case "true", "false":
		tok = BOOL
	default:
		tok = IDENT
	}
	return Token{tok, val, start, s.pos}
}

func (s *Scanner) scanStringLiteral(start Position) Token {
	var buf bytes.Buffer
	escape := false
	for {
		ch := s.readRune()
		if ch == EOFCHAR {
			break
		} else if escape {
			buf.WriteRune(ch)
			escape = false
			continue
		} else if ch == '\\' {
			escape = true
		} else if ch == '"' {
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return Token{STRING, buf.String(), start, s.pos}
}

func (s *Scanner) scanNumberLiteral(first rune, start Position) Token {
	var buf bytes.Buffer
	buf.WriteRune(first)
	for {
		ch := s.readRune()
		if isNumber(ch) || ch == '.' {
			buf.WriteRune(ch)
		} else {
			s.unreadRune()
			break
		}
	}
	return Token{NUMBER, buf.String(), start, s.pos}
}

const EOFCHAR = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isNumber(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
