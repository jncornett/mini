package mini

import "fmt"

type Vm struct{}

func NewVm() *Vm {
	return &Vm{}
}

func (vm Vm) Eval(s string) error {
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	fmt.Println("Would eval:", s)
	return nil
}
