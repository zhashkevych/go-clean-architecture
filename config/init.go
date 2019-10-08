package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port  string
	Auth  auth
	Mongo mongo
}

type auth struct {
	HashSalt   string
	SigningKey string
}

type mongo struct {
	URI string
}

func (c *Config) Set() {
	c.Port = viper.GetString("port")
	c.Auth.HashSalt = viper.GetString("auth.hash_salt")
	c.Auth.SigningKey = viper.GetString("auth.signing_key")
	c.Mongo.URI = viper.GetString("mongo.uri")
}

func Init() (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := new(Config)
	cfg.Set()

	return cfg, nil
}
