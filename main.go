package main

import (
	"fmt"

	acc "OOP/accounts"
	"OOP/holder"
)

func main() {
	holder := holder.Holder{Name: "John", Identity: 13292319, Profession: "Teacher"}
	account := acc.CurrentAccount{Holder: holder, AgencyNumber: 12133, AccountNumber: 123213}

	fmt.Println("Name:", account.Holder.Name)
	fmt.Println("John's balance is:", account.ReturnBalance())
}
