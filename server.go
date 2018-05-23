package main

import (
	"github.com/advancing-life/rabbit-can-middleware/config/router"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type (
	// Config ..
	Config struct {
		Server Server
		Redis  Redis
	}

	// Server ...
	Server struct {
		Host string `default:"http://localhost"`
		Port string `default:":1234"`
	}

	// Redis ...
	Redis struct {
		Host string `default:"redis"`
		Port string `default:":6379"`
	}
)

func main() {
	var config Config
	envconfig.Process("", &config)

	log.Println(config)
	router := router.Init()
	router.Logger.Fatal(router.Start(config.Server.Port))
}
