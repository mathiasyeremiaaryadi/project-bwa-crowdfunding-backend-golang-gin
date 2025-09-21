package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLoggerConfiguration() *logrus.Logger {
	logger := logrus.New()

	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableColors:    false,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05.999",
	})

	logger.SetLevel(logrus.TraceLevel)

	return logger
}
