package main

import (
	"go-nats/bank"
	"go-nats/publisher"
	"go-nats/subscriber"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

var (
	acc           bank.BankAccount
	withdrawalSub subscriber.WithdrawalSubscriber
	withdrawalPub publisher.WithdrawalPublisher
)

const (
	SUBJECT_WITHDRAWALS = "bank.withdrawals"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Printf("Error connecting to nats server %s", err.Error())
	} else {
		log.Printf("Connected to nats successfully: %s", nc.ConnectedUrl())
	}

	acc = bank.New(120)

	withdrawalSub = subscriber.NewWithdrawalSub(nc)
	withdrawalPub = publisher.NewWithdrawalPub(nc, SUBJECT_WITHDRAWALS)

	nc.Subscribe(SUBJECT_WITHDRAWALS, withdrawalSub.Withdraw)

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		withdrawalPub.InitiateWithdrawal(1, 10)
	}

	runtime.Goexit()
}
