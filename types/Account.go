package types

import "fmt"

type BaseAccount struct {
	owner   string
	balance float64
}

func NewBaseAccount(Owner string, Balance float64) *BaseAccount {
	return &BaseAccount{owner: Owner, balance: Balance}
}

func (b *BaseAccount) GetBalance() float64 {
	return b.balance
}

func (b *BaseAccount) GetOwner() string {
	return b.owner
}

func (b *BaseAccount) Deposit(amount float64) {
	if amount <= 0 {
		return
	}
	b.balance += amount
}

func (b *BaseAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("the amount should be positive")
	}
	if b.balance < amount {
		return fmt.Errorf("insufficient balance")
	}
	b.balance -= amount
	return nil
}

type SavingsAccount struct {
	BaseAccount  //no variable name needed
	interestRate float64
}

func (b *SavingsAccount) GetInterestRate() float64 {
	return b.interestRate
}

func NewSavingsAccount(base BaseAccount, InterestRate float64) *SavingsAccount {
	return &SavingsAccount{BaseAccount: base, interestRate: InterestRate}
}

func (s *SavingsAccount) ApplyInterest() {
	s.balance += s.balance * s.interestRate / 100
}
