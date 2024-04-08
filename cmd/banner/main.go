package main

import (
	"banner_service/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Panic(err.Error())
	}
}
