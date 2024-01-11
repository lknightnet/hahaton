package main

import (
	"bell/config"
	"bell/internal/app"
	"log"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
	}
	app.Run(cfg)
}
