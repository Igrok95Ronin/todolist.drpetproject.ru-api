package config

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Port string `yaml:"port"`
}

// Глобальная переменная для хранения экземпляра конфигурации
var instance *Config

// Синхронизатор для однократного создания экземпляра конфигурации
var once sync.Once

func GetConfig() *Config {
	logger := logging.GetLogger()
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig("./config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})

	return instance
}
