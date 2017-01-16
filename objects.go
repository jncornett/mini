package mini

import "math"

type Method string

type Object interface {
	Truthy() bool
	Nil() bool // FIXME rename IsNil
	Send(m Method, args ArgsObject) Object
}

type Nil struct{}

func (o Nil) Truthy() bool                   { return false }
func (o Nil) Nil() bool                      { return true }
func (o Nil) Send(Method, ArgsObject) Object { return nil }
func (o Nil) Eval(*Vm) (Object, error)       { return o, nil }

type ArgsObject []Object

func (o ArgsObject) Truthy() bool                   { return !o.Empty() }
func (o ArgsObject) Nil() bool                      { return false }
func (o ArgsObject) Send(Method, ArgsObject) Object { return nil }
func (o ArgsObject) Eval(*Vm) (Object, error)       { return o, nil }
func (o ArgsObject) Empty() bool                    { return o.Len() == 0 }
func (o ArgsObject) Len() int                       { return len(o) }
func (o ArgsObject) Arg(n int) Object {
	if n < 0 || o.Len() <= n {
		return &Nil{} // FIXME want singleton?
	}
	return o[n]
}

type FunctionObject func(ArgsObject) (Object, error)

func (o FunctionObject) Truthy() bool                         { return o != nil }
func (o FunctionObject) Nil() bool                            { return false }
func (o FunctionObject) Send(Method, ArgsObject) Object       { return nil }
func (o FunctionObject) Call(args ArgsObject) (Object, error) { return o(args) }
func (o FunctionObject) Eval(*Vm) (Object, error)             { return o, nil }

type StringObject string

func (o StringObject) Truthy() bool                   { return true }
func (o StringObject) Nil() bool                      { return false }
func (o StringObject) Send(Method, ArgsObject) Object { return nil }
func (o StringObject) Eval(*Vm) (Object, error)       { return o, nil }

func StringObjectFromString(v string) StringObject { return StringObject(v) }

type NumberObject float64

func (o NumberObject) Truthy() bool                   { return true }
func (o NumberObject) Nil() bool                      { return false }
func (o NumberObject) Send(Method, ArgsObject) Object { return nil }
func (o NumberObject) Eval(*Vm) (Object, error)       { return o, nil }
func (o NumberObject) ToInt() int                     { return int(o) }
func (o NumberObject) ToFloat() float64               { return float64(o) }
func (o NumberObject) IsFloat() bool {
	return float64(o) == math.Floor(float64(o))
}

func NumberObjectFromFloat(v float64) NumberObject { return NumberObject(v) }
func NumberObjectFromInt(v int) NumberObject       { return NumberObject(v) }

type BoolObject bool

func (o BoolObject) Truthy() bool                   { return bool(o) }
func (o BoolObject) Nil() bool                      { return false }
func (o BoolObject) Send(Method, ArgsObject) Object { return nil }
func (o BoolObject) Eval(*Vm) (Object, error)       { return o, nil }

func BoolObjectFromBool(v bool) BoolObject { return BoolObject(v) }
