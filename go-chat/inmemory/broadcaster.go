package inmemory

import (
	"discuss/core"
	"fmt"
)

var id = 1

type InMemoryBroadcaster struct {
	Emitters []core.Emitter
	EvtCh    chan core.Emitter
}

func (b *InMemoryBroadcaster) Add(c core.Emitter) {
	c.SetID(id)
	b.Emitters = append(b.Emitters, c)
	id += 3
	b.DisplayClients()
}

func (b *InMemoryBroadcaster) Broadcast(c core.Emitter) {
	for _, emitter := range b.Emitters {
		if emitter.GetID() != c.GetID() {
			fmt.Printf("emitter %v: %v\n", emitter.GetID(), c.GetMessage())
		}
	}
}

func (b *InMemoryBroadcaster) DisplayClients() {
	if len(b.Emitters) == 0 {
		fmt.Println("No emitter")
	} else {
		fmt.Println("Clients are:")
		for _, emitter := range b.Emitters {
			fmt.Println(emitter.GetID())
		}
	}
}

func (b *InMemoryBroadcaster) Remove(c core.Emitter) {
	for i, emitter := range b.Emitters {
		if emitter.GetID() == c.GetID() {
			b.Emitters[i] = b.Emitters[len(b.Emitters)-1]
			b.Emitters = b.Emitters[:len(b.Emitters)-1]
		}
	}
	b.DisplayClients()
}

func (b *InMemoryBroadcaster) GetEvtChan() chan core.Emitter { return b.EvtCh }
