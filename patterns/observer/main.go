package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	data int
}

type Observer interface {
	NotifyCallback(Event)
}

type Subject interface {
	AddListener(Observer)
	RemoveListener(Observer)
	Notify(Event)
}

type eventObserver struct {
	id   int
	time time.Time
}

type eventSubject struct {
	observers sync.Map
}

func (e eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Observer: %d received: %d after %v\n", e.id, event.data, time.Since(e.time))
}

func (s *eventSubject) AddListener(obs eventObserver) {
	fmt.Printf("Observer: add observer %d to the notify list\n", obs.id)
	obs.time = time.Now()
	s.observers.Store(obs.id, obs)
}

func (s *eventSubject) RemoveListener(obs eventObserver) {
	fmt.Printf("Observer: remove observer %d from the notify list\n", obs.id)
	s.observers.LoadAndDelete(obs.id)
}

func (s *eventSubject) Notify(event Event) {
	s.observers.Range(func(key interface{}, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}
		obs, _ := s.observers.Load(key)
		obs.(Observer).NotifyCallback(event)
		return true
	})
}

func main() {
	observers := eventSubject{observers: sync.Map{}}
	obs1 := eventObserver{id: 1}
	obs2 := eventObserver{id: 2}
	obs3 := eventObserver{id: 3}
	observers.AddListener(obs1)
	observers.AddListener(obs2)
	observers.AddListener(obs3)
	observers.Notify(Event{data: sum(5, 10)})
	observers.Notify(Event{data: sum(15, 2)})
	observers.RemoveListener(obs2)
	observers.Notify(Event{data: sum(5, 2)})
}

func sum(a, b int) int {
	time.Sleep(time.Millisecond * 2000)
	return a + b
}
