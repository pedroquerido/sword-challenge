package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/pedroquerido/sword-challenge/service-tasks/cmd/api/app"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln("failed to load env vars")
	}

	app := app.NewTaskAPI()
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
