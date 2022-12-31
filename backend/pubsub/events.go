package pubsub

type EventName int

const (
	OrderCreated EventName = iota
	OrderUpdated
)

type OrderEvent struct{}

func (ev EventName) String() string {
	switch ev {
	case OrderCreated:
		return "OrderCreated"
	case OrderUpdated:
		return "OrderUpdated"
	default:
		return "Undefined"
	}
}
