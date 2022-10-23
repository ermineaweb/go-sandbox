package main

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic         = "mytopic"
	brokerAddress = "localhost:9092"
)

func main() {
	ctx := context.Background()
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "mygroup",
	})
	defer reader.Close()
	fmt.Println("start consuming")
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		// after receiving the message, log its value
		fmt.Printf("Partition: %d\nOffset: %d\nKey: %s\nValue: %s\n------------\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
}
