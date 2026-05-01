package read_test

import (
	"errors"
	"testing"

	"github.com/JairAntonio22/scheme-R7RS/internal/eval"
	"github.com/JairAntonio22/scheme-R7RS/internal/read"
)

var tests = []struct {
	name       string
	input      string
	want       eval.Value
	wantErr    error
	skipReason string
}{
	// numbers
	{
		name:  "number",
		input: "42",
		want:  eval.Number(42),
	},
	{
		name:       "negative number",
		input:      "-15",
		want:       eval.Number(-15),
		skipReason: "TODO",
	},
	{
		name:       "float number",
		input:      "3.14",
		want:       eval.Number(3),
		skipReason: "TODO",
	},
	// booleans
	{
		name:  "true boolean",
		input: "#t",
		want:  eval.Boolean(true),
	},
	{
		name:  "false boolean",
		input: "#f",
		want:  eval.Boolean(false),
	},
	// symbols
	{
		name:  "symbol with letters",
		input: "hello",
		want:  eval.Symbol("hello"),
	},
	{
		name:  "symbol with no letters",
		input: "+",
		want:  eval.Symbol("+"),
	},
	{
		name:  "built in symbol",
		input: "car",
		want:  eval.Symbol("car"),
	},
	{
		name:  "special form symbol",
		input: "define",
		want:  eval.Symbol("define"),
	},
	{
		name:  "symbol with letters and graphs",
		input: "foo-bar",
		want:  eval.Symbol("foo-bar"),
	},
	{
		name:  "nil",
		input: "()",
		want:  eval.Nil{},
	},
	// lists
	{
		name:  "simple list",
		input: "(+ 1 2)",
		want:  eval.List(eval.Symbol("+"), eval.Number(1), eval.Number(2)),
	},
	{
		name:  "nested list",
		input: "(+ 1 (* 2 3))",
		want: eval.List(
			eval.Symbol("+"),
			eval.Number(1),
			eval.List(eval.Symbol("*"), eval.Number(2), eval.Number(3)),
		),
	},
	{
		name:  "multi-level list",
		input: "((1 2) (3 4))",
		want: eval.List(
			eval.List(eval.Number(1), eval.Number(2)),
			eval.List(eval.Number(3), eval.Number(4)),
		),
	},
	// quote
	{
		name:  "simple quote",
		input: "'a",
		want:  eval.List(eval.Quote, eval.Symbol("a")),
	},
	{
		name:  "list quote",
		input: "'(1 2 3)",
		want:  eval.List(eval.Quote, eval.List(eval.Number(1), eval.Number(2), eval.Number(3))),
	},
	{
		name:  "nested quote",
		input: "''a",
		want:  eval.List(eval.Quote, eval.List(eval.Quote, eval.Symbol("a"))),
	},
	// whitespace
	{
		name:  "extra whitespace",
		input: "   (+   1   2 )",
		want:  eval.List(eval.Symbol("+"), eval.Number(1), eval.Number(2)),
	},
	{
		name: "extra whitespace",
		input: `(
  +
  1
  2
)`,
		want: eval.List(eval.Symbol("+"), eval.Number(1), eval.Number(2)),
	},
	// errors
	{
		name:    "missing closing parenthesis",
		input:   "(+",
		want:    eval.List(eval.Symbol("+")),
		wantErr: read.ErrMissingClosingParen,
	},
	{
		name:    "extra parenthesis",
		input:   ")",
		want:    eval.Nil{},
		wantErr: read.ErrUnexpectedToken,
	},
	{
		name:    "unclosed list",
		input:   "(+ 1 2",
		want:    eval.List(eval.Symbol("+"), eval.Number(1), eval.Number(2)),
		wantErr: read.ErrMissingClosingParen,
	},
	{
		name:    "empty",
		input:   "",
		want:    eval.Nil{},
		wantErr: read.ErrUnexpectedToken,
	},
	{
		name:    "invalid quote",
		input:   "'",
		want:    eval.List(eval.Quote, eval.Nil{}),
		wantErr: read.ErrUnexpectedToken,
	},
	{
		name:       "invalid dot",
		input:      ".",
		skipReason: "TODO",
	},
}

func TestReadValue(t *testing.T) {
	t.Parallel()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.skipReason != "" {
				t.Skip(test.skipReason)
			}

			got, err := read.Read(test.input)
			if err != nil && !errors.Is(err, test.wantErr) {
				t.Errorf("got %v, want %v", err, test.wantErr)
			}

			if eval.Equal(got, test.want) == eval.False {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
