package pointerror

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		wallet := Wallet{}
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		wallet.Deposit(Bitcoint(15))
		got := wallet.Balance()
		want := Bitcoint(15)

		if got != want {
			t.Errorf("got: %s, but want: %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoint(10)}
		wallet.Withdraw(Bitcoint(5))

		got := wallet.Balance()
		want := Bitcoint(5)

		if got != want {
			t.Errorf("got: %s, but want: %s", got, want)
		}
	})
}
