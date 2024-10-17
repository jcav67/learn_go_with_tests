package pointers

import (
	"errors"
	"fmt"
)

// variables para todo el paquete
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// es posible sobreescribir tipos nativos
type Bitcoin int

// re definir como se representa el string
type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
