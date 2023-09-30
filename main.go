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

	message := account.withdraw(200)
	fmt.Println(message)
}

func (account *CurrentAccount) withdraw(value float64) string {
	if account.balance < value && value > 0 {
		return "Unfortunately, you cannot make a withdrawal because your balance is lower than your request."
	}

	account.balance -= value
	return "Your withdrawal has been successfully processed! And now your balance is: " + fmt.Sprintf("%.2f", account.balance)
}
