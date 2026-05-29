package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (wallet *Wallet) Deposit(value Bitcoin) {
	wallet.balance += value
}

func (wallet Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (w *Wallet) Withdraw(value Bitcoin) error {
	if w.balance >= value {
		w.balance -= value
		return nil
	}

	return errors.New("oh no")
}
