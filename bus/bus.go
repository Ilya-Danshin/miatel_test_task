package bus

import (
	"sync"

	"bus/bus/event"
)

type Bus struct {
	data chan event.Event
	m    sync.Mutex

	queue []event.Event

	subs []*sub
}

var bus Bus

func init() {
	bus.data = make(chan event.Event)

	go bus.run()
}

func (b *Bus) run() {
	for d := range b.data {
		for _, s := range b.subs {
			if s.isAllowedToRead() {
				s.data <- d
			}
		}
	}
}

func New() *Bus {
	return &bus
}

func (b *Bus) AddEvent(e event.Event, s *sub) {
	if s.isAllowedToWrite() {
		go e.Run(b.data)
	}
}

func (b *Bus) ConnectToRead() *sub {
	s := newSub(true, false)

	return b.addSub(s)
}

func (b *Bus) ConnectToReadAndWrite() *sub {
	s := newSub(true, true)

	return b.addSub(s)
}

func (b *Bus) addSub(s *sub) *sub {
	b.subs = append(b.subs, s)

	return s
}
