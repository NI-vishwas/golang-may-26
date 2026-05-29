package main

import (
	"fmt"
	"projectPackage/bank"
)

func main(){
	account := bank.NewAccount("Vishwas", 500.00)

	fmt.Println("Account Number:", account.Owner)

	account.Deposit(500.00)
	
	fmt.Printf("Current Balance: %.2f\n",account.GetBalance())

	// This part if enabled will give error as balance is private
	// account.balance = 1000000000.00
}