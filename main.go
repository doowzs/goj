package main

import (
	"fmt"
	"goj/app"
	"log"
	"os"
)

var (
	Version string = "0.1.1"
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
			fmt.Println(Version)
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
	_, _ = fmt.Fprintln(os.Stderr, "GOJ", Version, ": Generate OJ problems with Go")
	_, _ = fmt.Fprintln(os.Stderr, "Usage: goj job src [dst]")
	_, _ = fmt.Fprintln(os.Stderr, "\nGOJ supports two types of jobs:")
	_, _ = fmt.Fprintln(os.Stderr, "  - n, new: a new problem template")
	_, _ = fmt.Fprintln(os.Stderr, "  - g, gen: generate hust-oj files")

}

func injectVars() {
	app.Version = Version
}