package read

import (
	"errors"
	"fmt"
	"strconv"

	scheme "github.com/JairAntonio22/scheme-R7RS/internal/scm"
)

type parser struct {
	input []token
	pos   int
	curr  token
	errs  []error
}

func newParser(toks []token) *parser {
	p := parser{input: toks}
	p.advance()
	return &p
}

var (
	ErrMissingRightParen = errors.New("missing closing parenthesis")
	ErrUnexpectedToken   = errors.New("unexpected token")
	ErrTrailingTokens    = errors.New("unexpected trailing tokens")
	ErrInvalidNumber     = errors.New("invalid number")
)

func (p *parser) value() (scheme.Value, error) {
	val := p.parseValue()

	if p.curr.typ != eof {
		p.errs = append(p.errs, ErrTrailingTokens)
	}

	if len(p.errs) != 0 {
		return val, &ReadError{errs: p.errs}
	}

	return val, nil
}

func (p *parser) advance() {
	if p.pos >= len(p.input) {
		p.curr.typ = eof
		return
	}

	p.curr = p.input[p.pos]
	p.pos++
}

func (p *parser) parseValue() scheme.Value {
	val := scheme.Nil()

	switch p.curr.typ {
	case symbol:
		val = scheme.Symbol(p.curr.str)
		p.advance()

	case number:
		num, err := strconv.Atoi(p.curr.str)

		if err != nil {
			err := fmt.Errorf("%w: could not read %s", ErrInvalidNumber, p.curr.str)
			p.errs = append(p.errs, err)

		} else {
			val = scheme.Number(num)
		}

		p.advance()

	case lParen:
		p.advance()
		val = p.parseList()

	case quote:
		p.advance()
		val = p.parseValue()
		return scheme.List(scheme.Quote(), val)

	default:
		err := fmt.Errorf("%w: got %s", ErrUnexpectedToken, p.curr.typ)
		p.errs = append(p.errs, err)
		p.advance()
	}

	return val
}

func (p *parser) parseList() scheme.Value {
	val := scheme.Nil()

	switch p.curr.typ {
	case rParen:
		p.advance()

	case eof:
		p.errs = append(p.errs, ErrMissingRightParen)

	default:
		car := p.parseValue()
		cdr := p.parseList()
		val = scheme.Cons(car, cdr)
	}

	return val
}
