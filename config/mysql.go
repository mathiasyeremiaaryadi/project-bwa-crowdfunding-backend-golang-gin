package config

import (
	"fmt"
	"log"
	"os"
	"service-campaign-startup/model/entity"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySQLConnection(logrusLogger *logrus.Logger) *gorm.DB {
	logrusLogger.Info("Establishing MySQL connection . . .")

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, viper.GetString("MYSQL_USERNAME"), viper.GetString("MYSQL_PASSWORD"), viper.GetString("MYSQL_HOST"), viper.GetString("MYSQL_PORT"), viper.GetString("MYSQL_DATABASE"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	MySQLDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		logrusLogger.Fatalf("Failed to establish MySQL conenction: %+v", err)
	}

	logrusLogger.Info("MySQL connection successfull")

	logrusLogger.Info("Migrating MySQL tables . . .")
	err = MySQLDB.Debug().AutoMigrate(&entity.User{})
	if err != nil {
		logrusLogger.Fatalf("Failed to migrate MySQL table User: %+v", err)
	}
	logrusLogger.Info("Migrated User table successfully")

	err = MySQLDB.Debug().AutoMigrate(&entity.Campaign{})
	if err != nil {
		logrusLogger.Fatalf("Failed to migrate MySQL table Campaign: %+v", err)
	}
	logrusLogger.Info("Migrated Campign table successfully")

	err = MySQLDB.Debug().AutoMigrate(&entity.CampaignImage{})
	if err != nil {
		logrusLogger.Fatalf("Failed to migrate MySQL table CampaignImage: %+v", err)
	}
	logrusLogger.Info("Migrated CampignImage table successfully")

	err = MySQLDB.Debug().AutoMigrate(&entity.Transaction{})
	if err != nil {
		logrusLogger.Fatalf("Failed to migrate MySQL table Transaction: %+v", err)
	}
	logrusLogger.Info("Migrated Transaction table successfully")

	logrusLogger.Info("All tables are migrated successfully")

	return MySQLDB
}
