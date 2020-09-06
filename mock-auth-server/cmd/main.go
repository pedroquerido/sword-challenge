package main

import (
	"log"

	"github.com/pedroquerido/sword-challenge/mock-auth-server/cmd/app"
)

func main() {

	app := app.NewAuth()
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
