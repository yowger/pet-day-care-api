package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
}

func LoadConfig(configPath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("env")
	v.AddConfigPath(configPath)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found")
		}

		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
