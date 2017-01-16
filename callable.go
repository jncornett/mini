package mini

type Callable interface {
	Call(Args) (Object, error)
}
