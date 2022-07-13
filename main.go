package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	clihandler "gorski.mateusz/calc/cli"
)

// App is the main structure of CLI application
var App *cli.App

// RunApp runs the cli tool
func RunApp() {
	if err := App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	App = clihandler.CreateCLIHandler()
	RunApp()
}
