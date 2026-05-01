package eval

import "fmt"

var builtIns = map[Value]Value{
	Symbol("cons"):   &BuiltIn{name: "cons", argc: 2, callback: Cons},
	Symbol("list"):   &BuiltIn{name: "list", argc: 2, callback: List},
	Symbol("equal?"): &BuiltIn{name: "equal?", argc: 2, callback: Equal},
	Symbol("pair?"):  &BuiltIn{name: "pair?", argc: 1, callback: IsPair},
	Symbol("list?"):  &BuiltIn{name: "list?", argc: 1, callback: IsList},
}

type BuiltIn struct {
	name     string
	argc     int
	callback Callback
}

type Callback func(args ...Value) Value

func (bi BuiltIn) sealed() {}

func Cons(args ...Value) Value {
	return &Pair{Car: args[0], Cdr: args[1]}
}

func List(args ...Value) Value {
	var list Value = Nil{}

	for i := len(args) - 1; i >= 0; i-- {
		list = Cons(args[i], list)
	}

	return list
}

func Equal(args ...Value) Value { //nolint:cyclop // this method is straightfoward
	switch val1 := args[0].(type) {
	case Number:
		val2, ok := args[1].(Number)
		return Boolean(ok && val1 == val2)

	case Boolean:
		val2, ok := args[1].(Boolean)
		return Boolean(ok && val1 == val2)

	case Symbol:
		val2, ok := args[1].(Symbol)
		return Boolean(ok && val1 == val2)

	case Nil:
		_, ok := args[1].(Nil)
		return Boolean(ok)

	case *Pair:
		val2, ok := args[1].(*Pair)
		return Boolean(ok && Equal(val1.Car, val2.Car) == True && Equal(val1.Cdr, val2.Cdr) == True)

	default:
		panic(fmt.Sprintf("equal: unsupported type %T", val1))
	}
}

func IsPair(args ...Value) Value {
	_, ok := args[0].(*Pair)
	return Boolean(ok)
}

func IsList(args ...Value) Value {
	switch v := args[0].(type) {
	case Nil:
		return True

	case *Pair:
		return Boolean(IsList(v.Cdr) == True)

	default:
		return False
	}
}
