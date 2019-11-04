package main

import (
	"fmt"
	"goj/app"
	"log"
	"os"
)

var (
	Version string = "0.1.0"
	job     string
	src     string
	dst    string
)

func main() {
	parseArgs()
	injectVars()

	err := app.Run(job, src, dst)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error occurred:", err)
		os.Exit(-1)
	}
	os.Exit(0)
}

func parseArgs() {
	args := os.Args[1:]
	if len(args) < 2 || len(args) > 3 {
		if len(args) == 1 && args[0] == "version" {
			fmt.Println(Version)
			os.Exit(0)
		} else {
			printUsage()
			os.Exit(-1)
		}
	}

	job = args[0]
	src = args[1]
	dst = src + "/dist"
	if len(args) == 3 {
		dst = args[2]
	}
	log.Println("Arguments parsed:")
	log.Println(" - job:", job)
	log.Println(" - src:", src)
	log.Println(" - dst:", dst)
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