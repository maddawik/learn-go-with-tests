package main

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	sumAll := func(acc, x []int) []int {
		return append(acc, Sum(x))
	}
	return Reduce(numbersToSum, sumAll, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbersToSum, sumTail, []int{})
}

func SumAllHeads(numbersToSum ...[]int) int {
	heads := []int{}
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			continue
		}
		heads = append(heads, numbers[0])
	}
	return Sum(heads)
}
