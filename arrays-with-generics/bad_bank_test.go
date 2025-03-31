package main

import "testing"

func TestBadBank(t *testing.T) {
	joe := Account{Name: "Joe", Balance: 100}
	jane := Account{Name: "Jane", Balance: 75}
	jim := Account{Name: "Jim", Balance: 200}

	transactions := []Transaction{
		NewTransaction(jane, joe, 100),
		NewTransaction(jim, jane, 25),
	}

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(joe), 200.0)
	AssertEqual(t, newBalanceFor(jane), 0.0)
	AssertEqual(t, newBalanceFor(jim), 175.0)
}
