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
	obj := vm.Lookup(e)
	if obj == nil {
		return NIL, nil
	}
	return obj, nil
}

type NotExpr struct {
	Expr Expression
}

func (e NotExpr) Eval(vm *Vm) (Object, error) {
	if e.Expr == nil {
		return NIL, nil
	}
	obj, err := e.Expr.Eval(vm)
	if err != nil {
		return nil, err
	}
	return Bool(!obj.Truthy()), nil
}

// FIXME rename LHS => Lhs
// FIXME rename RHS => Rhs
type AndExpr struct {
	LHS Expression
	RHS Expression
}

func (e AndExpr) Eval(vm *Vm) (Object, error) {
	// preconditions
	if e.LHS == nil || e.RHS == nil {
		return NIL, nil
	}
	obj, err := e.LHS.Eval(vm)
	if err != nil {
		return nil, err
	}
	if !obj.Truthy() {
		// short circuit if possible
		return obj, nil
	}
	obj, err = e.RHS.Eval(vm)
	if err != nil {
		return nil, err
	}
	if !obj.Truthy() {
		return FALSE, nil
	}
	return obj, nil
}

// FIXME rename LHS => Lhs
// FIXME rename RHS => Rhs
type OrExpr struct {
	LHS Expression
	RHS Expression
}

func (e OrExpr) Eval(vm *Vm) (Object, error) {
	// preconditions
	if e.LHS == nil || e.RHS == nil {
		return NIL, nil
	}
	lhs, err := e.LHS.Eval(vm)
	if err != nil {
		return nil, err
	}
	if lhs.Truthy() {
		// short circuit if possible
		return lhs, nil
	}
	rhs, err := e.RHS.Eval(vm)
	if err != nil {
		return nil, err
	}
	if rhs.Truthy() {
		return rhs, nil
	}
	return FALSE, nil
}

type OpExpr struct {
	Base Expression
	Args []Expression
	Op   Op
}

func (e OpExpr) Eval(vm *Vm) (Object, error) {
	lhs, err := e.Base.Eval(vm)
	if err != nil {
		return nil, err
	}
	var args Args
	for _, expr := range e.Args {
		obj, err := expr.Eval(vm)
		if err != nil {
			return nil, err
		}
		args = append(args, obj) // FIXME implement args.Append or args.Push?
	}
	ret, err := lhs.Send(e.Op, args)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return nil, NewErrInvalidOp(e.Op, lhs)
	}
	return ret, nil
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
