package main

import (
	"app/config/Router"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type (
	Config struct {
		Server Server
		Redis  Redis
	}

	Server struct {
		Host string `default:"http://localhost"`
		Port string `default:":1234"`
	}

	Redis struct {
		Host string `default:"redis"`
		Port string `default:":6379"`
	}
)

func main() {
	var config Config
	envconfig.Process("", &config)

	log.Println(config)
	router := Router.Init()
	router.Logger.Fatal(router.Start(config.Server.Port))
}
