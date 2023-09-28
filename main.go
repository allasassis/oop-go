package main

import "fmt"

type CurrentAccount struct {
	name          string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	account := CurrentAccount{"Allas", 581, 10941, 0}
	account1 := CurrentAccount{name: "John", balance: 41.1}
	fmt.Println(account, account1)
}
