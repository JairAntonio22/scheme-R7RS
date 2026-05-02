package eval

type SpecialForm func(env *Env, args ...Value) Value

var specialForms = map[Value]SpecialForm{
	Quote:  QuoteForm,
	Define: DefineForm,
	If:     IfForm,
	Lambda: LambdaForm,
}

var Quote = Symbol("quote")

func QuoteForm(env *Env, args ...Value) Value {
	return args[0]
}

var Define = Symbol("define")

func DefineForm(env *Env, args ...Value) Value {
	env.Define(args[0], args[1])
	return Nil{}
}

var If = Symbol("if")

func IfForm(env *Env, args ...Value) Value {
	if args[0] == True {
		return args[0]
	}

	return args[1]
}

var Lambda = Symbol("lambda")

func LambdaForm(env *Env, args ...Value) Value {
	params := toSlice(args[0])
	body := args[1:]
	child := NewEnv(env)

	return Function{params: params, body: body, env: child}
}
