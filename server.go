package main

import (
    "app/config/Router"
)

func main() {
    router := Router.Init()
	router.Logger.Fatal(router.Start(":1234"))
}
