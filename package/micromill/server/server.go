package server

import (
	"context"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/1layar/universe/package/logger"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	mlogger "go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
)

type wmServer struct {
	// used for first registration
	registered bool
	started    bool
	sync.RWMutex
	opts   server.Options
	exit   chan chan error
	router *message.Router
}

// Handle implements server.Server.
func (*wmServer) Handle(server.Handler) error {
	panic("unimplemented")
}

// Init implements server.Server.
func (s *wmServer) Init(opts ...server.Option) error {
	options := newOptions(opts...)
	wmLogger := watermill.NewSlogLogger(slog.New(logger.NewMicroLogHandler()))
	router, err := message.NewRouter(message.RouterConfig{}, wmLogger)
	if err != nil {
		return err
	}

	router.AddMiddleware(middleware.Recoverer)

	s.router = router
	s.opts = options

	return nil
}

// NewHandler implements server.Server.
func (*wmServer) NewHandler(interface{}, ...server.HandlerOption) server.Handler {
	panic("unimplemented")
}

// NewSubscriber implements server.Server.
func (*wmServer) NewSubscriber(string, interface{}, ...server.SubscriberOption) server.Subscriber {
	panic("unimplemented")
}

// Options implements server.Server.
func (s *wmServer) Options() server.Options {
	return s.opts
}

func (s *wmServer) getListener() net.Listener {
	if s.opts.Context == nil {
		return nil
	}

	if l, ok := s.opts.Context.Value(netListener{}).(net.Listener); ok && l != nil {
		return l
	}

	return nil
}

// Start implements server.Server.
func (s *wmServer) Start() error {
	s.RLock()
	if s.started {
		s.RUnlock()
		return nil
	}
	s.RUnlock()

	s.opts.Logger.Logf(mlogger.InfoLevel, "Listening watermill")

	go func() {
		// processors are based on router, so they will work when router will start
		if err := s.router.Run(context.Background()); err != nil {
			s.opts.Logger.Logf(mlogger.ErrorLevel, "Server run error: ", err)
		}
	}()

	go func() {
		t := new(time.Ticker)

		// only process if it exists
		if s.opts.RegisterInterval > time.Duration(0) {
			// new ticker
			t = time.NewTicker(s.opts.RegisterInterval)
		}

		// return error chan
		var ch chan error

	Loop:
		for {
			select {
			// register self on interval
			case <-t.C:
				//todo: check if we need to register
			// wait for exit
			case ch = <-s.exit:
				break Loop
			}
		}

		// deregister
		// h.Deregister()

		s.opts.Broker.Disconnect()

		// Solve the problem of early exit
		ch <- s.router.Close()
	}()
	return nil
}

// Stop implements server.Server.
func (s *wmServer) Stop() error {
	ch := make(chan error)
	s.exit <- ch
	return <-ch
}

// String implements server.Server.
func (*wmServer) String() string {
	panic("unimplemented")
}

// Subscribe implements server.Server.
func (*wmServer) Subscribe(server.Subscriber) error {
	panic("unimplemented")
}

func NewServer(opts ...server.Option) server.Server {
	return &wmServer{}
}
