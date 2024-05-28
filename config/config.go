package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MaxIdleConn  time.Duration
}

var Cfg *Config = NewConfig()

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		ReadTimeout:  GetTimeDurationFromEnv("FHC_READ_TIMEOUT", 5*time.Second),
		WriteTimeout: GetTimeDurationFromEnv("FHC_WRITE_TIMEOUT", 5*time.Second),
		MaxIdleConn:  GetTimeDurationFromEnv("FHC_MAX_IDLE_CONN", 60*time.Second),
	}
}

func GetTimeDurationFromEnv(key string, defaultValue time.Duration) time.Duration {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}

	return duration
}
