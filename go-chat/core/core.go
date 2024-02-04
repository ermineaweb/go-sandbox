package core

import (
	"time"
)

type Status string

const (
	MESSAGE Status = "message"
	JOINING Status = "joining"
	LEAVING Status = "leaving"
)

type Broadcaster interface {
	GetEvtChan() chan Emitter
	Add(Emitter)
	Broadcast(Emitter)
	DisplayClients()
	Remove(Emitter)
}

type Emitter interface {
	GetID() int
	SetID(int)
	GetStatus() Status
	GetMessage() string
	Join()
	Leave()
	Send(string)
}

func Run(broadcaster Broadcaster) {
	ticker := time.NewTicker(5 * time.Second)

	// main loop
	for {
		select {
		case <-ticker.C:
			// ping clients
			broadcaster.DisplayClients()

		case evt := <-broadcaster.GetEvtChan():
			switch evt.GetStatus() {

			case MESSAGE:
				broadcaster.Broadcast(evt)

			case JOINING:
				broadcaster.Add(evt)

			case LEAVING:
				broadcaster.Remove(evt)
			}
		default:
		}
	}
}
