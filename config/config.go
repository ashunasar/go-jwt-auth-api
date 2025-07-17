package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	DbPath     string `yaml:"db_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

func LoadEnv() Config {

	var configPath string
	flags := flag.String("config", "", "config file path")
	flag.Parse()

	configPath = *flags

	if configPath == "" {
		log.Fatal("Please provide a config file")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {

		log.Fatal("File does not exist, Pleae check it again")

	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {

		log.Fatal("Please pass correct config file")
	}
	return cfg

}
