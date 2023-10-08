package main

import (
	"fmt"

	acc "OOP/accounts"
)

func main() {
	account := acc.CurrentAccount{}
	account.Name = "John"
	account.AccountNumber = 1230
	account.Balance = 500
	account.AgencyNumber = 8231

	account2 := acc.CurrentAccount{"Luke", 3212, 321, 1000}

	fmt.Println(account2.Transfer(832, &account))
	fmt.Println("Current balance:", account.Balance)
}
