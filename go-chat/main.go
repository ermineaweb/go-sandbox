package main

import (
	"discuss/core"
	"discuss/inmemory"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	var broadcaster core.Broadcaster = &inmemory.InMemoryBroadcaster{
		Emitters: []core.Emitter{}, EvtCh: make(chan core.Emitter),
	}

	var emitters = []core.Emitter{
		&inmemory.InMemoryEmitter{Broadcaster: broadcaster},
		&inmemory.InMemoryEmitter{Broadcaster: broadcaster},
	}

	for index, emitter := range emitters {
		DELAY := time.Duration(4+index*index) * time.Second
		go func(e core.Emitter, delay time.Duration) {
			e.Join()
			time.Sleep(delay)
			e.Send(fmt.Sprintf("Hello all, i am the emitter n°%v", e.GetID()))
			time.Sleep(delay)
			e.Leave()
		}(emitter, DELAY)
	}

	core.Run(broadcaster)
}
