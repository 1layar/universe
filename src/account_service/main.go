package main

import "github.com/1layar/universe/src/account_service/cmd/app"

var (
	name    = "account_service"
	version = "1.0.0"
)

func main() {
	app.Run(name, version)
}
