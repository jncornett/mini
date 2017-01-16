package mini

import "reflect"

// Function is a go function
type Function func(Args) (Object, error)

// Truthy helps Function implement the Object interface
func (o Function) Truthy() bool { return o != nil }

// IsNil helps Function implement the Object interface
func (o Function) IsNil() bool { return false }

// Send helps Function implement the Object interface
func (o Function) Send(op Op, args Args) (Object, error) {
	switch op {
	case OpEq, OpNe:
		rhs, ok := args.Arg(0).(Function)
		if !ok {
			return nil, newErrTypeBadRhs(op, o, rhs, funcType)
		}
		eq := reflect.DeepEqual(o, rhs)
		if op == OpEq {
			return Bool(eq), nil
		} else {
			return Bool(!eq), nil
		}
	}
	return nil, NewErrInvalidOp(op, o)
}

// Call helps Function implement the Callable interface
func (o Function) Call(args Args) (Object, error) { return o(args) }

// Eval helps Function implement the Expression interface
func (o Function) Eval(*Vm) (Object, error) { return o, nil }
