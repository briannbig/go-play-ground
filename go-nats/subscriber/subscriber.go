package subscriber

import (
	"encoding/json"
	"go-nats/bank"
	"log"

	"github.com/nats-io/nats.go"
)

type baseSubscriber struct {
	natsConn     *nats.Conn
	bankAccounts []bank.BankAccount
}

type WithdrawalSubscriber struct {
	baseSubscriber
}

func NewWithdrawalSub(natsConn *nats.Conn) WithdrawalSubscriber {
	var bankAccounts []bank.BankAccount = make([]bank.BankAccount, 1)
	bankAccounts = append(bankAccounts, bank.New(120))
	return WithdrawalSubscriber{baseSubscriber: baseSubscriber{
		natsConn:     natsConn,
		bankAccounts: bankAccounts,
	}}
}

func (ws WithdrawalSubscriber) Withdraw(m *nats.Msg) {
	var req bank.RequestWithdraw
	json.Unmarshal(m.Data, &req)

	var accFound bool
	for _, acc := range ws.bankAccounts {
		if acc.Id == req.Id {
			accFound = true
			acc.Withdraw(req.Amount)
		}
	}

	if !accFound {
		log.Printf("Account with id: %d not found", req.Id)
	}

}
