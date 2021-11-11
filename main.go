package main

import (
	"g/go/allsports/app"
	"g/go/allsports/logger"
)

func main() {

	logger.Info("Starting the application ....")

	app.Start()
}
