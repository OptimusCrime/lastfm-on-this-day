package config

import "os"

type Config struct {
	ApiKey       string
	SharedSecret string
}

func CreateConfig() *Config {
	return &Config{
		ApiKey:       os.Getenv("API_KEY"),
		SharedSecret: os.Getenv("SHARED_SECRET"),
	}
}
