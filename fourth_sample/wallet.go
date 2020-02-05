package fourthsample

import "fmt"

type Bitcoin int

/*type Stringer interface {
	String() string
}*/

/*func (b Bitcoin) String() {
	return fmt.Sprintf("%d BTC", b)
}*/

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct{ balance Bitcoin }

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func (w *Wallet) Withdraw(amount Bitcoin) {

}
