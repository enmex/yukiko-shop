package config

import (
	"log"
	"os"
	"strconv"
)

type ConfigRedis struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func NewConfigRedis() *ConfigRedis {
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &ConfigRedis{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     port,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	}
}
