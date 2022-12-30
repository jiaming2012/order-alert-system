package pubsub

type EventName int

const (
	NewOrderCreated EventName = iota
)

type NewOrderCreatedEvent struct {
	Data string
}

func (ev EventName) String() string {
	switch ev {
	case NewOrderCreated:
		return "NewOrderCreated"
	default:
		return "Undefined"
	}
}
