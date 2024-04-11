package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Println("Not able to connect to NATS Server: " + nats.DefaultURL)
	} else {
		log.Println("Connected to nc server. Connection info: " + nc.ConnectedAddr())
	}

	nc.Subscribe("greetings", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	runtime.Goexit()

}
