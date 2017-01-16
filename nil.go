package mini

// NIL is a nil singleton
const NIL = Nil(false)

// Nil is the type for the NIL singleton.
type Nil bool

// Truthy helps Nil implement the Object interface.
func (o Nil) Truthy() bool { return false }

// IsNil helps Nil implement the Object interface.
func (o Nil) IsNil() bool { return true }

// Send helps Nil implement the Object interface.
func (o Nil) Send(Op, Args) Object { return nil }

// Eval helps Nil implement the Expression interface.
func (o Nil) Eval(*Vm) (Object, error) { return o, nil }
