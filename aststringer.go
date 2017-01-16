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

func (e NotExpr) String() string {
	return fmt.Sprint("Not[", e.Expr, "]")
}

func (e AndExpr) String() string {
	return fmt.Sprint("Not[", e.LHS, " ", e.RHS, "]")
}

func (e OrExpr) String() string {
	return fmt.Sprint("Not[", e.LHS, " ", e.RHS, "]")
}

func (e OpExpr) String() string {
	children := []string{fmt.Sprint(e.Base)}
	for _, arg := range e.Args {
		children = append(children, fmt.Sprint(arg))
	}
	return fmt.Sprint("Op{", e.Op, "}[", strings.Join(children, " "), "]")
}
