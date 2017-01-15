package mini

func GetBoolValue(obj Object) bool {
	if obj == nil {
		return false
	}
	return obj.BoolValue()
}

type Callable interface {
	Call([]Object) (Object, error)
}

type Function func([]Object) (Object, error)

func (fn Function) Call(args []Object) (Object, error) {
	return fn(args)
}

func (fn Function) BoolValue() bool    { return fn == nil }
func (fn Function) Value() interface{} { return fn }
