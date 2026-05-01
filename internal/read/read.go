package read

import (
	"errors"
	"fmt"

	"github.com/JairAntonio22/scheme-R7RS/internal/eval"
)

func Read(s string) (eval.Value, error) {
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
