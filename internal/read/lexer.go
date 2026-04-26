package read

import (
	"strings"
	"unicode"
)

type lexer struct {
	input []rune
	pos   int
	ch    rune
}

func newLexer(s string) *lexer {
	l := lexer{input: []rune(s)}
	l.readRune()
	return &l
}

func (l *lexer) tokens() []token {
	tokens := make([]token, 0)

	for t := l.nextToken(); t.typ != eof; t = l.nextToken() {
		tokens = append(tokens, t)
	}

	return tokens
}

func (l *lexer) readRune() {
	if l.pos >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.pos]
	}

	l.pos++
}

func (l *lexer) nextToken() token {
	l.skipWhitespace()

	switch l.ch {
	case '(':
		l.readRune()
		return token{typ: lParen, str: "("}

	case ')':
		l.readRune()
		return token{typ: rParen, str: ")"}

	case '\'':
		l.readRune()
		return token{typ: quote, str: "'"}

	case 0:
		return token{typ: eof, str: ""}

	default:
		if unicode.IsDigit(l.ch) {
			return token{typ: number, str: l.readNumber()}
		} else {
			switch str := l.readStr(); str {
			case "#t", "#f":
				return token{typ: boolean, str: str}

			default:
				return token{typ: symbol, str: str}
			}
		}
	}
}

func (l *lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readRune()
	}
}

func (l *lexer) readNumber() string {
	start := l.pos - 1

	for unicode.IsDigit(l.ch) {
		l.readRune()
	}

	return string(l.input[start : l.pos-1])
}

func (l *lexer) readStr() string {
	start := l.pos - 1

	for l.ch != 0 && !unicode.IsSpace(l.ch) && !strings.ContainsRune("()'", l.ch) {
		l.readRune()
	}

	return string(l.input[start : l.pos-1])
}
