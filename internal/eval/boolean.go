package eval

type Boolean bool

func (b Boolean) sealed() {}

var (
	True  = Boolean(true)
	False = Boolean(false)
)

var _ Value = (*Boolean)(nil)
