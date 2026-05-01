package eval

// Pair
// car and cdr must be valid Values.
type Pair struct {
	Car Value
	Cdr Value
}

func (p Pair) sealed() {}

var _ Value = (*Pair)(nil)
