package config

type Config struct {
	Port           string `envconfig:"PORT" default:"5001"`
	TelegramApiKey string `envconfig:"TELEAGRAMBOT_KEY" default:""`
	GiphiKey       string `envconfig:"GIPHY_KEY" default:""`
}
