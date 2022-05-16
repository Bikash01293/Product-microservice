package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	Dburl string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/env")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		log.Fatalln("unable to read config file")
		return
	}

	viper.Unmarshal(&c)

	return
}
