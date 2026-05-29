package bank

import "fmt"

type BankAccount struct{
	Owner string
	balance float64
}

func NewAccount(owner string, initialBalance float64) BankAccount{
	return BankAccount{
		Owner: owner,
		balance: initialBalance,
	}
}

func (b *BankAccount) Deposit(amount float64){
	if amount > 0{
		b.balance += amount
		b.logTransaction("Deposit",amount)
	}
}

func (b BankAccount) GetBalance() float64{
	return b.balance
}

func (b BankAccount) logTransaction(action string, amount float64){
	fmt.Printf("[LOG] %s of %.2f processed for %s\n ",action,amount,b.Owner)
}