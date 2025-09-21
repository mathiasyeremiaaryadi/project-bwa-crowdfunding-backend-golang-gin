package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewConfigurationEnvironment(logrusLogger *logrus.Logger) {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logrusLogger.Fatalf("Failed to load environment configuration: %v", err)
	}
}
