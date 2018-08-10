package config

import "level-travel/models"

var Config models.Config

func GetConfig() models.Config {
	return Config
}
