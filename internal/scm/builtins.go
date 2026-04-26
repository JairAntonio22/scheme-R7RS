package scm

func Cons(args ...Value) Value {
	return &Pair{Car: args[0], Cdr: args[1]}
}

var _ BuiltIn = Cons

func List(args ...Value) Value {
	list := Nil()

	for i := len(args) - 1; i >= 0; i-- {
		list = Cons(args[i], list)
	}

	return list
}

var _ BuiltIn = List

func Equal(args ...Value) Value {
	if args[0].Type() != args[1].Type() {
		return False
	}

	switch args[0].Type() {
	case NumberType, BooleanType, SymbolType, NilType:
		return Boolean(args[0] == args[1])

	case PairType:
		pair1 := args[0].(*Pair) //nolint:forcetypeassert // type check done at beginning
		pair2 := args[1].(*Pair) //nolint:forcetypeassert // type check done at beginning

		if Equal(pair1.Car, pair2.Car) == False {
			return False
		}

		return Equal(pair1.Cdr, pair2.Cdr)

	default:
		return False
	}
}

var _ BuiltIn = Equal
