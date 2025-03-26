package pointerror

import "fmt"

type Bitcoint int

type Wallet struct {
	balance Bitcoint
}

func (w *Wallet) Withdraw(amount Bitcoint) {
	w.balance -= amount
}

func (w *Wallet) Deposit(amount Bitcoint) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
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
