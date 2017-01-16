# mini

[![GoDoc](https://godoc.org/github.com/jncornett/mini?status.svg)](https://godoc.org/github.com/jncornett/mini)

[![TravisCI](https://travis-ci.org/jncornett/mini.svg?branch=master)](https://travis-ci.org/jncornett/mini)

## quickstart

### install

    go get github.com/jncornett/mini/cmd/mini
    
### run

To view usage:

    mini -h
    
To run the interactive interpreter:

    mini
    
To run a script

    mini myscript.mini
    
## develop

    go get github.com/jncornett/mini
    
### structure

- scanner/lexer is in `scanner.go`
- AST notes are in `ast.go`
- parser is in `parser.go`
- see `cmd/mini/main.go` for an implementation example
- see `examples/` for script examples

### todo

- add functions
- write the grammar in EBNF
- add a language reference
