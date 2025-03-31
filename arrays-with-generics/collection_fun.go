package main

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	result := initialValue
	for _, v := range collection {
		result = f(result, v)
	}
	return result
}
