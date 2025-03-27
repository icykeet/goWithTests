package pointerror

import (
	"errors"
	"fmt"
)

type Bitcoint int

type Wallet struct {
	balance Bitcoint
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoint) error {
	if amount > w.Balance() {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Deposit(amount Bitcoint) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoint {
	return w.balance
}

type Stringer interface {
	String() string
}

func (b Bitcoint) String() string {
	return fmt.Sprintf("%d BTC", b)
}
