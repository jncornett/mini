package mini

import (
	"errors"
	"fmt"
	"reflect"
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
			func(args []Object) (Object, error) {
				return BoolObject(!GetBoolValue(getArg(args, 0))), nil
			},
		},
		{
			"__neg",
			func(args []Object) (Object, error) {
				// currently only defined for numbers
				switch v := getArg(args, 0).(type) {
				case NumberObject:
					return NumberObject(-int(v)), nil
				default:
					return nil, typeError("number", v)
				}
			},
		},
		{
			"__add",
			func(args []Object) (Object, error) {
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
			func(args []Object) (Object, error) {
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
			func(args []Object) (Object, error) {
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
			func(args []Object) (Object, error) {
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
			func(args []Object) (Object, error) {
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
			func(args []Object) (Object, error) {
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
			func(args []Object) (Object, error) {
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
			func(args []Object) (Object, error) {
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
			func(args []Object) (Object, error) {
				return BoolObject(reflect.DeepEqual(getArg(args, 0), getArg(args, 1))), nil
			},
		},
		{
			"__ne",
			func(args []Object) (Object, error) {
				return BoolObject(!reflect.DeepEqual(getArg(args, 0), getArg(args, 1))), nil
			},
		},
		{
			"__and",
			func(args []Object) (Object, error) {
				lhs := GetBoolValue(getArg(args, 0))
				if !lhs {
					return BoolObject(false), nil // Shortcircuit
				}
				rhs := GetBoolValue(getArg(args, 1))
				if !rhs {
					return BoolObject(false), nil
				}
				return getArg(args, 1), nil
			},
		},
		{
			"__or",
			func(args []Object) (Object, error) {
				lhs := GetBoolValue(getArg(args, 0))
				if lhs {
					return getArg(args, 0), nil // Shortcircuit
				}
				rhs := GetBoolValue(getArg(args, 1))
				if rhs {
					return getArg(args, 1), nil
				}
				return BoolObject(false), nil
			},
		},
		// stdlib
		{
			"print",
			func(args []Object) (Object, error) {
				_, err := fmt.Println(objectsToEmpties(args)...)
				return nil, err
			},
		},
	}
}

func objectsToEmpties(args []Object) []interface{} {
	var out []interface{}
	for _, arg := range args {
		out = append(out, arg.Value())
	}
	return out
}

func getArg(args []Object, n int) Object {
	if n < 0 || n >= len(args) {
		return nil
	}
	return args[n]
}

func getTwoIntArgs(args []Object) (int, int, error) {
	var lhs, rhs int
	switch v := getArg(args, 0).(type) {
	case NumberObject:
		lhs = int(v)
	default:
		return 0, 0, typeError("number", v)
	}
	switch v := getArg(args, 1).(type) {
	case NumberObject:
		rhs = int(v)
	default:
		return 0, 0, typeError("number", v)
	}
	return lhs, rhs, nil
}

func getTwoBoolValues(args []Object) (bool, bool) {
	return GetBoolValue(getArg(args, 0)), GetBoolValue(getArg(args, 1))
}

func typeError(expected string, actual interface{}) error {
	return fmt.Errorf("TypeError: expected %v, got %T", expected, actual)
}
