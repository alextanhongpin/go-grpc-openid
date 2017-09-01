package main

import (
	"fmt"

	"github.com/nats-io/go-nats"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("received message: %s", string(m.Data))
	})

	for {
	}
}
