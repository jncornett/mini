package mini

type Object interface {
	BoolValue() bool
	Value() interface{}
}

type Expression interface {
	// FIXME maybe replace interface{} with an Object type
	Eval(*Vm) (Object, error) // walk the children
}

type TreeExpr struct {
	Children []Expression
}

func (e *TreeExpr) Eval(vm *Vm) (obj Object, err error) {
	for _, expr := range e.Children {
		obj, err = expr.Eval(vm)
		if err != nil {
			break
		}
	}
	return
}

type IfExpr struct {
	Condition Expression
	IfBlock   Expression
	ElseBlock Expression
}

func (e *IfExpr) Eval(vm *Vm) (Object, error) {
	obj, err := e.Condition.Eval(vm)
	if err != nil {
		return nil, err
	}
	if GetBoolValue(obj) {
		return e.IfBlock.Eval(vm)
	}
	return e.ElseBlock.Eval(vm)
}

type ForExpr struct {
	Condition Expression
	Block     Expression
}

func (e *ForExpr) Eval(vm *Vm) (Object, error) {
	var (
		obj Object
		err error
	)
	for {
		var cond Object
		cond, err = e.Condition.Eval(vm)
		if err != nil {
			break
		}
		if !GetBoolValue(cond) {
			break
		}
		obj, err = e.Block.Eval(vm)
		if err != nil {
			break
		}
	}
	return obj, err
}

type AssignExpr struct {
	LHSSymbol string
	RHS       Expression
}

func (e *AssignExpr) Eval(vm *Vm) (obj Object, err error) {
	obj, err = e.RHS.Eval(vm)
	if err != nil {
		vm.Assign(e.LHSSymbol, obj)
	}
	return
}

type CallExpr struct {
	Symbol string
	Args   []Expression
}

func (e *CallExpr) Eval(vm *Vm) (Object, error) {
	args := make([]Object, len(e.Args))
	for i, expr := range e.Args {
		var err error
		args[i], err = expr.Eval(vm)
		if err != nil {
			return nil, err
		}
	}
	return vm.Call(e.Symbol, args)
}

type ConstExpr struct {
	Value Object
}

func (e *ConstExpr) Eval(vm *Vm) (Object, error) {
	return e.Value, nil
}

type LookupExpr struct {
	Symbol string
}

func (e *LookupExpr) Eval(vm *Vm) (Object, error) {
	return vm.Lookup(e.Symbol), nil
}
