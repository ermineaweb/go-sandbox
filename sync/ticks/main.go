package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	INITIAL_MESSAGE_COUNTER = 2
	DELAY_MAX_ORIGINAL_CHAT = 1300
)

type MessageSlower struct {
	messageChan    chan string
	displayChan    chan string
	speed          float64
	delay          int
	messageCounter int
}

func main() {
	// speed = msg / time -> time = msg / speed = 1 / speed
	ms := MessageSlower{messageChan: make(chan string), displayChan: make(chan string, 3), speed: 1 / float64(DELAY_MAX_ORIGINAL_CHAT) * 1000, delay: DELAY_MAX_ORIGINAL_CHAT, messageCounter: INITIAL_MESSAGE_COUNTER}

	go ms.send()
	go ms.slow()

	for {
		select {
		case msg := <-ms.displayChan:
			fmt.Printf("Delay:%vms Speed:%.1fmsg/s\n%v\n", ms.delay, ms.speed, msg)
			ms.resetMessageCounter()
		default:
			ms.decreaseMessageCounter()
			if ms.messageCounter == 0 {
				fmt.Println(format(fmt.Sprintf("%v void messages, increase the speed", INITIAL_MESSAGE_COUNTER), "red"))
				ms.faster(50)
			}
		}
		time.Sleep(time.Duration(ms.delay) * time.Millisecond)
	}
}

func (m *MessageSlower) send() {
	var n = 0
	ticker := time.NewTicker(1 * time.Millisecond)

	for {
		n++
		<-ticker.C
		t := <-ticker.C
		message := generateMessage(t.Format("15:04:05"))
		// fmt.Printf("%v: %v\n", format("sender", "yellow"), message)
		m.messageChan <- message
		ticker.Reset(time.Duration(rand.Intn(800)+1) * time.Millisecond)
	}
}

func (m *MessageSlower) slow() {
	for {
		message := <-m.messageChan

		select {
		case m.displayChan <- message:
			// fmt.Printf("%v: %v\n", format("displayer", "magenta"), message)
		default:
			m.slower(200)
			fmt.Println(format("Channel full, messages are lose", "red"))
		}
	}
}

func (m *MessageSlower) slower(n int) {
	if m.delay-n <= 0 {
		m.delay = 1
	} else {
		m.delay -= n
	}
	m.speed = 1 / float64(m.delay) * 1000
}

func (m *MessageSlower) faster(n int) {
	m.delay += n
	m.speed = 1 / float64(m.delay) * 1000
}

func (m *MessageSlower) decreaseMessageCounter() { m.messageCounter-- }

func (m *MessageSlower) resetMessageCounter() { m.messageCounter = INITIAL_MESSAGE_COUNTER }
