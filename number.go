package mini

import "math"

// Number is a number type. There is not int data type.
type Number float64

// Truthy helps Number implement the Object interface
func (o Number) Truthy() bool { return true }

// IsNil helps Number implement the Object interface
func (o Number) IsNil() bool { return false }

// Send helps Number implement the Object interface
func (o Number) Send(op Op, args Args) (Object, error) {
	switch op {
	case OpNeg:
		return -o, nil
	case OpAdd, OpSub, OpMul, OpDiv, OpLt, OpLe, OpGt, OpGe, OpEq, OpNe:
		rhs, ok := args.Arg(0).(Number)
		if !ok {
			return nil, newErrTypeBadRhs(op, o, rhs, numberType)
		}
		switch op {
		case OpAdd:
			return o + rhs, nil
		case OpSub:
			return o - rhs, nil
		case OpMul:
			return o * rhs, nil
		case OpDiv:
			if rhs == 0 {
				return nil, ErrZeroDivision
			}
			return o / rhs, nil
		case OpLt:
			return Bool(o < rhs), nil
		case OpLe:
			return Bool(o <= rhs), nil
		case OpGt:
			return Bool(o > rhs), nil
		case OpGe:
			return Bool(o >= rhs), nil
		}
	}
	return nil, NewErrInvalidOp(op, o)
}

// Eval helps Number implement the Expression interface
func (o Number) Eval(*Vm) (Object, error) { return o, nil }

// ToInt converts Number to an int
func (o Number) ToInt() int { return int(o) }

// ToFloat converts Number to a float64
func (o Number) ToFloat() float64 { return float64(o) }

// IsFloat returns true if there is a fractional component to o
func (o Number) IsFloat() bool {
	return float64(o) == math.Floor(float64(o))
}

// NewNumberFromFloat constructs a Number from a float64
func NewNumberFromFloat(v float64) Number { return Number(v) }
