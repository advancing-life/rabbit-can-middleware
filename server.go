package main

import (
	"app/config/Router"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	router := Router.Init()
	router.Logger.Fatal(router.Start(":1234"))
}
