package example

import (
	"io/ioutil"
	"path/filepath"

	"github.com/ThatsMrTalbot/example/server"
	"gopkg.in/yaml.v2"
)

// RedisConfig holds config items for mongo
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database int64  `yaml:"database"`
}

// MongoConfig holds config items for mongo
type MongoConfig struct {
	URL      string `yaml:"url"`
	Database string `yaml:"database"`
}

// SendGridConfig holds config items for sendgrid
type SendGridConfig struct {
	Username string `yaml:"username"`
	Key      string `yaml:"key"`
}

// Config holds application configuration
type Config struct {
	Servers  server.Group    `yaml:"servers"`
	Mongo    *MongoConfig    `yaml:"mongo"`
	Redis    *RedisConfig    `yaml:"redis"`
	SendGrid *SendGridConfig `yaml:"sendgrid"`
}

// Open the config and populate the object
func (config *Config) Open() error {
	var (
		data []byte
		err  error
	)

	path := filepath.Join("config", "config.yml")
	if data, err = ioutil.ReadFile(path); err != nil {
		return err
	}

	return yaml.Unmarshal(data, config)
}
