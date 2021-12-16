package config

import (
	"encoding/json"

	secrets "github.com/ijustfool/docker-secrets"
	"github.com/kelseyhightower/envconfig"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	TelegramApiKey string `envconfig:"TELEAGRAMBOT_KEY" default:""`
	Port           string `envconfig:"PORT" default:"5000"`
}

// GetConfig - создаёт конфиг на основе переменных окружения  префиксом WORKERBOT_ и docker secrets.
func GetConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("TELEGRAM_BOT", cfg)
	if err != nil {
		return nil, err
	}

	err = updateFromDockerSecrets(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func updateFromDockerSecrets(cfg *Config) error {
	dockerSecrets, err := secrets.NewDockerSecrets("")
	if err != nil {
		return nil
	}
	return mapstructure.Decode(dockerSecrets.GetAll(), cfg)
}

// String - генерирует строковое представление конфигурации.
func (c *Config) String() string {
	data, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(data)
}
