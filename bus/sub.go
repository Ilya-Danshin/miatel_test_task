package bus

import "bus/bus/event"

const channelSize = 1024

type Sub struct {
	data chan event.Event

	isAllowToRead  bool
	isAllowToWrite bool
}

func newSub(readAllow, writeAllow bool) *Sub {
	s := &Sub{
		isAllowToRead:  readAllow,
		isAllowToWrite: writeAllow,
	}

	s.data = make(chan event.Event, channelSize)

	return s
}

func (s *Sub) isAllowedToRead() bool {
	return s.isAllowToRead
}

func (s *Sub) isAllowedToWrite() bool {
	return s.isAllowToWrite
}

func (s *Sub) Read() event.Event {
	return <-s.data
}
