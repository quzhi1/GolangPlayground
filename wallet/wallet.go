package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if wallet.balance < amount {
		return ErrInsufficientFunds
	} else {
		wallet.balance -= amount
		return nil
	}
}

// Technically you do not need to change Balance to use a pointer receiver
// as taking a copy of the balance is fine.
// However, by convention you should keep your method receiver types the same for consistency.
func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
