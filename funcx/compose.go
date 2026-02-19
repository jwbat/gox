package funcx


// Pipe applies a sequence of functions A -> A
func Pipe[A any](funcs ...func(A) A) func(A) A {
	return func(a A) A {
		for _, f := range funcs {
			a = f(a)
		}
		return a
	}
}


// Compose takes functions f & g that returns function h such that h(x) = f(g(x)).
func Compose[A, B, C any](f func(B) C, g func(A) B) func(A) C {
	return func(a A) C {
		return f(g(a))
	}
}


func Identity[T any](t T) T { 
	return t 
}
