package eval

// Value
// Everything that gets evaluated is a Value
// Never use golang's nil as a Value.
type Value interface {
	sealed()
}
