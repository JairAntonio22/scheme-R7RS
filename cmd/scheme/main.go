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
	(define make-adder
		(lambda (x)
			(lambda (y) (+ x y))))

	(define add5 (make-adder 5))

	(add5 3)
	`

	value, err := read.Read(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	env := eval.DefaultEnv()
	value, err = eval.Eval(value, env)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(print.Print(value))
	}
}
