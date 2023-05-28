package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress    string `mapstructure:"SERVER_ADDRESS, default=localhost:8080"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
}

func NewConfig() (*Config, error) {
	var c Config

	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
