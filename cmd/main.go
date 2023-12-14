package main

import (
	"bus/bus"
	"bus/bus/event"
)

func main() {
	bus.New(&event.Event{Name: "1"}, &event.Event{Name: "2"})

}
