package main

import (
	"gateway/config"
	"gateway/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	app.Run(cfg)
}
