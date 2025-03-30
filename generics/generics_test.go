package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "hey")
	})
	t.Run("asserting on bools", func(t *testing.T) {
		AssertTrue(t, true)
		AssertFalse(t, false)
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(123)

		AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()

		AssertEqual(t, firstNum+secondNum, 3)
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(Stack[string])

		AssertTrue(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("123")

		AssertFalse(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "456")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "123")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
}
