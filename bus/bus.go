package bus

import (
	"sync"

	"bus/bus/event"
)

type Bus struct {
	data chan event.Event
	m    sync.Mutex

	subs []*sub
}

func (b *Bus) run() {
	for d := range b.data {
		b.m.Lock()
		for _, s := range b.subs {
			s.data <- d
		}
		b.m.Unlock()
	}
}

func New() *Bus {
	b := &Bus{}
	b.data = make(chan event.Event)
	go b.run()

	return b
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
	b.m.Lock()
	b.subs = append(b.subs, s)
	b.m.Unlock()

	return s
}
