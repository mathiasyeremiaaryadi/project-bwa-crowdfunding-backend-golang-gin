package config

import (
	"fmt"
	"log"
	"service-campaign-startup/model/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() *gorm.DB {
	dsn := fmt.Sprintf(
		`%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		CONFIG["MYSQL_USER"],
		CONFIG["MYSQL_HOST"],
		CONFIG["MYSQL_PORT"],
		CONFIG["MYSQL_SCHEMA"],
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("MySQL Database is Connected")

	// err = db.AutoMigrate(&entity.User{})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// err = db.AutoMigrate(&entity.Campaign{})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// err = db.AutoMigrate(&entity.CampaignImage{})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	err = db.AutoMigrate(&entity.Transaction{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("All Schema is Migrated")

	return db
}
