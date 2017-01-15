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

type FunctionObject func([]Object) (Object, error)

func (o FunctionObject) BoolValue() bool    { return o == nil }
func (o FunctionObject) Value() interface{} { return o }
func (fn FunctionObject) Call(args []Object) (Object, error) {
	return fn(args)
}

type StringObject string

func (o StringObject) BoolValue() bool    { return true }
func (o StringObject) Value() interface{} { return string(o) }

// FIXME need to transparently support floating point
type NumberObject int

func (o NumberObject) BoolValue() bool    { return true }
func (o NumberObject) Value() interface{} { return int(o) }

type BoolObject bool

func (o BoolObject) BoolValue() bool    { return bool(o) }
func (o BoolObject) Value() interface{} { return bool(o) }
