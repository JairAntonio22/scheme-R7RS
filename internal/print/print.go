package print

import (
	"strconv"

	"github.com/JairAntonio22/scheme-R7RS/internal/eval"
)

func Print(val eval.Value) string {
	switch v := val.(type) {
	case eval.Number:
		return strconv.Itoa(int(v))

	case eval.Boolean:
		if v == eval.True {
			return "#t"
		}

		return "#f"

	case eval.Symbol:
		return string(v)

	case eval.Nil:
		return "()"

	case *eval.Pair:
		car := Print(v.Car)
		cdr := Print(v.Cdr)

		return "(" + car + " . " + cdr + ")"

	default:
		return "<unknown>"
	}
}
