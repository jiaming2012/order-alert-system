package pubsub

import (
	"fmt"
	"github.com/asaskevich/EventBus"
)

var bus EventBus.Bus

func Setup() error {
	bus = EventBus.New()

	return nil
}

func Publish(topic EventName, event interface{}) {
	bus.Publish(topic.String(), event)
}

func Subscribe(topic EventName, callbackFn interface{}) error {
	if err := bus.SubscribeAsync(topic.String(), callbackFn, false); err != nil {
		return err
	}

	fmt.Printf("Subscribed to topic %s\n", topic)
	return nil
}
