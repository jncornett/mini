package mini

import "strings"

// String is a string type
type String string

// Truthy helps String implement the Object interface
func (o String) Truthy() bool { return true }

// IsNil helps String implement the Object interface
func (o String) IsNil() bool { return false }

// Send helps String implement the Object interface
func (o String) Send(op Op, args Args) (Object, error) {
	switch op {
	case OpMul:
		rhs, ok := args.Arg(0).(Number)
		if !ok {
			return nil, newErrTypeBadRhs(op, o, rhs, stringType)
		}
	case OpAdd:
		rhs, ok := args.Arg(0).(String)
		if !ok {
			return nil, newErrTypeBadRhs(op, o, rhs, stringType)
		}
		return o + rhs, nil
	case OpEq, OpNe, OpLt, OpLe, OpGt, OpGe:
		rhs, ok := args.Arg(0).(String)
		if !ok {
			return nil, newErrTypeBadRhs(op, o, rhs, stringType)
		}
		cmp := strings.Compare(string(o), string(rhs))
		switch op {
		case OpEq:
			return Bool(cmp == 0), nil
		case OpNe:
			return Bool(cmp != 0), nil
		case OpLt:
			return Bool(cmp < 0), nil
		case OpLe:
			return Bool(cmp <= 0), nil
		case OpGt:
			return Bool(cmp > 0), nil
		case OpGe:
			return Bool(cmp >= 0), nil
		}
	}
	return nil, NewErrInvalidOp(op, o)
}

// Eval helps String implement the Expression interface
func (o String) Eval(*Vm) (Object, error) { return o, nil }

// NewStringFromString constructs a String from a v
func NewStringFromString(v string) String { return String(v) }
