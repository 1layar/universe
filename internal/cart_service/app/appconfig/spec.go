package appconfig

import "github.com/1layar/universe/internal/cart_service/app/appcontext"

// ConfigSpec is the configuration specification.
type ConfigSpec struct {
	// ServiceListenAddress is the address that the Fiber HTTP server will listen on.
	ServiceListenAddress string `split_words:"true" required:"true" default:":3000"`

	// LogJSONStdout is the flag to enable JSON logging to stdout and disable logging to file.
	LogJSONStdout bool `split_words:"true" required:"true" default:"false"`

	// LogLevel is the log level. Valid values are: trace, debug, info, warn, error, fatal, panic.
	LogLevel ConfigLogLevel `split_words:"true" required:"true" default:"info"`

	// Database
	DatabaseUrl string `split_words:"true" required:"true" default:"postgres://user:password@localhost:5432/postgres?sslmode=disable"`

	// Redis
	RedisUrl string `split_words:"true" required:"true" default:"redis://localhost:6379"`

	RedisAddr string `split_words:"true" required:"true" default:"localhost:6379"`

	RedisPassword string `split_words:"true" required:"false" default:""`

	// RabbitMQ
	AmqpUrl string `split_words:"true" required:"true" default:"amqp://guest:guest@localhost:5672/"`
}

type Config struct {
	// ConfigSpec is the configuration specification injected to the config.
	ConfigSpec

	// AppContext is the application context
	AppContext appcontext.Ctx
}
