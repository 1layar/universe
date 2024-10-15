package appconfig

import "github.com/1layar/universe/internal/auth_service/app/appcontext"

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

	// OpenFga
	FgaApiUrl string `split_words:"true" required:"true" default:"localhost:8080"`
	FgaStore  string `split_words:"true" required:"true" default:"default"`

	// RabbitMQ
	AmqpUrl string `split_words:"true" required:"true" default:"localhost:5672"`

	// Database
	DatabaseUrl string `split_words:"true" required:"true" default:"postgres://postgres:postgres@localhost:5432/postgres"`
}

type Config struct {
	// ConfigSpec is the configuration specification injected to the config.
	ConfigSpec

	// AppContext is the application context
	AppContext appcontext.Ctx
}
