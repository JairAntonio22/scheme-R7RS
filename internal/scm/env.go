package scm

// Env
// bindings are never nil, parent is nil on global.
type Env struct {
	bindings map[Symbol]Value
	parent   *Env
}

func NewEnv(parent *Env) *Env {
	bindings := make(map[Symbol]Value)
	return &Env{bindings: bindings, parent: parent}
}

func (e *Env) LookUp(s Symbol) (Value, bool) {
	curr := e

	for curr != nil {
		value, exists := curr.bindings[s]

		if exists {
			return value, true
		}

		curr = curr.parent
	}

	return nil, false
}

func (e *Env) Define(s Symbol, v Value) {
	e.bindings[s] = v
}
