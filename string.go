package mini

// String is a string type
type String string

// Truthy helps String implement the Object interface
func (o String) Truthy() bool { return true }

// IsNil helps String implement the Object interface
func (o String) IsNil() bool { return false }

// Send helps String implement the Object interface
func (o String) Send(Op, Args) Object { return nil }

// Eval helps String implement the Expression interface
func (o String) Eval(*Vm) (Object, error) { return o, nil }

// NewStringFromString constructs a String from a v
func NewStringFromString(v string) String { return String(v) }
