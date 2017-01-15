package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jncornett/mini"
)

var prompt = "> "

func main() {
	var (
		debug = flag.Bool("debug", false, "turn on debug logging")
		repl  = flag.Bool("repl", false, "enter REPL mode")
	)
	flag.Parse()
	if flag.NArg() == 0 {
		*repl = true
	}
	vm := mini.NewVm()
	vm.Debug = *debug
	for _, script := range flag.Args() {
		err := runScript(vm, script)
		if err != nil {
			log.Fatal(err)
		}
	}
	if *repl {
		enterRepl(vm, ":-) ")
	}
}

func runScript(vm *mini.Vm, p string) error {
	f, err := os.Open(p)
	defer f.Close()
	if err != nil {
		return err
	}
	return vm.Eval(f)
}

func enterRepl(vm *mini.Vm, prompt string) {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		code, _ := r.ReadString('\n')
		err := vm.EvalString(code)
		if err != nil {
			fmt.Println("error:", err)
		} else if vm.Result != nil {
			fmt.Println("=>", vm.Result)
		}
	}
}
