package config

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DependencyFacade struct {
	LogrusLogger *logrus.Logger
	MySQLDB      *gorm.DB
}

func NewDepenencies() *DependencyFacade {
	dependencies := &DependencyFacade{}
	dependencies.LogrusLogger = NewLoggerConfiguration()
	NewConfigurationEnvironment(dependencies.LogrusLogger)
	dependencies.MySQLDB = NewMySQLConnection(dependencies.LogrusLogger)

	return dependencies
}
