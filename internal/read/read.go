package read

import (
	"errors"
	"fmt"

	"github.com/JairAntonio22/scheme-R7RS/internal/eval"
)

func Value(s string) (eval.Value, error) {
	parse := newParser(s)

	val, err := parse.value()
	toks := make([]token, 0)

	for parse.curr.typ != eof {
		parse.advance()
		toks = append(toks, parse.curr)
	}

	if len(toks) > 0 {
		trailErr := fmt.Errorf("%w: got '%v'", ErrTrailingTokens, toks)
		err = errors.Join(err, trailErr)
	}

	return val, err
}

func Program(s string) ([]eval.Value, error) {
	parse := newParser(s)
	program := make([]eval.Value, 0)
	var err error

	for parse.curr.typ != eof {
		val, valErr := parse.value()
		program = append(program, val)
		err = errors.Join(err, valErr)
	}

	return program, err
}
