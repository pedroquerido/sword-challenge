package main

import (
	"log"

	"tasks-service/cmd/api/app"
	/* 	"github.com/joho/godotenv"
	 */)

func main() {

	/* 	if err := godotenv.Load(); err != nil {
		log.Fatalln("failed to load env vars")
	} */

	app := app.NewTaskAPI()
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
