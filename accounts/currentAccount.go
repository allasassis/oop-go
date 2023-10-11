package accounts

import (
	"OOP/holder"
	"fmt"
)

type CurrentAccount struct {
	Holder        holder.Holder
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (account *CurrentAccount) Withdraw(value float64) string {
	if account.balance < value || value < 0 {
		return "Unfortunately, you cannot make a withdrawal because your balance is lower than your request."
	}

	account.balance -= value
	return "Your withdrawal has been processed successfully! Your new balance is now: " + fmt.Sprintf("%.2f", account.balance)
}

func (account *CurrentAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		account.balance += value
		return "Your deposit has been processed successfully! Your new balance is now: ", account.balance
	}
	return "Unfortunately, you cannot make a deposit because the value is less than zero. Your balance remains unchanged: ", account.balance
}

func (account *CurrentAccount) Transfer(transferValue float64, destinyAccount *CurrentAccount) string {
	if transferValue > 0 && account.balance < transferValue {
		return "Unfortunately, you cannot make a withdrawal because your balance is lower than your request."
	} else {
		account.Withdraw(transferValue)
		destinyAccount.Deposit(transferValue)
		return "Your transfer has been processed successfully! Your new balance is now: " + fmt.Sprintf("%.2f", account.balance)
	}
}

func (c CurrentAccount) GetBalance() float64 {
	return c.balance
}
