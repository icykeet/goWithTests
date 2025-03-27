package pointerror

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoint(15))

		assertBalance(t, wallet, Bitcoint(15))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoint(10)}
		err := wallet.Withdraw(Bitcoint(5))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoint(5))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoint(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoint(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

var assertBalance = func(t testing.TB, wallet Wallet, want Bitcoint) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, but want: %s", got, want)
	}
}

var assertError = func(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

var assertNoError = func(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("get an error, but didn't get one")
	}
}
