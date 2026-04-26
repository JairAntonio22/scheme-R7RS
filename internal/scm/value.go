package scm

type ValueType int

const (
	NumberType ValueType = iota
	BooleanType
	SymbolType
	NilType
	PairType
	FunctionType
	BuiltInType
)

// Value
// Everything that gets evaluated is a Value
// Never use golang's nil as a Value.
//
//sumtype:decl
type Value interface {
	Type() ValueType
	sealed()
}

// Number
// This will get more robust in the future.
type Number int

func (n Number) Type() ValueType {
	return NumberType
}

func (n Number) sealed() {}

var _ Value = (*Number)(nil)

type Boolean bool

func (b Boolean) Type() ValueType {
	return BooleanType
}

func (b Boolean) sealed() {}

var (
	True  = Boolean(true)
	False = Boolean(false)
)

var _ Value = (*Boolean)(nil)

type Symbol string

func (s Symbol) Type() ValueType {
	return SymbolType
}

func (s Symbol) sealed() {}

var _ Value = (*Symbol)(nil)

type NilStruct struct{}

func (n NilStruct) Type() ValueType {
	return NilType
}

func (n NilStruct) sealed() {}

var nilValue = &NilStruct{}

var _ Value = (*NilStruct)(nil)

func Nil() Value {
	return nilValue
}

// Pair
// car and cdr must be valid Values.
type Pair struct {
	Car Value `json:"car"`
	Cdr Value `json:"cdr"`
}

func (p Pair) Type() ValueType {
	return PairType
}

func (p Pair) sealed() {}

var _ Value = (*Pair)(nil)

// Function
// env is never nil
// params are symbols
// body has at least one expression.
type Function struct {
	params []Symbol
	body   []Value
	env    *Env
}

func (f Function) Type() ValueType {
	return FunctionType
}

func (f Function) sealed() {}

var _ Value = (*Function)(nil)

type BuiltIn func(args ...Value) Value

func (bi BuiltIn) Type() ValueType {
	return BuiltInType
}

func (bi BuiltIn) sealed() {}

var _ Value = (*BuiltIn)(nil)
