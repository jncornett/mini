package mini

// Op is the index of a builtin operation to invoke on an Object
type Op int

// Object represents a generic data type
type Object interface {
	Truthy() bool
	IsNil() bool
	Send(m Op, args Args) (Object, error)
}
