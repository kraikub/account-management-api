package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server server
}

type server struct {
	Name string
	Port int
}

func GetRuntimeConfig() (Config, error) {

	var conf Config

	viper.SetConfigName("config")
	viper.AddConfigPath("./internal/config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&conf)
	return conf, err
}
