package eval

type Symbol string

func (s Symbol) sealed() {}

var _ Value = (*Symbol)(nil)
