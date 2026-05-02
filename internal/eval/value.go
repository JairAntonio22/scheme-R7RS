package eval

// Value
// Everything that gets evaluated is a Value
// Never use golang's nil as a Value.
type Value interface {
	sealed()
}

// Number
// This will get more robust in the future.
type Number int

func (n Number) sealed() {}

type Boolean bool

func (b Boolean) sealed() {}

type Symbol string

func (s Symbol) sealed() {}

var (
	True  = Boolean(true)
	False = Boolean(false)
)

type Nil struct{}

func (n Nil) sealed() {}

// Pair
// car and cdr must be valid Values.
type Pair struct {
	Car Value
	Cdr Value
}

func (p *Pair) sealed() {}

// Function
// env is never nil
// params are symbols
// body has at least one expression.
type Function struct {
	params []Value
	body   []Value
	env    *Env
}

func (f Function) sealed() {}

type BuiltIn struct {
	name     string
	argc     int
	callback Callback
}

type Callback func(args ...Value) Value

func (bi BuiltIn) sealed() {}
