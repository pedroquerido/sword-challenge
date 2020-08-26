package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/pedroquerido/sword-challenge/service-tasks/cmd/task/app"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln("failed to load env vars")
	}

	app := app.NewTaskApp()
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
