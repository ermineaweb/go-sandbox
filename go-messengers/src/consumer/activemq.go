package consumer

import (
	"fmt"
	"log"

	"github.com/go-stomp/stomp"
)

var queues = []string{
	"/queue/RTU_CONFIGURATION_REQ.Q",
	"/queue/RTU_CONFIGURATION_RESP.Q",
	"/topic/OPTICALROUTE_REQ.Q",
	"/topic/TESTSETUP_REQ.Q",
	"/topic/RTU_TEST_RESULT.T",
	"/topic/RTU_NOTIFICATION.T",
	"/topic/RTU_CONFIGURATION_REQ.Q",
	"/topic/ActiveMQ.Advisory.MasterBroker",
}

type Activemq struct {
	Consumer
}

func NewAmqConsumer() *Activemq {
	return &Activemq{}
}

func (c *Activemq) Consume() {
	fmt.Println("activemq consume")
}

func connect() {
	conn, err := stomp.Dial(
		"tcp",
		"localhost:61613",
		stomp.ConnOpt.Login("smx", "smx"),
		stomp.ConnOpt.HeartBeat(500, 500),
		stomp.ConnOpt.HeartBeatGracePeriodMultiplier(5))
	if err != nil {
		log.Fatal("fail to connect to server ", err.Error())
	}
	defer conn.Disconnect()

	messageChan := make(chan *stomp.Message)

	for _, queue := range queues {
		go process(conn, queue, messageChan)
	}

	for {
		message := <-messageChan
		fmt.Println("----------------------------------")
		fmt.Println(string(message.Body))
		fmt.Println("----------------------------------")
		// todo save in a file
	}
}

func process(conn *stomp.Conn, queue string, messageChan chan *stomp.Message) {
	log.Printf("Subscribe: %v\n", queue)

	subscription, err := conn.Subscribe(queue, stomp.AckClient)
	if err != nil {
		println("Cannot subscribe to", queue, err.Error())
		return
	}

	for {
		message := <-subscription.C
		if message.Body != nil {
			log.Printf("Message from the queue: %v\n", queue)
			messageChan <- message
			return
		} else {
			log.Printf("Error consuming message from queue %v\n%v", queue, message.Err.Error())
			return
		}
	}
}
