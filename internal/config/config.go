package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/paveltyukin/practicum-go-shortener/pkg/logger"
)

type Config struct {
	Host string `env:"HOST"`
	Port string `env:"PORT"`
}

var instance *Config
var once sync.Once

func InitConfig(logger *logger.Logger) *Config {
	once.Do(func() {
		logger.Info("read application configuration")
		instance = &Config{}
		err := godotenv.Load(".env")
		if err != nil {
			logger.Info("failed to load env from godotenv")
			logger.Fatal(err)
		}

		if err = cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})

	return instance
}
