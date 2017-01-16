package mini

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type SymbolTable map[Symbol]Object

type Vm struct {
	Symbols SymbolTable
	Result  Object
	Debug   bool
}

func NewVm() *Vm {
	return &Vm{
		Symbols: getDefaultSymbols(),
	}
}

func (vm *Vm) Eval(r io.Reader) error {
	expr, err := NewParser(r).Parse()
	if vm.Debug {
		log.Println("AST:", expr)
	}
	if err != nil {
		return err
	}
	vm.Result, err = expr.Eval(vm)
	if vm.Debug {
		log.Println("Symbols:", vm.Symbols)
	}
	return err
}

func (vm *Vm) EvalString(s string) error {
	return vm.Eval(strings.NewReader(s))
}

func (vm *Vm) Assign(sym Symbol, obj Object) {
	vm.Symbols[sym] = obj
}

func (vm *Vm) Lookup(sym Symbol) Object {
	return vm.Symbols[sym]
}

func (vm *Vm) Call(sym Symbol, args Args) (Object, error) {
	fn, ok := vm.Lookup(sym).(Callable)
	if !ok {
		return nil, fmt.Errorf("TypeError: %v is not a function", sym)
	}
	return fn.Call(args)
}

func getDefaultSymbols() SymbolTable {
	symbols := make(SymbolTable)
	for _, entry := range GetStdlib() {
		symbols[entry.Name] = entry.Func
	}
	return symbols
}
