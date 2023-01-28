package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func NewAliyunConfig() (*AliyunConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, errors.New("Error loading .env file")
	}

	return &AliyunConfig{
		Access: Access{
			AccessKeyId:     os.Getenv("ACCESS_KEY_ID"),
			AccessKeySecret: os.Getenv("ACCESS_KEY_SECRET"),
		},
		Endpoint: os.Getenv("ENDPOINT"),
	}, nil
}

type Access struct {
	AccessKeyId     string
	AccessKeySecret string
}

type AliyunConfig struct {
	Access   Access
	Endpoint string
}
