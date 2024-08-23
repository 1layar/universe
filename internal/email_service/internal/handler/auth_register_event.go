package handler

import (
	"context"

	"github.com/1layar/universe/pkg/shared/constant"
)

func RegAuthRegisterEvent(c EmailHandler) {
	messages, err := c.Subscriber.Subscribe(context.Background(), constant.AUTH_REGISTER_EVENT)
	if err != nil {
		panic(err)
	}

	go c.RegEmailHandler(messages)
}
