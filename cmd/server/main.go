package main

import "github.com/marcosduarte-dev/go-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}