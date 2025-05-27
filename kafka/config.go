package kafka

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	DefaultEnvPrefix = "kafka"
)

type Config struct {
	BrokerAddrs []string `required:"true"`
}

// ConfigFromEnv generates config from environment variables with default envconfig prefix 'kafka'
func ConfigFromEnv() *Config {
	config := &Config{}
	envconfig.MustProcess(DefaultEnvPrefix, config)
	return config
}

// ConfigFromEnv generates config from environment variables with custom envconfig prefix
func ConfigFromEnvPrefix(prefix string) *Config {
	config := &Config{}
	envconfig.MustProcess(prefix, config)
	return config
}
