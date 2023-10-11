package main

import (
	"fmt"

	acc "OOP/accounts"
	hol "OOP/holder"
)

func main() {
	holder := hol.Holder{Name: "John", Identity: 13292319, Profession: "Teacher"}
	account := acc.CurrentAccount{Holder: holder, AgencyNumber: 12133, AccountNumber: 123213}

	fmt.Println("Name:", account.Holder.Name)
	fmt.Println("John's balance is:", account.GetBalance())

	holder1 := hol.Holder{Name: "Luke", Identity: 4231321, Profession: "Engineer"}
	account1 := acc.SavingsAccount{Holder: holder1, AgencyNumber: 23112, AccountNumber: 412321}
	account1.Deposit(1000)
	account1.Withdraw(-1322)
	fmt.Println("Name:", account1.Holder.Name)
	fmt.Println("Luke's balance is:", account1.GetBalance())

	PayBill(&account1, 60)
	fmt.Println("Luke's balance after paying the bill is:", account1.GetBalance())
}

func PayBill(account VerifyAcccount, billValue float64) {
	account.Withdraw(billValue)
}

type VerifyAcccount interface {
	Withdraw(value float64) string
}
