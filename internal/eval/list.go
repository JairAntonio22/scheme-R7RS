package eval

import "fmt"

func toSlice(input Value) []Value {
	switch val := input.(type) {
	case Nil:
		return []Value{}

	case *Pair:
		slice := make([]Value, 0)
		curr := val
		isPair := true

		for isPair {
			slice = append(slice, curr.Car)
			curr, isPair = curr.Cdr.(*Pair)
		}

		return slice

	default:
		panic(fmt.Sprintf("cannot value %s to slice", val))
	}
}
