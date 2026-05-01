package main

import (
	"fmt"

	"github.com/JairAntonio22/scheme-R7RS/internal/eval"
	"github.com/JairAntonio22/scheme-R7RS/internal/read"
	"github.com/k0kubun/pp/v3"
)

func main() {
	fmt.Println()

	s := "(list 'a b)"

	value, err := read.Read(s)
	if err != nil {
		fmt.Println(err)
	}

	value, err = eval.Eval(value, eval.DefaultEnv())
	if err != nil {
		fmt.Println(err)
	}

	_, _ = pp.Println(value)
}
