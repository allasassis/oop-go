package main

import "fmt"

type CurrentAccount struct {
	name          string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	account := CurrentAccount{}
	account.name = "John"
	account.accountNumber = 1230
	account.balance = 500
	account.agencyNumber = 8231

	account2 := CurrentAccount{"Luke", 3212, 321, 1000}

	fmt.Println(account2.transfer(832, &account))
	fmt.Println("Current balance:", account.balance)
}

func (account *CurrentAccount) withdraw(value float64) string {
	if account.balance < value && value > 0 {
		return "Unfortunately, you cannot make a withdrawal because your balance is lower than your request."
	}

	account.balance -= value
	return "Your withdrawal has been processed successfully! Your new balance is now: " + fmt.Sprintf("%.2f", account.balance)
}

func (account *CurrentAccount) deposit(value float64) (string, float64) {
	if value > 0 {
		account.balance += value
		return "Your deposit has been processed successfully! Your new balance is now: ", account.balance
	}
	return "Unfortunately, you cannot make a deposit because the value is less than zero. Your balance remains unchanged: ", account.balance
}

func (account *CurrentAccount) transfer(transferValue float64, destinyAccount *CurrentAccount) string {
	if transferValue > 0 && account.balance < transferValue {
		return "Unfortunately, you cannot make a withdrawal because your balance is lower than your request."
	} else {
		account.withdraw(transferValue)
		destinyAccount.deposit(transferValue)
		return "Your transfer has been processed successfully! Your new balance is now: " + fmt.Sprintf("%.2f", account.balance)
	}
}
