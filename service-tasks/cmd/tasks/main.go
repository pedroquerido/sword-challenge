package main

import (
	"fmt"
	"os"

	"github.com/pedroquerido/sword-challenge/service-tasks/cmd/tasks/app"
)

func main() {

	app := app.NewTasksApp()
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
