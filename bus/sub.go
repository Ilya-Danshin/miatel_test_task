package bus

import "bus/bus/event"

// channelSize is size of buffered channel for sub
const channelSize = 1024

// sub is bus subscriber abstraction
type sub struct {
	data chan event.Event

	isAllowToRead  bool
	isAllowToWrite bool
}

// newSub create new bus subscriber
func newSub(readAllow, writeAllow bool) *sub {
	s := &sub{
		isAllowToRead:  readAllow,
		isAllowToWrite: writeAllow,
	}

	s.data = make(chan event.Event, channelSize)

	return s
}

// isAllowedToRead allow to read getter
func (s *sub) isAllowedToRead() bool {
	return s.isAllowToRead
}

// isAllowedToWrite allow to write getter
func (s *sub) isAllowedToWrite() bool {
	return s.isAllowToWrite
}

// Read is method for read events that come from bus
func (s *sub) Read() event.Event {
	return <-s.data
}
