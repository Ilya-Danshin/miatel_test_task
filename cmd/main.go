package main

import (
	"fmt"
	"time"

	"bus/bus"
	"bus/bus/event"
)

type Event struct {
	str string

	ch chan event.Event
}

func (e *Event) Run(ch chan event.Event) {
	e.ch = ch
	for {
		time.Sleep(time.Second * 1)
		e.Signal()
	}
}

func (e *Event) Signal() {
	e.ch <- e
}

func main() {

	e1 := Event{str: "sas"}
	e2 := Event{str: "rar"}

	b := bus.New()
	s1 := b.ConnectToRead()
	s2 := b.ConnectToReadAndWrite()

	b.AddEvent(&e1, s1)
	b.AddEvent(&e1, s2)
	b.AddEvent(&e2, s2)

	for {
		fmt.Println(s1.Read())
		fmt.Println(s2.Read())
	}
}
