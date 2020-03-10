package main

import (
	"fmt"
	"goj/app"
	"goj/template"
	"log"
	"os"
)

var (
	Version string = "0.4.3"
	job     string
	path    string
)

func main() {
	parseArgs()
	injectVars()

	err := app.Run(job, path)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error occurred:", err)
		os.Exit(-1)
	}
	os.Exit(0)
}

func parseArgs() {
	args := os.Args[1:]
	if len(args) != 2 {
		if len(args) == 1 && args[0] == "version" {
			fmt.Println("GOJ", Version)
			os.Exit(0)
		} else {
			printUsage()
			os.Exit(-1)
		}
	}

	job = args[0]
	path = args[1]
	log.Println("Arguments parsed:")
	log.Println(" - job:", job)
	log.Println(" - path:", path)
}

func printUsage() {
	_, _ = fmt.Fprintf(os.Stderr, "GOJ v%s: Generate OJ problems with Go\n", Version)
	_, _ = fmt.Fprintln(os.Stderr, "Usage: goj job pathname")
	_, _ = fmt.Fprintln(os.Stderr, "Supported jobs are:")
	_, _ = fmt.Fprintln(os.Stderr, "  - n, new: create a new folder with problem files")
	_, _ = fmt.Fprintln(os.Stderr, "  - g, gen: generate test data and FPS-format file")

}

func injectVars() {
	app.Version = Version
	template.Version = Version
}
