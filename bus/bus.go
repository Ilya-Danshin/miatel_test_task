package bus

import (
	"sync"

	"bus/bus/event"
)

// Bus is bus abstraction
type Bus struct {
	// data is channel where send signals all events
	data chan event.Event
	// m is mutex for adding subs and send events to subs
	m sync.Mutex

	// subs is slice of all bus subs
	subs []*sub
}

// run is function that always run in bus goroutine.
// Just read from data-channel and send to sub's channels.
func (b *Bus) run() {
	for d := range b.data {
		b.m.Lock()
		for _, s := range b.subs {
			s.data <- d
		}
		b.m.Unlock()
	}
}

// New is creating new bus object and start bus goroutine
func New() *Bus {
	b := &Bus{}
	b.data = make(chan event.Event)
	go b.run()

	return b
}

// AddEvent is method for add mex event to bus
func (b *Bus) AddEvent(e event.Event, s *sub) {
	if s.isAllowedToWrite() {
		go e.Run(b.data)
	}
}

// ConnectToRead is method for subscribe to bus
func (b *Bus) ConnectToRead() *sub {
	s := newSub(true, false)

	return b.addSub(s)
}

// ConnectToReadAndWrite is method for subscribe to bus with allow to add new events
func (b *Bus) ConnectToReadAndWrite() *sub {
	s := newSub(true, true)

	return b.addSub(s)
}

// addSub is internal method for add subs
func (b *Bus) addSub(s *sub) *sub {
	b.m.Lock()
	b.subs = append(b.subs, s)
	b.m.Unlock()

	return s
}
