package read

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/JairAntonio22/scheme-R7RS/internal/eval"
)

var (
	ErrTrailingTokens      = errors.New("unexpected trailing tokens")
	ErrUnexpectedToken     = errors.New("unexpected token")
	ErrMissingClosingParen = errors.New("missing closing parenthesis")
	ErrInvalidNumber       = errors.New("invalid number")
	ErrInvalidBoolean      = errors.New("invalid boolean")
	ErrInvalidToken        = errors.New("invalid token")
)

type parser struct {
	lex  *lexer
	curr token
}

func newParser(s string) *parser {
	p := parser{lex: newLexer(s)}
	p.advance()
	return &p
}

func (parse *parser) advance() {
	if parse.curr.typ == eof {
		return
	}

	parse.curr = parse.lex.token()
}

func (parse *parser) value() (eval.Value, error) {
	switch parse.curr.typ {
	case symbol:
		val := eval.Symbol(parse.curr.str)
		parse.advance()
		return val, nil

	case number:
		val, err := parseNumber(parse.curr)
		parse.advance()
		return val, err

	case boolean:
		val, err := parseBoolean(parse.curr)
		parse.advance()
		return val, err

	case lParen:
		parse.advance()
		return parse.list()

	case quote:
		parse.advance()
		quoted, err := parse.value()
		return eval.List(eval.Quote, quoted), err

	case eof:
		err := fmt.Errorf("%w: got %s", ErrUnexpectedToken, parse.curr.typ)
		return eval.Nil{}, err

	case illegal, rParen:
		err := fmt.Errorf("%w: got %s", ErrUnexpectedToken, parse.curr.str)
		parse.advance()
		return eval.Nil{}, err

	default:
		panic(fmt.Errorf("%w: got %v", ErrInvalidToken, parse.curr.typ))
	}
}

func parseNumber(tok token) (eval.Value, error) {
	num, err := strconv.Atoi(tok.str)
	if err != nil {
		err = fmt.Errorf("%w: could not read %s", ErrInvalidNumber, tok.str)
	}

	return eval.Number(num), err
}

func parseBoolean(tok token) (eval.Value, error) {
	switch tok.str {
	case "#t":
		return eval.True, nil

	case "#f":
		return eval.False, nil

	default:
		panic(fmt.Errorf("%w: got %s", ErrInvalidBoolean, tok.str))
	}
}

func (parse *parser) list() (eval.Value, error) {
	switch parse.curr.typ {
	case lParen, quote, number, boolean, symbol:
		car, carErr := parse.value()
		cdr, cdrErr := parse.list()
		return eval.Cons(car, cdr), errors.Join(carErr, cdrErr)

	case rParen:
		parse.advance()
		return eval.Nil{}, nil

	case eof:
		return eval.Nil{}, ErrMissingClosingParen

	case illegal:
		parse.advance()
		return eval.Nil{}, fmt.Errorf("%w: got %s", ErrUnexpectedToken, parse.curr.str)

	default:
		panic(fmt.Errorf("%w: got %v", ErrInvalidToken, parse.curr.typ))
	}
}
