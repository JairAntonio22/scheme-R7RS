package main

import (
	"fmt"

	"github.com/JairAntonio22/scheme-R7RS/internal/read"
	"github.com/k0kubun/pp/v3"
)

func main() {
	fmt.Println()
	s := "(+ (* 2 3) '1)"

	value, err := read.ReadValue(s)
	if err != nil {
		fmt.Println(err)
	}

	_, _ = pp.Println(value)
}
