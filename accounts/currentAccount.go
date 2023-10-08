package accounts

import "fmt"

type CurrentAccount struct {
	Name          string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (account *CurrentAccount) Withdraw(value float64) string {
	if account.Balance < value && value > 0 {
		return "Unfortunately, you cannot make a withdrawal because your balance is lower than your request."
	}

	account.Balance -= value
	return "Your withdrawal has been processed successfully! Your new balance is now: " + fmt.Sprintf("%.2f", account.Balance)
}

func (account *CurrentAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		account.Balance += value
		return "Your deposit has been processed successfully! Your new balance is now: ", account.Balance
	}
	return "Unfortunately, you cannot make a deposit because the value is less than zero. Your balance remains unchanged: ", account.Balance
}

func (account *CurrentAccount) Transfer(transferValue float64, destinyAccount *CurrentAccount) string {
	if transferValue > 0 && account.Balance < transferValue {
		return "Unfortunately, you cannot make a withdrawal because your balance is lower than your request."
	} else {
		account.Withdraw(transferValue)
		destinyAccount.Deposit(transferValue)
		return "Your transfer has been processed successfully! Your new balance is now: " + fmt.Sprintf("%.2f", account.Balance)
	}
}
