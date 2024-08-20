package server

import (
	"crypto/tls"
	"net"

	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/codec"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/transport"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/encoding"
)

type netListener struct{}
type maxMsgSizeKey struct{}
type maxConnKey struct{}
type tlsAuth struct{}

// AuthTLS should be used to setup a secure authentication using TLS.
func AuthTLS(t *tls.Config) server.Option {
	return setServerOption(tlsAuth{}, t)
}

// MaxConn specifies maximum number of max simultaneous connections to server.
func MaxConn(n int) server.Option {
	return setServerOption(maxConnKey{}, n)
}

// Listener specifies the net.Listener to use instead of the default.
func Listener(l net.Listener) server.Option {
	return setServerOption(netListener{}, l)
}

// MaxMsgSize set the maximum message in bytes the server can receive and
// send.  Default maximum message size is 4 MB.
func MaxMsgSize(s int) server.Option {
	return setServerOption(maxMsgSizeKey{}, s)
}

func newOptions(opt ...server.Option) server.Options {
	opts := server.Options{
		Codecs:        make(map[string]codec.NewCodec),
		Metadata:      map[string]string{},
		Broker:        broker.DefaultBroker,
		Registry:      registry.DefaultRegistry,
		RegisterCheck: server.DefaultRegisterCheck,
		Transport:     transport.DefaultTransport,
		Address:       server.DefaultAddress,
		Name:          server.DefaultName,
		Id:            server.DefaultId,
		Version:       server.DefaultVersion,
		Logger:        logger.DefaultLogger,
	}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}
