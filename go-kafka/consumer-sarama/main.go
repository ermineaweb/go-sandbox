package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {
	initialOffset := sarama.OffsetNewest
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, config)
	if err != nil {
		return
	}

	partitionList, err := consumer.Partitions("mytopic")
	if err != nil {
		return
	}

	var (
		messages = make(chan *sarama.ConsumerMessage, 256)
		closing  = make(chan struct{})
		wg       sync.WaitGroup
	)

	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition("mytopic", partition, initialOffset)
		if err != nil {
			return
		}

		go func(pc sarama.PartitionConsumer) {
			<-closing
			pc.AsyncClose()
		}(pc)

		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for message := range pc.Messages() {
				messages <- message
			}
		}(pc)
	}

	go func() {
		for msg := range messages {
			fmt.Printf("Partition: %d\nOffset: %d\nKey: %s\nValue: %s\n------------\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		}
	}()

	wg.Wait()
	close(messages)

	if err := consumer.Close(); err != nil {
		fmt.Println("Failed to close consumer: ", err)
	}
}
