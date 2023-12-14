package event

type Event interface {
	Run(chan Event)
	Signal()
}
