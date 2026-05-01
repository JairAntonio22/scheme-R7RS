package eval

type Nil struct{}

func (n Nil) sealed() {}

var _ Value = (*Nil)(nil)
