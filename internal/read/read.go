package read

import "github.com/JairAntonio22/scheme-R7RS/internal/scm"

func ReadValue(s string) (scm.Value, error) {
	l := newLexer(s)
	toks := l.tokens()
	p := newParser(toks)
	return p.value()
}
