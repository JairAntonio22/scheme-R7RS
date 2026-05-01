package eval

// Env
// bindings are never nil, parent is nil on global.
type Env struct {
	bindings map[Value]Value
	parent   *Env
}

func NewEnv(parent *Env) *Env {
	bindings := make(map[Value]Value)
	return &Env{bindings: bindings, parent: parent}
}

func DefaultEnv() *Env {
	return &Env{bindings: builtIns, parent: nil}
}

func (e *Env) LookUp(s Value) (Value, bool) {
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

func (e *Env) Define(s, v Value) {
	e.bindings[s] = v
}
