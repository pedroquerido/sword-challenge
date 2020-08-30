package main

import (
	"flag"
	"log"

	"tasks-service/cmd/api/app"
)

var (
	autoMigrate = flag.Bool("auto-migrate", false, "run db migrations on startup")
)

func main() {

	flag.Parse()

	app := app.NewTaskAPI(*autoMigrate)
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
