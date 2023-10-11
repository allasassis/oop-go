package accounts

import (
	"OOP/holder"
	"fmt"
)

type SavingsAccount struct {
	Holder                                 holder.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (account *SavingsAccount) Withdraw(value float64) string {
	if account.balance < value || value < 0 {
		return "Unfortunately, you cannot make a withdrawal because your balance is lower than your request."
	}

	account.balance -= value
	return "Your withdrawal has been processed successfully! Your new balance is now: " + fmt.Sprintf("%.2f", account.balance)
}

func (account *SavingsAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		account.balance += value
		return "Your deposit has been processed successfully! Your new balance is now: ", account.balance
	}
	return "Unfortunately, you cannot make a deposit because the value is less than zero. Your balance remains unchanged: ", account.balance
}

func (c SavingsAccount) GetBalance() float64 {
	return c.balance
}
