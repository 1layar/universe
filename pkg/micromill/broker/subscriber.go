package broker

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"go-micro.dev/v4/broker"
	mbroker "go-micro.dev/v4/broker"
)

type wmSubscriber struct {
	opts  broker.SubscribeOptions
	msgs  <-chan *message.Message
	topic string
	fn    mbroker.Handler
	hb    *wmBroker
}

// Options implements broker.Subscriber.
func (s *wmSubscriber) Options() broker.SubscribeOptions {
	return s.opts
}

// Topic implements broker.Subscriber.
func (s *wmSubscriber) Topic() string {
	return s.topic
}

// Unsubscribe implements broker.Subscriber.
func (s *wmSubscriber) Unsubscribe() error {
	s.msgs = nil
	s.hb.unsubscribe(s.topic, s)
	return nil
}
