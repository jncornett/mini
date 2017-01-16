package mini

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	numberType = reflect.TypeOf(NumberObject(0))
)

type StdlibEntry struct {
	Name string
	Func FunctionObject
}

// FIXME this isn't the best way to organize functionality
func GetStdlib() []StdlibEntry {
	return []StdlibEntry{
		// builtins
		{
			"__not",
			func(args ArgsObject) (Object, error) {
				return BoolObjectFromBool(args.Arg(0).Truthy()), nil
			},
		},
		{
			"__neg",
			func(args ArgsObject) (Object, error) {
				num, ok := args.Arg(0).(NumberObject)
				// currently only defined for numbers
				if !ok {
					return nil, typeError(numberType, num)
				}
				switch v := getArg(args, 0).(type) {
				case NumberObject:
					return NumberObject(-int(v)), nil
				default:
					return nil, typeError(numberType, num)
				}
			},
		},
		{
			"__add",
			func(args ArgsObject) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return NumberObject(lhs + rhs), nil
			},
		},
		{
			"__sub",
			func(args ArgsObject) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return NumberObject(lhs - rhs), nil
			},
		},
		{
			"__mul",
			func(args ArgsObject) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return NumberObject(lhs * rhs), nil
			},
		},
		{
			"__div",
			func(args ArgsObject) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				if rhs == 0 {
					return nil, errors.New("ZeroDivisionError")
				}
				return NumberObject(lhs / rhs), nil
			},
		},
		{
			"__lt",
			func(args ArgsObject) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return BoolObject(lhs < rhs), nil
			},
		},
		{
			"__le",
			func(args ArgsObject) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return BoolObject(lhs <= rhs), nil
			},
		},
		{
			"__gt",
			func(args ArgsObject) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return BoolObject(lhs > rhs), nil
			},
		},
		{
			"__ge",
			func(args ArgsObject) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return BoolObject(lhs >= rhs), nil
			},
		},
		{
			"__eq",
			func(args ArgsObject) (Object, error) {
				return BoolObject(reflect.DeepEqual(getArg(args, 0), getArg(args, 1))), nil
			},
		},
		{
			"__ne",
			func(args ArgsObject) (Object, error) {
				return BoolObject(!reflect.DeepEqual(getArg(args, 0), getArg(args, 1))), nil
			},
		},
		{
			"__and",
			func(args ArgsObject) (Object, error) {
				lhs := args.Arg(0)
				if !lhs.Truthy() {
					// short circuit if we can
					return BoolObject(false), nil
				}
				rhs := args.Arg(1)
				if !rhs.Truthy() {
					return BoolObject(false), nil
				}
				return rhs, nil
			},
		},
		{
			"__or",
			func(args ArgsObject) (Object, error) {
				lhs := args.Arg(0)
				if lhs.Truthy() {
					// short circuit if possible
					return lhs, nil
				}
				rhs := args.Arg(1)
				if rhs.Truthy() {
					return rhs, nil
				}
				return BoolObject(false), nil
			},
		},
		// stdlib
		{
			"print",
			func(args ArgsObject) (Object, error) {
				_, err := fmt.Println(objectsToEmpties(args)...)
				return nil, err
			},
		},
	}
}

func objectsToEmpties(args ArgsObject) []interface{} {
	var out []interface{}
	for _, arg := range args {
		out = append(out, arg)
	}
	return out
}

func getArg(args ArgsObject, n int) Object {
	if n < 0 || n >= len(args) {
		return nil
	}
	return args[n]
}

func getTwoIntArgs(args ArgsObject) (int, int, error) {
	var lhs, rhs int
	switch v := getArg(args, 0).(type) {
	case NumberObject:
		lhs = int(v)
	default:
		return 0, 0, typeError(numberType, v)
	}
	switch v := getArg(args, 1).(type) {
	case NumberObject:
		rhs = int(v)
	default:
		return 0, 0, typeError(numberType, v)
	}
	return lhs, rhs, nil
}

func typeError(want reflect.Type, have interface{}) error {
	return fmt.Errorf("Wrong type, want %v, have (%T)(%v)", want, have, have)
}
