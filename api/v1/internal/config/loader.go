package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Server server
	Db     db
}

type db struct {
	Mongo mongo
}

type server struct {
	Name string
	Port int
}

type mongo struct {
	Uri string
	Name string
}

func getRuntimeConfigFromEnv() (Config, error) {

	// Parse KRAIKUB_SERVER_PORT from string to int.
	port, err := strconv.Atoi(os.Getenv("KRAIKUB_SERVER_PORT"))
	if err != nil {
		return Config{}, err
	}

	// Create server config
	s := server{
		Name: os.Getenv("KRAIKUB_SERVER_NAME"),
		Port: port,
	}
	return Config{
		Server: s,
	}, nil
}

func getRuntimeConfigFromYamlFile() (Config, error) {
	var conf Config
	viper.SetConfigName("config")
	viper.AddConfigPath("./api/v1/internal/config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&conf)
	return conf, err
}

func GetRuntimeConfig() (Config, error) {

	if os.Getenv("KRAIKUB_ENV") == "production" {
		return getRuntimeConfigFromEnv()
	}
	return getRuntimeConfigFromYamlFile()
}
