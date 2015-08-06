package main

import (
	"flag"
	"os"
)

func main() {
	path := flag.String("config", "", "Relative path to configuration")
	flag.Parse()

	config, err := System.read(*path)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	System.NewTasks(config)
}
