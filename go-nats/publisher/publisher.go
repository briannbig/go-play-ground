package publisher

import (
	"go-nats/bank"
	"log"

	"github.com/nats-io/nats.go"
)

type basePublisher struct {
	natsConn *nats.Conn
}

type WithdrawalPublisher struct {
	basePublisher
	subject string
}

func NewWithdrawalPub(natsConn *nats.Conn, subject string) WithdrawalPublisher {
	return WithdrawalPublisher{basePublisher: basePublisher{
		natsConn: natsConn,
	}, subject: subject,
	}
}

func (wp WithdrawalPublisher) InitiateWithdrawal(id uint, amt uint) {
	log.Printf("Publisher initiate withdrawal --- amount: %d", amt)
	req := bank.RequestWithdraw{Id: id, Amount: amt}
	wp.natsConn.Publish(wp.subject, []byte(req.String()))
}
