package main

import (
	"log"

	"github.com/pedroquerido/sword-challenge/notifications-agent/cmd/agent/app"
)

func main() {

	app := app.NewAgent()
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
