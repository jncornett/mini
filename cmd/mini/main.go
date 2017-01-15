package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jncornett/mini"
)

var prompt = "> "

func main() {
	var (
		vm = mini.NewVm()
		r  = bufio.NewReader(os.Stdin)
	)
	for {
		fmt.Print(prompt)
		code, _ := r.ReadString('\n')
		err := vm.EvalString(code)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(vm.Result)
		}
	}
}
