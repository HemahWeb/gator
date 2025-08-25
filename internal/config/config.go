package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBURL           string `mapstructure:"db_url"`
	CurrentUserName string `mapstructure:"current_user_name"`
}

func loadConfig() {
	viper.SetConfigName(".gatorconfig")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../../")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found: %v", err)
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	}
}

func GetConfig() Config {
	loadConfig()
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
	return config
}

func SetUser(user string) {
	loadConfig()
	viper.Set("current_user_name", user)
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Error writing config: %v", err)
	}
}

func GetUser() string {
	loadConfig()
	return viper.GetString("current_user_name")
}
