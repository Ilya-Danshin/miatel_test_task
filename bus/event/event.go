package event

// Event is interface for events that can write to bus
type Event interface {
	Run(chan Event)
	Signal()
}
