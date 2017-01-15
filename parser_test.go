package mini_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jncornett/mini"
)

func TestParserParse(t *testing.T) {
	tests := []struct {
		Program        string
		ExpectError    bool
		ExpectedOutput string
	}{
		{
			"print()",
			false,
			"Tree[@print([])]",
		},
		{
			"print()print()",
			false,
			"Tree[@print([]) @print([])]",
		},
		{
			"print()",
			false,
			"Tree[@print([])]",
		},
		{
			"(,,)",
			false,
			"Tree[Tree[<nil>]]", // FIXME we should optimize out empty branches
		},
		{
			"print(foo)",
			false,
			"Tree[@print([@foo])]",
		},
		{
			"print(foo,)",
			false,
			"Tree[@print([@foo])]",
		},
		{
			"print(foo,bar)",
			false,
			"Tree[@print([@foo @bar])]",
		},
		{
			"print(foo,bar,)",
			false,
			"Tree[@print([@foo @bar])]",
		},
		{
			"print(\"hello\", 123, \"world\", false)",
			false,
			"Tree[@print([#(string)hello #(int)123 #(string)world #(bool)false])]",
		},
		{
			"foo",
			false,
			"Tree[@foo]",
		},
		{
			"foo =",
			true,
			"",
		},
		{
			"foo = bar",
			false,
			"Tree[@foo=(@bar)]",
		},
	}
	for _, test := range tests {
		t.Run(test.Program, func(t *testing.T) {
			ast, err := mini.NewParser(strings.NewReader(test.Program)).Parse()
			if test.ExpectError {
				if err == nil {
					t.Fatal("expected an error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			output := fmt.Sprint(ast)
			if test.ExpectedOutput != output {
				t.Errorf("expected %q, got %q", test.ExpectedOutput, output)
			}
		})
	}
}
