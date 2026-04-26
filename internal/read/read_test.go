package read_test

import (
	"testing"

	"github.com/JairAntonio22/scheme-R7RS/internal/read"
	"github.com/JairAntonio22/scheme-R7RS/internal/scm"
)

var tests = []struct {
	name       string
	input      string
	want       scm.Value
	wantErr    error
	skipReason string
}{
	// numbers
	{
		name:  "number",
		input: "42",
		want:  scm.Number(42),
	},
	{
		name:       "negative number",
		input:      "-15",
		want:       scm.Number(-15),
		skipReason: "not implemented",
	},
	{
		name:       "float number",
		input:      "3.14",
		want:       scm.Number(3),
		skipReason: "not implemented",
	},
	// booleans
	{
		name:  "true boolean",
		input: "#t",
		want:  scm.Boolean(true),
	},
	{
		name:  "false boolean",
		input: "#f",
		want:  scm.Boolean(false),
	},
	// symbols
	{
		name:  "symbol with letters",
		input: "hello",
		want:  scm.Symbol("hello"),
	},
	{
		name:  "symbol with no letters",
		input: "+",
		want:  scm.Symbol("+"),
	},
	{
		name:  "built in symbol",
		input: "car",
		want:  scm.Symbol("car"),
	},
	{
		name:  "special form symbol",
		input: "define",
		want:  scm.Symbol("define"),
	},
	{
		name:  "symbol with letters and graphs",
		input: "foo-bar",
		want:  scm.Symbol("foo-bar"),
	},
	{
		name:  "nil",
		input: "()",
		want:  scm.Nil(),
	},
	// lists
	{
		name:  "simple list",
		input: "(+ 1 2)",
		want:  scm.List(scm.Symbol("+"), scm.Number(1), scm.Number(2)),
	},
	{
		name:  "nested list",
		input: "(+ 1 (* 2 3))",
		want: scm.List(
			scm.Symbol("+"),
			scm.Number(1),
			scm.List(scm.Symbol("*"), scm.Number(2), scm.Number(3)),
		),
	},
	{
		name:  "multi-level list",
		input: "((1 2) (3 4))",
		want: scm.List(
			scm.List(scm.Number(1), scm.Number(2)),
			scm.List(scm.Number(3), scm.Number(4)),
		),
	},
	// quote
	{
		name:  "simple quote",
		input: "'a",
		want:  scm.List(scm.Quote(), scm.Symbol("a")),
	},
	{
		name:  "list quote",
		input: "'(1 2 3)",
		want:  scm.List(scm.Quote(), scm.List(scm.Number(1), scm.Number(2), scm.Number(3))),
	},
	{
		name:  "nested quote",
		input: "''a",
		want:  scm.List(scm.Quote(), scm.List(scm.Quote(), scm.Symbol("a"))),
	},
	// whitespace
	{
		name:  "extra whitespace",
		input: "   (+   1   2 )",
		want:  scm.List(scm.Symbol("+"), scm.Number(1), scm.Number(2)),
	},
	{
		name: "extra whitespace",
		input: `(
  +
  1
  2
)`,
		want: scm.List(scm.Symbol("+"), scm.Number(1), scm.Number(2)),
	},
	// errors
	{
		name:       "missing closing parenthesis",
		input:      "(+",
		skipReason: "TODO",
	},
	{
		name:       "extra parenthesis",
		input:      ")",
		skipReason: "TODO",
	},
	{
		name:       "unclosed list",
		input:      "(+ 1 2",
		skipReason: "TODO",
	},
	{
		name:       "empty",
		input:      "",
		skipReason: "TODO",
	},
	{
		name:       "invalid quote",
		input:      "'",
		skipReason: "TODO",
	},
	{
		name:       "invalid quote",
		input:      "'",
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

			got, err := read.ReadValue(test.input)
			if err != nil {
				t.Errorf("error %v", err)
			}

			if scm.Equal(got, test.want) == scm.False {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
