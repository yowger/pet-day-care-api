package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT         string `mapstructure:"PORT"`
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
}

func InitConfig(path string) *Config {
	cfgFile, err := LoadConfig(path)
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	return cfg
}

func LoadConfig(configPath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(".env")
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
