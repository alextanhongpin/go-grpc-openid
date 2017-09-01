package main

import (
	"fmt"

	"github.com/nats-io/go-nats"

	"github.com/alextanhongpin/grpc-openid/model"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	c.Subscribe("NewUser", func(usr *model.User) {
		fmt.Printf("received message: %+v", usr)
	})

	for {
	}
}
