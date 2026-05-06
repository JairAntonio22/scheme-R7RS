package eval

import (
	"errors"
	"fmt"
)

var (
	ErrSymbolNotDefined  = errors.New("symbol is not defined")
	ErrSymbolNotCallable = errors.New("symbol not callable")
	ErrCannotEval        = errors.New("cannot evaluate")
	ErrArityMismatch     = errors.New("mismatch in number of arguments")
	ErrImproperList      = errors.New("improper list found")
)

func Eval(input Value, env *Env) (Value, error) {
	switch val := input.(type) {
	case Number:
		return input, nil

	case Boolean:
		return input, nil

	case Nil:
		return input, nil

	case Symbol:
		envVal, exists := env.LookUp(val)

		if !exists {
			return Nil{}, fmt.Errorf("%w: got %s", ErrSymbolNotDefined, input)
		}

		return envVal, nil

	case *Pair:
		slice := toSlice(val)
		sym, ok := slice[0].(Symbol)

		if ok {
			switch sym {
			case Quote:
				return QuoteForm(env, slice[1:]...), nil

			case Define:
				return DefineForm(env, slice[1:]...), nil

			case If:
				return IfForm(env, slice[1:]...), nil

			case Lambda:
				return LambdaForm(env, slice[1:]...), nil
			}
		}

		evalSlice := make([]Value, 0, len(slice))

		for i := range slice {
			evalArg, err := Eval(slice[i], env)
			if err != nil {
				return Nil{}, err
			}

			evalSlice = append(evalSlice, evalArg)
		}

		return Apply(evalSlice[0], evalSlice[1:]...)

	default:
		return Nil{}, fmt.Errorf("%w: got %s", ErrCannotEval, input)
	}
}

func Apply(symbol Value, args ...Value) (Value, error) {
	switch function := symbol.(type) {
	case *BuiltIn:
		if len(args) != function.argc {
			return Nil{}, fmt.Errorf(
				"%w: got %d, expected %d",
				ErrArityMismatch,
				len(args),
				function.argc,
			)
		}

		return function.callback(args...), nil

	case Function:
		child := NewEnv(function.env)

		if len(function.params) != len(args) {
			return Nil{}, fmt.Errorf(
				"%w: got %d, expected %d",
				ErrArityMismatch,
				len(args),
				len(function.params),
			)
		}

		for i := range function.params {
			child.Define(function.params[i], args[i])
		}

		var lastVal Value = Nil{}

		for _, expr := range function.body {
			val, err := Eval(expr, child)
			if err != nil {
				return val, err
			}

			lastVal = val
		}

		return lastVal, nil

	default:
		return Nil{}, fmt.Errorf("%w: got %s", ErrSymbolNotCallable, symbol)
	}
}
