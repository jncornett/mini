package mini

import (
	"fmt"
	"reflect"
)

func (e *TreeExpr) String() string {
	return fmt.Sprintf("Tree%v", e.Children)
}

func (e *LookupExpr) String() string {
	return fmt.Sprint("@", e.Symbol)
}

func (e *CallExpr) String() string {
	return fmt.Sprintf("@%v(%v)", e.Symbol, e.Args)
}

func (e *ConstExpr) String() string {
	val := e.Value.Value()
	return fmt.Sprintf("#(%v)%+v", reflect.TypeOf(val).Name(), val)
}

func (e *AssignExpr) String() string {
	return fmt.Sprintf("@%v=(%v)", e.LHSSymbol, e.RHS)
}
