package controller

import (
	"github.com/1layar/merasa/backend/src/shared/command"
	"github.com/1layar/merasa/backend/src/shared/constant"
)

func RegAddProduct(c ProductHandler) {
	c.RegHandler(RouterItem[command.AddProductCommand, command.AddProductResult]{
		Cmd:     constant.ADD_PRODUCT_CMD,
		Handler: c.HandleAddProduct,
	})
}
