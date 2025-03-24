package main

func walk(x interface{}, fn func(input string)) {
	fn("I'm as confused as a goat on astroturf")
}
