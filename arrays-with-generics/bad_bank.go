package main

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, transaction := range transactions {
		if transaction.From == name {
			balance -= transaction.Sum
		}
		if transaction.To == name {
			balance += transaction.Sum
		}
	}
	return balance
}
