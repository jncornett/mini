package mini

const (
	// TRUE is a truthy Bool singleton
	TRUE = Bool(true)
	// FALSE is a falsy Bool singleton
	FALSE = Bool(false)
)

// Bool is the bool type
type Bool bool

// Truthy helps Bool implement the Object interface
func (o Bool) Truthy() bool { return bool(o) }

// IsNil helps Bool implement the Object interface
func (o Bool) IsNil() bool { return false }

// Send helps Bool implement the Object interface
func (o Bool) Send(Op, Args) (Object, error) { return nil, nil }

// Eval helps Bool implement the Expression interface
func (o Bool) Eval(*Vm) (Object, error) { return o, nil }

// NewBoolFromBool constructs a Bool from a bool
func NewBoolFromBool(v bool) Bool { return Bool(v) }
