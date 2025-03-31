package main

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Joe",
			To:   "Jane",
			Sum:  100,
		},
		{
			From: "Jim",
			To:   "Joe",
			Sum:  25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Jane"), 100.0)
	AssertEqual(t, BalanceFor(transactions, "Joe"), -75.0)
	AssertEqual(t, BalanceFor(transactions, "Jim"), -25.0)
}
