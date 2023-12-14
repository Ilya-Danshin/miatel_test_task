package bus

import (
	"bus/bus/event"
	"bus/bus/subscriber"
)

type Bus struct {
	events []*event.Event
	subs   []*subscriber.Sub
}

func New(events ...*event.Event) *Bus {
	b := &Bus{
		events: events,
	}

	for _, e := range b.events {
		go e.Run()
	}

	return b
}

func (b *Bus) ConnectRead() {

}

func (b *Bus) ConnectReadAndWrite() {

}
