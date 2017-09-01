package queue

import (
	"log"

	nats "github.com/nats-io/go-nats"
)

// Queue represents the schema for queue
type Queue struct {
	Ref *nats.EncodedConn
}

// New creates a new nats queue reference
func New() *Queue {
	var q Queue
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalf("error connecting to nats: %v", err)
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	q.Ref = c
	return &q
}
