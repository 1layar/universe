package config

import (
	"fmt"

	"github.com/pkg/errors"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
)

type Config struct {
	Port     int
	Tracing  TracingConfig
	Database DatabaseConfig
}

type TracingConfig struct {
	Enable bool
	Jaeger JaegerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type JaegerConfig struct {
	URL string
}

var cfg *Config = &Config{
	Port: 3550,
}

func Address() string {
	return fmt.Sprintf(":%d", cfg.Port)
}

func Tracing() TracingConfig {
	return cfg.Tracing
}

func Database() DatabaseConfig {
	return cfg.Database
}

func MysqlUrl() (string, error) {
	databaseConfig := Database()
	dbUser := databaseConfig.User
	dbPass := databaseConfig.Password
	dbName := databaseConfig.Name
	dbHost := databaseConfig.Host
	dbPort := databaseConfig.Port

	if dbPort == 0 {
		dbPort = 3306
	}

	if !(dbUser != "" && dbPass != "" && dbName != "" && dbHost != "") {
		return "", errors.New("configor.MysqlUrl: missing database config")
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	return url, nil
}

func Load() error {
	configor, err := config.NewConfig(config.WithSource(env.NewSource()))
	if err != nil {
		return errors.Wrap(err, "configor.New")
	}
	if err := configor.Load(); err != nil {
		return errors.Wrap(err, "configor.Load")
	}
	if err := configor.Scan(cfg); err != nil {
		return errors.Wrap(err, "configor.Scan")
	}
	return nil
}
