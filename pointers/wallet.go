package pointers

import (
	"errors"
	"fmt"
)

const insufficientFundsToWithdrawErrorMessage = "cannot withdraw insufficient funds!"

var ErrInsuficientFunds = errors.New(insufficientFundsToWithdrawErrorMessage)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin { return w.balance }
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsuficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
