package mini

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	boolType   reflect.Type = reflect.TypeOf(Bool(false))
	funcType   reflect.Type = reflect.TypeOf(Function(nil))
	numberType reflect.Type = reflect.TypeOf(Number(0))
	stringType reflect.Type = reflect.TypeOf(String(""))
)

var (
	ErrZeroDivision = errors.New("Divide by zero")
)

func NewErrInvalidOp(op Op, obj Object) error {
	return fmt.Errorf("InvalidOp: %v is invalid for object of type %T", op, obj)
}

func newErrTypeBadRhs(op Op, lhs, rhs interface{}, rhsType reflect.Type) error {
	return fmt.Errorf(
		"TypeError: expected rhs of %T %v %T to be %v",
		lhs,
		op,
		rhs,
		rhsType,
	)
}

func newTypeError(arg int, want reflect.Type, have interface{}) error {
	return nil // FIXME implement
}
