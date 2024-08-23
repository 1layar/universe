package client

import (
	"context"
	"sync"

	"go-micro.dev/v4/client"
)

type wmClient struct {
	sync.RWMutex
	opts client.Options
}

// Call implements client.Client.
func (c *wmClient) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	panic("unimplemented")
}

// Init implements client.Client.
func (c *wmClient) Init(opts ...client.Option) error {
	c.Lock()
	defer c.Unlock()

	c.opts = client.NewOptions(opts...)

	return nil
}

// NewMessage implements client.Client.
func (*wmClient) NewMessage(topic string, msg interface{}, opts ...client.MessageOption) client.Message {
	panic("unimplemented")
}

// NewRequest implements client.Client.
func (*wmClient) NewRequest(service string, endpoint string, req interface{}, reqOpts ...client.RequestOption) client.Request {
	panic("unimplemented")
}

// Options implements client.Client.
func (c *wmClient) Options() client.Options {
	return c.opts
}

// Publish implements client.Client.
func (*wmClient) Publish(ctx context.Context, msg client.Message, opts ...client.PublishOption) error {
	panic("unimplemented")
}

// Stream implements client.Client.
func (*wmClient) Stream(ctx context.Context, req client.Request, opts ...client.CallOption) (client.Stream, error) {
	panic("unimplemented")
}

// String implements client.Client.
func (c *wmClient) String() string {
	return "wm-client"
}

func NewClient(opts ...client.Option) client.Client {
	return &wmClient{}
}
