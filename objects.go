package mini

import "math"

const (
	NIL = Nil(false)
)

type Method string

type Object interface {
	Truthy() bool
	IsNil() bool
	Send(m Method, args Args) Object
}

type Nil bool

func (o Nil) Truthy() bool             { return false }
func (o Nil) IsNil() bool              { return true }
func (o Nil) Send(Method, Args) Object { return nil }
func (o Nil) Eval(*Vm) (Object, error) { return o, nil }

func NewNil() Nil { return NIL }

type Args []Object

func (o Args) Truthy() bool             { return !o.Empty() }
func (o Args) IsNil() bool              { return false }
func (o Args) Send(Method, Args) Object { return nil }
func (o Args) Eval(*Vm) (Object, error) { return o, nil }
func (o Args) Empty() bool              { return o.Len() == 0 }
func (o Args) Len() int                 { return len(o) }
func (o Args) Arg(n int) Object {
	if n < 0 || o.Len() <= n {
		return NIL
	}
	return o[n]
}

type Function func(Args) (Object, error)

func (o Function) Truthy() bool                   { return o != nil }
func (o Function) IsNil() bool                    { return false }
func (o Function) Send(Method, Args) Object       { return nil }
func (o Function) Call(args Args) (Object, error) { return o(args) }
func (o Function) Eval(*Vm) (Object, error)       { return o, nil }

type String string

func (o String) Truthy() bool             { return true }
func (o String) IsNil() bool              { return false }
func (o String) Send(Method, Args) Object { return nil }
func (o String) Eval(*Vm) (Object, error) { return o, nil }

func NewString() String                   { return String("") }
func NewStringFromString(v string) String { return String(v) }

type Number float64

func (o Number) Truthy() bool             { return true }
func (o Number) IsNil() bool              { return false }
func (o Number) Send(Method, Args) Object { return nil }
func (o Number) Eval(*Vm) (Object, error) { return o, nil }
func (o Number) ToInt() int               { return int(o) }
func (o Number) ToFloat() float64         { return float64(o) }
func (o Number) IsFloat() bool {
	return float64(o) == math.Floor(float64(o))
}

func NewNumber() Number                   { return Number(0) }
func NewNumberFromFloat(v float64) Number { return Number(v) }
func NewNumberFromInt(v int) Number       { return Number(v) }

type Bool bool

func (o Bool) Truthy() bool             { return bool(o) }
func (o Bool) IsNil() bool              { return false }
func (o Bool) Send(Method, Args) Object { return nil }
func (o Bool) Eval(*Vm) (Object, error) { return o, nil }

func NewBool() Bool               { return false }
func NewBoolFromBool(v bool) Bool { return Bool(v) }
