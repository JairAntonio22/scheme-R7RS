package scm

var quote = Symbol("quote")

func Quote() Value {
	return &quote
}
