package main

import "fmt"

// Account is a interface which defines function definations
type Account interface {
	WithdrawMoney(amount int)
}

// SavingsAccount can be a Account if it implements the functions defined by the interface
type SavingsAccount struct {
	minimumBalance int
	accountbalance int
}

// Function implementation , the function can have it's custom logic
func (s *SavingsAccount) WithdrawMoney(amount int) {
	// For savings account the amount cannot go below minimum balance
	if (s.accountbalance-amount) > 0 && (s.accountbalance-amount) >= s.minimumBalance {
		s.accountbalance -= amount
		fmt.Println("Withdraw successful , remaining balance is ", s.accountbalance)
	} else {
		fmt.Println("Account cannot go below ", s.minimumBalance)
		fmt.Println("Denied withdrawal")
	}
}

// CurrentAccount can be a Account if it implements the functions defined by the interface
type CurrentAccount struct {
	accountbalance int
}

func (c *CurrentAccount) WithdrawMoney(amount int) {
	if c.accountbalance-amount >= 0 {
		c.accountbalance -= amount
		fmt.Println("Withdraw successful , remaining balance is ", c.accountbalance)
	} else {
		fmt.Println("Account cannot be in negative")
		fmt.Println("Denied withdrawal")
	}
}

func main() {

	// Here we can see , We have created a variable of type Account & it can accomodate either Saving or Current Account
	var account1 Account = &SavingsAccount{100, 1000}
	var account2 Account = &CurrentAccount{1000}

	account1.WithdrawMoney(900)
	account2.WithdrawMoney(10000)
}
