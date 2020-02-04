package fourthsample

import "fmt"

type Bitcoin int

type Wallet struct{ balance Bitcoin }

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
