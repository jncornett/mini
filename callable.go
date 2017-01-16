package mini

type Callable interface {
	Call(ArgsObject) (Object, error)
}
