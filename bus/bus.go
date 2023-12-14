package bus

import (
	"sync"

	"bus/bus/event"
)

type Bus struct {
	data chan event.Event
	m    sync.Mutex

	queue []event.Event

	subs []*Sub
}

var bus Bus

func init() {
	bus.data = make(chan event.Event)

	go bus.run()
}

func (b *Bus) run() {
	for d := range b.data {
		for _, sub := range b.subs {
			if sub.isAllowedToRead() {
				sub.data <- d
			}
		}
	}
}

func New() *Bus {
	return &bus
}

func (b *Bus) AddEvent(e event.Event, s *Sub) {
	if s.isAllowedToWrite() {
		go e.Run(b.data)
	}
}

func (b *Bus) ConnectToRead() *Sub {
	s := newSub(true, false)

	return b.addSub(s)
}

func (b *Bus) ConnectToWrite() *Sub {
	s := newSub(true, true)

	return b.addSub(s)
}

func (b *Bus) addSub(s *Sub) *Sub {
	b.subs = append(b.subs, s)

	return s
}
