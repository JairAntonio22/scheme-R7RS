package eval

// Function
// env is never nil
// params are symbols
// body has at least one expression.
type Function struct {
	params []Symbol
	body   []Value
	env    *Env
}

func (f Function) sealed() {}

var _ Value = (*Function)(nil)
