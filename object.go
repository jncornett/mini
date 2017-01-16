package mini

// Object represents a generic data type
type Object interface {
	Truthy() bool
	IsNil() bool
	Send(m Op, args Args) (Object, error)
}
