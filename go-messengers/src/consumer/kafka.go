package consumer

import (
	"fmt"
)

type Kafka struct {
	Consumer
	id string
}

func NewKafkaConsumer(host string) *Kafka {
	return &Kafka{
		Consumer: Consumer{
			host: host, login: "", password: "",
		},
	}
}

func (k *Kafka) Consume() {
	fmt.Println("kafka consume " + k.host)
}
