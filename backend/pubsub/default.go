package pubsub

import (
	"github.com/asaskevich/EventBus"
	"github.com/sirupsen/logrus"
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

	logrus.Infof("Subscribed to topic %s", topic)
	return nil
}
