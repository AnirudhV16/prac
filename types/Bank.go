package types

import "fmt"

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	GetBalance() float64
}

type BankAccount struct {
	AccountNumber string
	Owner         string
	Balance       float64
}

// constructor
func NewBankAccount(AccountNumber string, Owner string, Balance float64) *BankAccount {
	return &BankAccount{AccountNumber: AccountNumber, Owner: Owner, Balance: Balance}
}

// methods
func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("the amount should be positive")
	}
	if b.Balance < amount {
		return fmt.Errorf("insufficient balance") //no punctuations at the end of errors
	}
	b.Balance -= amount
	return nil
}
func (b *BankAccount) GetBalance() float64 {
	return b.Balance
}
