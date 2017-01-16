package mini

import (
	"fmt"
	"strings"
)

func (e Tree) String() string {
	var children []string
	for _, child := range e.Children {
		children = append(children, fmt.Sprint(child))
	}
	return fmt.Sprint("Tree[", strings.Join(children, " "), "]")
}

func (cb ConditionalBlock) String() string {
	return fmt.Sprint("Cond(", cb.Condition, "=>", cb.Block, ")")
}

func (e IfExpr) String() string {
	return fmt.Sprint("If(", e.If, " ", e.Else, ")")
}

func (e ForExpr) String() string {
	return fmt.Sprint("For(", e.For, ")")
}

func (e AssignExpr) String() string {
	return fmt.Sprint(e.Name, "=", e.Expr)
}

func (e CallExpr) String() string {
	return fmt.Sprint(e.Name, e.Args)
}

func (e Symbol) String() string {
	return fmt.Sprint("@", string(e))
}
