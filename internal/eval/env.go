package eval

import "maps"

// Env
// bindings are never nil, parent is nil on global.
type Env struct {
	Bindings map[Value]Value
	Parent   *Env
}

func NewEnv(parent *Env) *Env {
	bindings := make(map[Value]Value)
	return &Env{Bindings: bindings, Parent: parent}
}

func DefaultEnv() *Env {
	bindings := maps.Clone(defaultBindings)
	return &Env{Bindings: bindings, Parent: nil}
}

func (e *Env) LookUp(s Value) (Value, bool) {
	curr := e

	for curr != nil {
		value, exists := curr.Bindings[s]

		if exists {
			return value, true
		}

		curr = curr.Parent
	}

	return nil, false
}

func (e *Env) Define(s, v Value) {
	e.Bindings[s] = v
}
