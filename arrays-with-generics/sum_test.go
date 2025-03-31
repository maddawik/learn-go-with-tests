package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, numbers)
		}
	})

	t.Run("collection of varying size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 6})
		want := []int{0, 10}
		checkSums(t, got, want)
	})
}

func TestSumAllHeads(t *testing.T) {
	checkSums := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("gets the sum of the head of some slices", func(t *testing.T) {
		got := SumAllHeads([]int{1, 2}, []int{5, 6})
		want := 6
		checkSums(t, got, want)
	})

	t.Run("sums empty slices", func(t *testing.T) {
		got := SumAllHeads([]int{}, []int{20, 30})
		want := 20
		checkSums(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		got := Reduce([]int{1, 2, 3}, multiply, 1)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("concatenation of strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		got := Reduce([]string{"1", "2", "3"}, concatenate, "")
		want := "123"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
