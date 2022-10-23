package main

import "go-stomp-sniff/src/consumer"

func main() {
	kafka := consumer.NewKafkaConsumer("user")
	amq := consumer.NewAmqConsumer()
	kafka.Consume()
	amq.Consume()
}
