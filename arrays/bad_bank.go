package arrays

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func applyTransaction(curAccount Account, transaction Transaction) Account {
	if transaction.From == curAccount.Name {
		curAccount.Balance -= transaction.Sum
	}

	if transaction.To == curAccount.Name {
		curAccount.Balance += transaction.Sum
	}

	return curAccount
}
