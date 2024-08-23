include .env
export

.PHONY: app_start

app_start:
	go run main.go start
doc:
	swag init --parseDependency -g ./main.go