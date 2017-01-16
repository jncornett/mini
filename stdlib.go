package mini

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	numberType = reflect.TypeOf(Number(0))
)

type StdlibEntry struct {
	Name Symbol
	Func Function
}

// FIXME this isn't the best way to organize functionality
func GetStdlib() []StdlibEntry {
	return []StdlibEntry{
		// builtins
		{
			"__not",
			func(args Args) (Object, error) {
				return NewBoolFromBool(!args.Arg(0).Truthy()), nil
			},
		},
		{
			"__neg",
			func(args Args) (Object, error) {
				num, ok := args.Arg(0).(Number)
				// currently only defined for numbers
				if !ok {
					return nil, typeError(numberType, num)
				}
				switch v := getArg(args, 0).(type) {
				case Number:
					return Number(-int(v)), nil
				default:
					return nil, typeError(numberType, num)
				}
			},
		},
		{
			"__add",
			func(args Args) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return Number(lhs + rhs), nil
			},
		},
		{
			"__sub",
			func(args Args) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return Number(lhs - rhs), nil
			},
		},
		{
			"__mul",
			func(args Args) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return Number(lhs * rhs), nil
			},
		},
		{
			"__div",
			func(args Args) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				if rhs == 0 {
					return nil, errors.New("ZeroDivisionError")
				}
				return Number(lhs / rhs), nil
			},
		},
		{
			"__lt",
			func(args Args) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return Bool(lhs < rhs), nil
			},
		},
		{
			"__le",
			func(args Args) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return Bool(lhs <= rhs), nil
			},
		},
		{
			"__gt",
			func(args Args) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return Bool(lhs > rhs), nil
			},
		},
		{
			"__ge",
			func(args Args) (Object, error) {
				// currently only defined for numbers
				lhs, rhs, err := getTwoIntArgs(args)
				if err != nil {
					return nil, err
				}
				return Bool(lhs >= rhs), nil
			},
		},
		{
			"__eq",
			func(args Args) (Object, error) {
				return Bool(reflect.DeepEqual(getArg(args, 0), getArg(args, 1))), nil
			},
		},
		{
			"__ne",
			func(args Args) (Object, error) {
				return Bool(!reflect.DeepEqual(getArg(args, 0), getArg(args, 1))), nil
			},
		},
		{
			"__and",
			func(args Args) (Object, error) {
				lhs := args.Arg(0)
				if !lhs.Truthy() {
					// short circuit if we can
					return Bool(false), nil
				}
				rhs := args.Arg(1)
				if !rhs.Truthy() {
					return Bool(false), nil
				}
				return rhs, nil
			},
		},
		{
			"__or",
			func(args Args) (Object, error) {
				lhs := args.Arg(0)
				if lhs.Truthy() {
					// short circuit if possible
					return lhs, nil
				}
				rhs := args.Arg(1)
				if rhs.Truthy() {
					return rhs, nil
				}
				return Bool(false), nil
			},
		},
		// stdlib
		{
			"print",
			func(args Args) (Object, error) {
				_, err := fmt.Println(objectsToEmpties(args)...)
				return nil, err
			},
		},
	}
}

func objectsToEmpties(args Args) []interface{} {
	var out []interface{}
	for _, arg := range args {
		out = append(out, arg)
	}
	return out
}

func getArg(args Args, n int) Object {
	if n < 0 || n >= len(args) {
		return nil
	}
	return args[n]
}

func getTwoIntArgs(args Args) (int, int, error) {
	var lhs, rhs int
	switch v := getArg(args, 0).(type) {
	case Number:
		lhs = int(v)
	default:
		return 0, 0, typeError(numberType, v)
	}
	switch v := getArg(args, 1).(type) {
	case Number:
		rhs = int(v)
	default:
		return 0, 0, typeError(numberType, v)
	}
	return lhs, rhs, nil
}

func typeError(want reflect.Type, have interface{}) error {
	return fmt.Errorf("Wrong type, want %v, have (%T)(%v)", want, have, have)
}
