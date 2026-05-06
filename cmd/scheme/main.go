package main

import (
	"fmt"

	"github.com/JairAntonio22/scheme-R7RS/internal/eval"
	"github.com/JairAntonio22/scheme-R7RS/internal/print"
	"github.com/JairAntonio22/scheme-R7RS/internal/read"
)

func main() {
	fmt.Println()

	input := `
		(define add5 ((lambda (x) (lambda (y) (+ x y))) 5))
		(add5 3)
	`

	program, err := read.Program(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	env := eval.DefaultEnv()

	for _, value := range program {
		print.Print(value)

		value, err = eval.Eval(value, env)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(print.Print(value))
		}
	}
}
