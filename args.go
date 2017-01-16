package mini

// Args represents an argument list.
type Args []Object

// Truthy helps Args implement the Object interface.
func (o Args) Truthy() bool { return !o.Empty() }

// IsNil helps Args implement the Object interface.
func (o Args) IsNil() bool { return false }

// Send helps Args implement the Object interface.
func (o Args) Send(Op, Args) (Object, error) { return nil, nil }

// Eval helps Args implement the Expression interface.
func (o Args) Eval(*Vm) (Object, error) { return o, nil }

// Empty returns true if the argument list is nil or empty, false otherwise.
func (o Args) Empty() bool { return o.Len() == 0 }

// Len returns the length of the argument list.
func (o Args) Len() int { return len(o) }

// Arg returns the argument at n if it exits, or NIL otherwise
func (o Args) Arg(n int) Object {
	if n < 0 || o.Len() <= n {
		return NIL
	}
	return o[n]
}
