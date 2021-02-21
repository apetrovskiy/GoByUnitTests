package fourth_sample

import (
	"fmt"
	"github.com/cucumber/godog"
)

var wallet Wallet

func beforeScenario(interface{}) {
	wallet = Wallet{}
}

func thereIsInWallet(amount int) error {
	wallet.Deposit(Bitcoin(amount))
	// fmt.Fprintf(os.Stdout, "actual balance (background) %s", wallet.Balance())
	return nil
}

func iDeposit(amount int) error {
	wallet.Deposit(Bitcoin(amount))
	// fmt.Fprintf(os.Stdout, "actual balance (deposit) %s", wallet.Balance())
	return nil
}

func iWithdraw(amount int) error {
	wallet.Withdraw(Bitcoin(amount))
	// fmt.Fprintf(os.Stdout, "actual balance (withdrawal) %s", wallet.Balance())
	return nil
}

func amountIs(expectedTotalAmount int) error {
	// fmt.Fprintf(os.Stdout, "expected %d", expectedTotalAmount)
	// fmt.Fprintf(os.Stdout, "actual %s", wallet.Balance())
	// fmt.Fprintf(os.Stdout, "%v", wallet.Balance() != Bitcoin(expectedTotalAmount))
	if wallet.Balance() != Bitcoin(expectedTotalAmount) {
		return fmt.Errorf("expected balance %s, actual balance %s", Bitcoin(expectedTotalAmount), wallet.Balance())
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.BeforeScenario(beforeScenario)
	s.Step(`^there is (\d+) in wallet$`, thereIsInWallet)
	s.Step(`^I deposit (\d+)$`, iDeposit)
	s.Step(`^I withdraw (\d+)$`, iWithdraw)
	s.Step(`^amount is (\d+)$`, amountIs)
}
