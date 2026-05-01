package eval

// Number
// This will get more robust in the future.
type Number int

func (n Number) sealed() {}

var _ Value = (*Number)(nil)
