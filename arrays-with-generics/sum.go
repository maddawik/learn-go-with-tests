package main

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func SumAllHeads(numbersToSum ...[]int) int {
	var sum int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			continue
		}
		sum += numbers[0]
	}

	return sum
}

func Reduce[A any](collection []A, f func(A, A) A, initialValue A) A {
	result := initialValue
	for _, v := range collection {
		result = f(result, v)
	}
	return result
}
