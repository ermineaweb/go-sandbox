package inmemory

import (
	"discuss/core"
	"fmt"
)

type InMemoryEmitter struct {
	Id          int
	Status      core.Status
	Message     string
	Broadcaster core.Broadcaster
}

func (c *InMemoryEmitter) Join() {
	fmt.Println("new emitter is joining")
	c.Status = core.JOINING
	c.Broadcaster.GetEvtChan() <- c
}

func (c *InMemoryEmitter) Leave() {
	fmt.Printf("emitter %v is leaving\n", c.Id)
	c.Status = core.LEAVING
	c.Broadcaster.GetEvtChan() <- c
}

func (c *InMemoryEmitter) Send(message string) {
	c.Status = core.MESSAGE
	c.Message = message
	c.Broadcaster.GetEvtChan() <- c
}

func (c *InMemoryEmitter) GetID() int             { return c.Id }
func (c *InMemoryEmitter) SetID(id int)           { c.Id = id }
func (c *InMemoryEmitter) GetStatus() core.Status { return c.Status }
func (c *InMemoryEmitter) GetMessage() string     { return c.Message }
