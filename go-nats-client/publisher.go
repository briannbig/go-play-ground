package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Printf("Could not connect to nats Server: %s ", err)
	} else {
		log.Printf("Connected to nats server. Connection info: %v", nc.ConnectedServerName())
	}

	defer nc.Close()

	nc.Publish("greetings", []byte("hello through nats"))
}
