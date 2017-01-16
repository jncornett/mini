package mini

import "fmt"

var StdLib []Entry = []Entry{
	{
		"print",
		func(args Args) (Object, error) {
			_, err := fmt.Println(objectsToEmpties(args)...)
			return nil, err
		},
	},
}

func objectsToEmpties(args Args) []interface{} {
	var out []interface{}
	for _, arg := range args {
		out = append(out, arg)
	}
	return out
}
