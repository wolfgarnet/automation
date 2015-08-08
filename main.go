package main

import (
	"flag"
	"os"
	"wolfgarnet/automation/system"
	"wolfgarnet/automation/tasks"
)

func main() {
	path := flag.String("config", "", "Relative path to configuration")
	flag.Parse()

	tasks.Do()

	config, err := system.System.Read(*path)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	system.System.NewTasks(config)
}
