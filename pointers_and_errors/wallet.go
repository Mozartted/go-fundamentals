package pointersanderrors

import "fmt"

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.balance += deposit
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
