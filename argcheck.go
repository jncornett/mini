package mini

import (
	"fmt"
	"reflect"
)

func NewErrInvalidOp(op Op, obj Object) error {
	return fmt.Errorf("InvalidOp: %v is invalid for object of type %T", op, obj)
}

func newTypeError(arg int, want reflect.Type, have interface{}) error {
	return nil // FIXME implement
}
