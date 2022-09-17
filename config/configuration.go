package config

import (
	"os"
)

const (
	LOCAL = "local"
)

const ENVIRONMENT string = LOCAL

var env = map[string]map[string]string{
	"local": {
		"PORT": "8080",

		"MYSQL_HOST":   "127.0.0.1",
		"MYSQL_PORT":   "3306",
		"MYSQL_USER":   "root",
		"MYSQL_PASS":   "",
		"MYSQL_SCHEMA": "db_food_startup",
	},
}

var CONFIG = env[ENVIRONMENT]

func Getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = Getenv(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}
