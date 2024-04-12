package bank

import (
	"encoding/json"
	"log"
)

type BankAccount struct {
	Id      uint `json:"id"`
	Balance uint `json:"balance"`
}

func (b BankAccount) GetBalance() uint {
	return b.Balance
}

func (b BankAccount) Withdraw(amount uint) uint {
	b.Balance -= amount
	log.Printf("Withdrew %d, new balance: %d", amount, b.Balance)
	return b.Balance
}
func (b BankAccount) Deposit(amount uint) uint {
	b.Balance += amount
	return b.Balance
}
func New(amount uint) BankAccount {

	return BankAccount{Id: 1, Balance: amount}
}

func (bA BankAccount) String() string {
	out, err := json.MarshalIndent(bA, "", "   ")
	if err != nil {
		log.Printf("Could not mashal data --- error: %s", err.Error())
		return ""
	}
	return string(out)

}
