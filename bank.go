package bank

import "fmt"

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
	fmt.Printf("Deposited: %.2f, Balance: %.2f\n", amount, ba.Balance)
}

func (ba *BankAccount) Withdraw(amount float64) {
	if ba.Balance >= amount {
		ba.Balance -= amount
		fmt.Printf("Withdrawn: %.2f, Balance: %.2f\n", amount, ba.Balance)
	} else {
		fmt.Println("Insufficient funds.")
	}
}

func (ba BankAccount) GetBalance() {
	fmt.Printf("Balance: %.2f\n", ba.Balance)
}
