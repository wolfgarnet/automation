package main

import (
	"flag"
	"os"
	"wolfgarnet/automation/system"
	"wolfgarnet/automation/tasks"
	"log"
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

	tasks, err := system.System.NewTasks(config)
	if err != nil {
		log.Fatalf("Could not create tasks, %v", err)
	}

	tr := system.NewTaskRunner(tasks, false)
	cache := make(map[string]interface{})
	tr.Run(cache)
}
