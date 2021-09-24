package main

import (
	"github.com/killtheverse/nitd-results/app"
	"github.com/killtheverse/nitd-results/config"
)

func main() {
	config := config.NewConfig()
	app.ConfigAndRun(config)
}