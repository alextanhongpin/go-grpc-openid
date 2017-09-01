// Webhook server subscribes to the events that is happening
package main

import (
	"fmt"

	"github.com/nats-io/go-nats"

	"github.com/alextanhongpin/grpc-openid/model"
)

// type Webhook interface {
// }

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	// Subscribe to the creation of new users
	c.Subscribe("NewUser", func(usr *model.User) {
		fmt.Printf("received message: %+v", usr)
	})

	for {
	}
}
