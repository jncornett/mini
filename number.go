package mini

import "math"

// Number is a number type. There is not int data type.
type Number float64

// Truthy helps Number implement the Object interface
func (o Number) Truthy() bool { return true }

// IsNil helps Number implement the Object interface
func (o Number) IsNil() bool { return false }

// Send helps Number implement the Object interface
func (o Number) Send(Op, Args) (Object, error) { return nil, nil }

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
