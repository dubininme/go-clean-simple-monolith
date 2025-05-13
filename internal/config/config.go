package config

import (
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	User             string `envconfig:"DB_USER"`
	Password         string `envconfig:"DB_PASSWORD"`
	Host             string `envconfig:"DB_HOST"`
	Port             int    `envconfig:"DB_PORT" default:"3306"`
	Name             string `envconfig:"DB_DATABASE"`
	ConnectionsCount int    `envconfig:"DB_CONNECTIONS" default:"10"`
}

type Config struct {
	DB             DBConfig
	JWTSecret      string `envconfig:"JWT_SECRET"`
	ElasticAddress string `envconfig:"ELASTIC_ADDRESS" default:"http://localhost:9200"`
	OutboxPoller   OutboxPollerConfig
}

type OutboxPollerConfig struct {
	BatchSize int `envconfig:"OUTBOX_POLLER_BATCH_SIZE" default:"100"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
