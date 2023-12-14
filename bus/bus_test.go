package bus

import (
	"testing"
	"time"

	"bus/bus/event"
)

// Event interface implementation
type Event struct {
	ch     chan event.Event
	period time.Duration
}

func (e *Event) Run(ch chan event.Event) {
	e.ch = ch
	for i := 0; i < 5; i++ {
		time.Sleep(e.period)
		e.Signal()
	}
}

func (e *Event) Signal() {
	e.ch <- e
}

func TestBus_ConnectToRead(t *testing.T) {
	b := New()
	s := b.ConnectToRead()

	if !s.isAllowedToRead() {
		t.Error("Sub can't read")
	}
	if s.isAllowedToWrite() {
		t.Error("Sub can write")
	}
	if len(b.subs) != 1 {
		t.Errorf("Number of subs of bus - %d; want 1", len(b.subs))
	}
}

func TestBus_ConnectToReadAndWrite(t *testing.T) {
	b := New()
	s := b.ConnectToReadAndWrite()

	if !s.isAllowedToRead() {
		t.Error("Sub can't read")
	}
	if !s.isAllowedToWrite() {
		t.Error("Sub can't write")
	}
	if len(b.subs) != 1 {
		t.Errorf("Number of subs of bus - %d; want 1", len(b.subs))
	}
}

func TestBus_AddEvent(t *testing.T) {
	b := New()
	s1 := b.ConnectToRead()         // Connection that can only read
	s2 := b.ConnectToReadAndWrite() // Connection that can add event to bus

	e := &Event{
		period: time.Millisecond,
	}

	b.AddEvent(e, s1) // Do nothing because s1 can only read
	// Wait some time
	time.Sleep(time.Millisecond * 100)
	// Nothing happens

	// Add event
	b.AddEvent(e, s2)

	// Here got fatal error if something incorrect
	// Read to 1st sub 5 times
	for i := 0; i < 5; i++ {
		s1.Read()
	}

	// Read to 2nd sub 5 times
	for i := 0; i < 5; i++ {
		s2.Read()
	}
}

func TestSub_Read(t *testing.T) {
	b := New()
	s1 := b.ConnectToReadAndWrite() // Connection that can add event to bus
	s2 := b.ConnectToRead()         // Connection that can only read

	e := &Event{
		period: time.Millisecond,
	}

	// Add event
	b.AddEvent(e, s1)

	// Do nothing
	b.AddEvent(e, s2)
	// Read from both 5 times
	for i := 0; i < 5; i++ {
		s1.Read()
		s2.Read()
	}
}
