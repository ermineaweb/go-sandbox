package consumer

import (
	"log"
)

type IConsumer interface {
	Consume()
}

type Consumer struct {
	log      *log.Logger
	host     string
	login    string
	password string
}
