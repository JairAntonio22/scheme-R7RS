package eval_test

import (
	"errors"
	"testing"

	"github.com/JairAntonio22/scheme-R7RS/internal/eval"
	"github.com/JairAntonio22/scheme-R7RS/internal/read"
)

var tests = []struct {
	name       string
	input      string
	env        *eval.Env
	want       eval.Value
	wantErr    error
	skipReason string
}{
	// self evaluation
	{
		name:  "number self evaluation",
		input: "42",
		env:   eval.DefaultEnv(),
		want:  eval.Number(42),
	},
	{
		name:  "boolean self evaluation",
		input: "#f",
		env:   eval.DefaultEnv(),
		want:  eval.False,
	},
	{
		name:  "nil self evaluation",
		input: "()",
		env:   eval.DefaultEnv(),
		want:  eval.Nil{},
	},
	{
		name:  "symbol lookup",
		input: "x",
		env: &eval.Env{
			Bindings: map[eval.Value]eval.Value{
				eval.Symbol("x"): eval.Number(10),
			},
			Parent: nil,
		},
		want: eval.Number(10),
	},
	// basic apply
	{
		name:  "simple apply",
		input: "(+ 1 2)",
		env:   eval.DefaultEnv(),
		want:  eval.Number(3),
	},
	{
		name:  "nested apply",
		input: "(+ 1 (+ 2 3))",
		env:   eval.DefaultEnv(),
		want:  eval.Number(6),
	},
	{
		name:       "multiple args apply",
		input:      "(+ 1 2 3 4)",
		env:        eval.DefaultEnv(),
		want:       eval.Number(10),
		skipReason: "TODO",
	},
	// quote
	{
		name:  "quote symbol",
		input: "'x",
		env:   eval.DefaultEnv(),
		want:  eval.Symbol("x"),
	},
	{
		name:  "quote list",
		input: "'(1 2 3)",
		env:   eval.DefaultEnv(),
		want:  eval.List(eval.Number(1), eval.Number(2), eval.Number(3)),
	},
	{
		name:  "quote nested",
		input: "'(1 (2 3) 4)",
		env:   eval.DefaultEnv(),
		want:  eval.List(eval.Number(1), eval.List(eval.Number(2), eval.Number(3)), eval.Number(4)),
	},
	// lambda
	{
		name:  "lambda inmediate application",
		input: "((lambda (x) (+ x 1)) 5)",
		env:   eval.DefaultEnv(),
		want:  eval.Number(6),
	},
	{
		name:  "lambda multi params",
		input: "((lambda (x y) (+ x y)) 3 4)",
		env:   eval.DefaultEnv(),
		want:  eval.Number(7),
	},
	{
		name:  "lambda body multiple expressions",
		input: "((lambda (x) (+ x 1) (+ x 2)) 5)",
		env:   eval.DefaultEnv(),
		want:  eval.Number(7),
	},
	// closure
	{
		name:  "closure application",
		input: "(((lambda (x) (lambda (y) (+ x y))) 5) 3)",
		env:   eval.DefaultEnv(),
		want:  eval.Number(8),
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

			val, _ := read.Read(test.input)
			got, err := eval.Eval(val, test.env)
			if err != nil && !errors.Is(err, test.wantErr) {
				t.Errorf("got %v, want %v", err, test.wantErr)
			}

			if eval.Equal(got, test.want) == eval.False {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
