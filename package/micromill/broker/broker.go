package broker

import (
	"context"
	"log/slog"
	"sync"

	"github.com/1layar/universe/package/logger"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	mbroker "go-micro.dev/v4/broker"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/cmd"
)

type wmBroker struct {
	opts mbroker.Options
	sync.RWMutex
	running     bool
	pubSub      *gochannel.GoChannel
	subscribers map[string][]*wmSubscriber
	exit        chan chan error
}

// Address implements broker.Broker.
func (b *wmBroker) Address() string {
	return ""
}

// Connect implements broker.Broker.
func (b *wmBroker) Connect() error {
	b.RLock()
	if b.running {
		b.RUnlock()
		return nil
	}
	b.RUnlock()

	b.Lock()
	defer b.Unlock()

	wmLogger := watermill.NewSlogLogger(slog.New(logger.NewMicroLogHandler()))
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		wmLogger,
	)

	b.pubSub = pubSub
	// set running
	b.running = true
	// run the routes
	return nil
}

// Disconnect implements broker.Broker.
func (b *wmBroker) Disconnect() error {
	return b.pubSub.Close()
}

// Init implements broker.Broker.
func (b *wmBroker) Init(opts ...mbroker.Option) error {
	b.RLock()
	if b.running {
		b.RUnlock()
		return nil
	}
	b.RUnlock()

	b.Lock()
	defer b.Unlock()

	b.opts = *mbroker.NewOptions(opts...)

	return nil
}

// Options implements broker.Broker.
func (b *wmBroker) Options() mbroker.Options {
	return b.opts
}

// Publish implements broker.Broker.
func (*wmBroker) Publish(topic string, m *mbroker.Message, opts ...mbroker.PublishOption) error {
	panic("unimplemented")
}

// String implements broker.Broker.
func (b *wmBroker) String() string {
	return "watermill"
}

func (b *wmBroker) unsubscribe(topic string, subscriber *wmSubscriber) {
	b.Lock()
	defer b.Unlock()
	// remove subscriber
	for i, v := range b.subscribers[topic] {
		if v == subscriber {
			b.subscribers[topic] = append(b.subscribers[topic][:i], b.subscribers[topic][i+1:]...)
			break
		}
	}
}

// Subscribe implements broker.Broker.
func (b *wmBroker) Subscribe(topic string, h mbroker.Handler, opts ...mbroker.SubscribeOption) (mbroker.Subscriber, error) {
	options := mbroker.NewSubscribeOptions(opts...)
	msgs, err := b.pubSub.Subscribe(options.Context, topic)

	if err != nil {
		return nil, err
	}

	subscriber := &wmSubscriber{
		msgs:  msgs,
		topic: topic,
		opts:  options,
		fn:    h,
	}

	b.Lock()
	defer b.Unlock()

	b.subscribers[topic] = append(b.subscribers[topic], subscriber)

	return subscriber, nil
}

func init() {
	cmd.DefaultBrokers["watermill"] = NewBroker
}

func NewBroker(opts ...mbroker.Option) mbroker.Broker {
	options := mbroker.Options{
		Context:  context.TODO(),
		Registry: registry.DefaultRegistry,
	}

	for _, o := range opts {
		o(&options)
	}
	return &wmBroker{
		opts: options,
		exit: make(chan chan error),
	}
}
