package eval_test

import (
	"errors"
	"testing"

	"github.com/JairAntonio22/scheme-R7RS/internal/eval"
	"github.com/JairAntonio22/scheme-R7RS/internal/read"
)

var tests = []struct {
	name       string
	input      []string
	want       eval.Value
	wantErr    error
	skipReason string
}{
	// self evaluation
	{
		name:  "number self evaluation",
		input: []string{"42"},
		want:  eval.Number(42),
	},
	{
		name:  "boolean self evaluation",
		input: []string{"#f"},
		want:  eval.False,
	},
	{
		name:  "nil self evaluation",
		input: []string{"()"},
		want:  eval.Nil{},
	},
	{
		name:  "symbol lookup",
		input: []string{"(define x 10)", "x"},
		want:  eval.Number(10),
	},
	// basic apply
	{
		name:  "simple apply",
		input: []string{"(+ 1 2)"},
		want:  eval.Number(3),
	},
	{
		name:  "nested apply",
		input: []string{"(+ 1 (+ 2 3))"},
		want:  eval.Number(6),
	},
	{
		name:       "multiple args apply",
		input:      []string{"(+ 1 2 3 4)"},
		want:       eval.Number(10),
		skipReason: "TODO",
	},
	// quote
	{
		name:  "quote symbol",
		input: []string{"'x"},
		want:  eval.Symbol("x"),
	},
	{
		name:  "quote list",
		input: []string{"'(1 2 3)"},
		want:  eval.List(eval.Number(1), eval.Number(2), eval.Number(3)),
	},
	{
		name:  "quote nested",
		input: []string{"'(1 (2 3) 4)"},
		want:  eval.List(eval.Number(1), eval.List(eval.Number(2), eval.Number(3)), eval.Number(4)),
	},
	// lambda
	{
		name:  "lambda inmediate application",
		input: []string{"((lambda (x) (+ x 1)) 5)"},
		want:  eval.Number(6),
	},
	{
		name:  "lambda multi params",
		input: []string{"((lambda (x y) (+ x y)) 3 4)"},
		want:  eval.Number(7),
	},
	{
		name:  "lambda body multiple expressions",
		input: []string{"((lambda (x) (+ x 1) (+ x 2)) 5)"},
		want:  eval.Number(7),
	},
	// closure
	{
		name:  "closure application",
		input: []string{"(((lambda (x) (lambda (y) (+ x y))) 5) 3)"},
		want:  eval.Number(8),
	},
	{
		name:  "closure definition",
		input: []string{"(define add5 ((lambda (x) (lambda (y) (+ x y))) 5))", "(add5 3)"},
		want:  eval.Number(8),
	},
	// scope and shadowing
	{
		name:  "parameter shadowing",
		input: []string{"(define x 10)", "((lambda (x) (+ x 1)) 5)"},
		want:  eval.Number(6),
	},
	{
		name: "closure shadowing",
		input: []string{
			"(define x 100)",
			"((lambda (x) (lambda () x)) 5)",
			"(((lambda (x) (lambda () x)) 5))",
		},
		want: eval.Number(5),
	},
	// dynamic operator
	{
		name:  "operator is expression",
		input: []string{"((lambda (f) (f 2 3)) +)"},
		want:  eval.Number(5),
	},
	{
		name:  "operator from if",
		input: []string{"((if #t + cons) 3 2)"},
		want:  eval.Number(5),
	},
	// special forms
	{
		name:  "if true branch only",
		input: []string{"(if #t 1 (/ 1 0))"},
		want:  eval.Number(1),
	},
	{
		name:  "if false branch only",
		input: []string{"(if #f (/ 1 0) 1)"},
		want:  eval.Number(1),
	},
	// errors
	{
		name:       "undefined symbol",
		input:      []string{"x"},
		skipReason: "TODO",
	},
	{
		name:       "arity mismatch",
		input:      []string{"(+ 1)"},
		skipReason: "TODO",
	},
	{
		name:       "calling non function",
		input:      []string{"(1 2 3)"},
		skipReason: "TODO",
	},
}

func TestEval(t *testing.T) {
	t.Parallel()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.skipReason != "" {
				t.Skip(test.skipReason)
			}

			var lastGot eval.Value
			env := eval.DefaultEnv()

			for i := range test.input {
				val, _ := read.Value(test.input[i])
				got, err := eval.Eval(val, env)
				if err != nil && !errors.Is(err, test.wantErr) {
					t.Errorf("got %v, want %v", err, test.wantErr)
				}

				lastGot = got
			}

			if eval.Equal(lastGot, test.want) == eval.False {
				t.Errorf("got %v, want %v", lastGot, test.want)
			}
		})
	}
}
