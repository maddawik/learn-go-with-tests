package main

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	result := initialValue
	for _, v := range collection {
		result = f(result, v)
	}
	return result
}

func Find[A any](collection []A, f func(A) bool) (value A, found bool) {
	for _, v := range collection {
		if f(v) {
			return v, true
		}
	}
	return
}
