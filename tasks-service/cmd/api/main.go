package main

import (
	"flag"
	"log"

	"github.com/pedroquerido/sword-challenge/tasks-service/cmd/api/app"
)

var (
	migrateDB = flag.Bool("run-migrations", false, "run db migrations on startup")
)

func main() {

	flag.Parse()

	app := app.NewTaskAPI(*migrateDB)
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
