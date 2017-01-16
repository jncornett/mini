package mini

type Expression interface {
	Eval(*Vm) (Object, error)
}

type Tree struct {
	Children []Expression
}

func (e *Tree) Eval(vm *Vm) (obj Object, err error) {
	for _, expr := range e.Children {
		obj, err = expr.Eval(vm)
		if err != nil {
			break
		}
	}
	return
}

type ConditionalBlock struct {
	Condition Expression
	Block     Expression
}

type IfExpr struct {
	If   ConditionalBlock
	Else ConditionalBlock
}

func (e *IfExpr) Eval(vm *Vm) (Object, error) {
	didEval, obj, err := evalBranch(e.If, vm)
	if !didEval || err != nil {
		return obj, err
	}
	_, obj, err = evalBranch(e.Else, vm)
	return obj, err
}

type ForExpr struct {
	For ConditionalBlock
}

func (e *ForExpr) Eval(vm *Vm) (Object, error) {
	var (
		obj Object
		err error
	)
	for {
		var didEval bool
		didEval, obj, err = evalBranch(e.For, vm)
		if err != nil || !didEval {
			break
		}
	}
	return obj, err
}

type AssignExpr struct {
	Name Symbol
	Expr Expression
}

func (e *AssignExpr) Eval(vm *Vm) (obj Object, err error) {
	obj, err = e.Expr.Eval(vm)
	if err == nil {
		vm.Assign(e.Name, obj)
	}
	return
}

type CallExpr struct {
	Name Symbol
	Args []Expression
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
	return vm.Call(e.Name, args)
}

type Symbol string

func (e Symbol) Eval(vm *Vm) (Object, error) {
	return vm.Lookup(e), nil
}

func evalBranch(cb ConditionalBlock, vm *Vm) (bool, Object, error) {
	if cb.Condition == nil {
		return false, NIL, nil
	}
	obj, err := cb.Condition.Eval(vm)
	if err != nil || !obj.Truthy() {
		return false, NIL, err
	}
	if cb.Block == nil {
		return false, NIL, nil
	}
	obj, err = cb.Block.Eval(vm)
	return true, obj, err
}
