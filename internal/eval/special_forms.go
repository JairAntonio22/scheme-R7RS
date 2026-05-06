package eval

type SpecialForm func(env *Env, args ...Value) Value

var Quote = Symbol("quote")

func QuoteForm(env *Env, args ...Value) Value {
	return args[0]
}

var Define = Symbol("define")

func DefineForm(env *Env, args ...Value) Value {
	sym := args[0]
	val, _ := Eval(args[1], env)
	env.Define(sym, val)
	return Nil{}
}

var If = Symbol("if")

func IfForm(env *Env, args ...Value) Value {
	if args[0] == True {
		val, _ := Eval(args[1], env)
		return val
	}

	val, _ := Eval(args[2], env)
	return val
}

var Lambda = Symbol("lambda")

func LambdaForm(env *Env, args ...Value) Value {
	params := toSlice(args[0])
	body := args[1:]
	child := NewEnv(env)

	return Function{params: params, body: body, env: child}
}
