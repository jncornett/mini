package mini_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/jncornett/mini"
)

func TestScannerScan(t *testing.T) {
	tests := []struct {
		Program string
		Type    mini.TokenType
		Value   string
	}{
		{"", mini.EOF, ""},
		{"  ", mini.WS, "  "},
		{"\"", mini.STRING, ""},
		{"\"\\\"\"", mini.STRING, "\""},
		{"\"foo\"", mini.STRING, "foo"},
		{"f", mini.IDENT, "f"},
		{"foo", mini.IDENT, "foo"},
		{"f123_oo", mini.IDENT, "f123_oo"},
		{"_foo", mini.IDENT, "_foo"},
		{
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_",
			mini.IDENT,
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_",
		},
		{"123", mini.NUMBER, "123"},
		{".123", mini.NUMBER, ".123"},
		{".", mini.NUMBER, "."},
		{"123.", mini.NUMBER, "123."},
		{"123.456", mini.NUMBER, "123.456"},
		{"...456...", mini.NUMBER, "...456..."},
		{"true", mini.BOOL, "true"},
		{"false", mini.BOOL, "false"},
		{"if", mini.IF, "if"},
		{"else", mini.ELSE, "else"},
		{"for", mini.FOR, "for"},
		{"break", mini.BREAK, "break"},
		{"continue", mini.CONTINUE, "continue"},
		{"and", mini.AND, "and"},
		{"or", mini.OR, "or"},
		{"{", mini.CURLYOPEN, "{"},
		{"}", mini.CURLYCLOSE, "}"},
		{"(", mini.ROUNDOPEN, "("},
		{")", mini.ROUNDCLOSE, ")"},
		{"[", mini.SQUAREOPEN, "["},
		{"]", mini.SQUARECLOSE, "]"},
		{",", mini.COMMA, ","},
		{"=", mini.ASSIGN, "="},
		{"+", mini.ADD, "+"},
		{"-", mini.SUBTRACT, "-"},
		{"*", mini.MULTIPLY, "*"},
		{"/", mini.DIVIDE, "/"},
		{"!", mini.NOT, "!"},
		{"<", mini.LESS, "<"},
		{">", mini.GREATER, ">"},
		{"<=", mini.LESSEQUAL, "<="},
		{">=", mini.GREATEREQUAL, ">="},
		{"==", mini.EQUAL, "=="},
		{"!=", mini.NOTEQUAL, "!="},
	}

	for _, test := range tests {
		t.Run(test.Program, func(t *testing.T) {
			s := mini.NewScanner(strings.NewReader(test.Program))
			tok := s.Scan()
			if test.Type != tok.Type {
				t.Errorf("expected type %v, got %v", test.Type, tok.Type)
			}
			if test.Value != tok.Value {
				// t.Logf("expected: %v", []byte(test.Value))
				// t.Logf("actual: %v", []byte(test.Value))
				t.Errorf("expected value %q, got %q", test.Value, tok.Value)
			}
		})
	}
}

func TestScannerScanRepeated(t *testing.T) {
	program := "(hello !=\ngoodbye)\n"
	tokens := []mini.Token{
		{mini.ROUNDOPEN, "(", pos(0, 0), pos(0, 1)},
		{mini.IDENT, "hello", pos(0, 1), pos(0, 6)},
		{mini.WS, " ", pos(0, 6), pos(0, 7)},
		{mini.NOTEQUAL, "!=", pos(0, 7), pos(0, 9)},
		{mini.WS, "\n", pos(0, 9), pos(1, 0)},
		{mini.IDENT, "goodbye", pos(1, 0), pos(1, 7)},
		{mini.ROUNDCLOSE, ")", pos(1, 7), pos(1, 8)},
		{mini.WS, "\n", pos(1, 8), pos(2, 0)},
	}
	s := mini.NewScanner(strings.NewReader(program))
	for i, tok := range tokens {
		got := s.Scan()
		if !reflect.DeepEqual(tok, got) {
			t.Errorf("expected token %v to be %v, got %v", i, tok, got)
		}
	}
}

func pos(row, col int) mini.Position {
	return mini.Position{Row: row, Col: col}
}
