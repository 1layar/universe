package appconfig

import "github.com/1layar/universe/pkg/email_service/internal/app/appcontext"

// ConfigSpec is the configuration specification.
type ConfigSpec struct {
	// ServiceListenAddress is the address that the Fiber HTTP server will listen on.
	ServiceListenAddress string `split_words:"true" required:"true" default:":3000"`

	// LogJSONStdout is the flag to enable JSON logging to stdout and disable logging to file.
	LogJSONStdout bool `split_words:"true" required:"true" default:"false"`

	// LogLevel is the log level. Valid values are: trace, debug, info, warn, error, fatal, panic.
	LogLevel ConfigLogLevel `split_words:"true" required:"true" default:"info"`

	// JWT Secret
	JwtSecret  string `split_words:"true" required:"true" default:"test"`
	JwtExpTime string `split_words:"true" required:"true" default:"1h"`

	// redis addr
	RedisAddr string `split_words:"true" required:"true" default:"localhost:6379"`
	// redis username
	RedisUsername string `split_words:"true" required:"false" default:""`
	// redis password
	RedisPassword string `split_words:"true" required:"false" default:""`

	// email delivery
	TypeEmailDelivery string `split_words:"true" required:"true" default:"email_delivery"`

	// Database
	DatabaseUrl string `split_words:"true" required:"true" default:"postgres://postgres:postgres@localhost:5432/postgres"`

	// Redis
	RedisUrl string `split_words:"true" required:"true" default:"redis://localhost:6379"`

	// RabbitMQ
	AmqpUrl string `split_words:"true" required:"true" default:"amqp://guest:guest@localhost:5672/"`
}

type Config struct {
	// ConfigSpec is the configuration specification injected to the config.
	ConfigSpec

	// AppContext is the application context
	AppContext appcontext.Ctx
}
