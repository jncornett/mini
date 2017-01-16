package mini

// Function is a go function
type Function func(Args) (Object, error)

// Truthy helps Function implement the Object interface
func (o Function) Truthy() bool { return o != nil }

// IsNil helps Function implement the Object interface
func (o Function) IsNil() bool { return false }

// Send helps Function implement the Object interface
func (o Function) Send(Op, Args) (Object, error) { return nil, nil }

// Call helps Function implement the Callable interface
func (o Function) Call(args Args) (Object, error) { return o(args) }

// Eval helps Function implement the Expression interface
func (o Function) Eval(*Vm) (Object, error) { return o, nil }
