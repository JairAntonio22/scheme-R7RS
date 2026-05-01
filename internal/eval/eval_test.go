package eval_test

import (
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
}{}

func TestReadValue(t *testing.T) {
	t.Parallel()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.skipReason != "" {
				t.Skip(test.skipReason)
			}

			value, _ := read.Read(test.input)
			got, err := eval.Eval(value, eval.DefaultEnv())
			if err != nil {
				t.Errorf("error %v", err)
			}

			if eval.Equal(got, test.want) == eval.False {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
