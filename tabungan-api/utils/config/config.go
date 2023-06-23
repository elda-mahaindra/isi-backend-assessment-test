package config

import (
	"fmt"

	"tabungan-api/utils/errs"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variables.
type Config struct {
	Port string `mapstructure:"PORT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	const op errs.Op = "config/LoadConfig"

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, errs.E(op, errs.IO, fmt.Sprintf("failed to read configuration file: %s", err))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, errs.E(op, errs.Unanticipated, fmt.Sprintf("failed to unmarshal configuration: %s", err))
	}
	return
}
