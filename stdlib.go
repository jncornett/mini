package mini

import "fmt"

type StdlibEntry struct {
	Name string
	Func Function
}

func GetStdlib() []StdlibEntry {
	return []StdlibEntry{
		{
			"print",
			func(args []Object) (Object, error) {
				var out []interface{}
				for _, arg := range args {
					out = append(out, arg.Value())
				}
				fmt.Println(out...)
				return nil, nil
			},
		},
	}
}
