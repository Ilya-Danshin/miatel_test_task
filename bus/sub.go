package bus

import "bus/bus/event"

const channelSize = 1024

type sub struct {
	data chan event.Event

	isAllowToRead  bool
	isAllowToWrite bool
}

func newSub(readAllow, writeAllow bool) *sub {
	s := &sub{
		isAllowToRead:  readAllow,
		isAllowToWrite: writeAllow,
	}

	s.data = make(chan event.Event, channelSize)

	return s
}

func (s *sub) isAllowedToRead() bool {
	return s.isAllowToRead
}

func (s *sub) isAllowedToWrite() bool {
	return s.isAllowToWrite
}

func (s *sub) Read() event.Event {
	return <-s.data
}
