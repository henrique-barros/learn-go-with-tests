package arrays

import "testing"

func TestBadBank(t *testing.T) {
	var (
		ryia  = Account{Name: "Ryia", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, ryia, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(ryia), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}

func AssertEqual(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.0f, want %.0f", got, want)
	}
}
