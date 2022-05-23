package main

import (
	"GO_APP/app"
	"GO_APP/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}