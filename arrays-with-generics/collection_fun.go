package main

func Reduce[A any](collection []A, f func(A, A) A, initialValue A) A {
	result := initialValue
	for _, v := range collection {
		result = f(result, v)
	}
	return result
}
