package main

import (
	"todo/config"
	"todo/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic("failed to run app")
	}

	app.Run(cfg)
}
