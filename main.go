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
		log.Printf("ERROR, %v", err)
		os.Exit(1)
	}

	log.Printf("Creating tasks")
	tasks, err := system.System.NewTasksFromConfig(config)
	if err != nil {
		log.Fatalf("Could not create tasks, %v", err)
	}

	log.Printf("Running")
	tr := system.NewTaskRunner(tasks)
	tr.Run()
}
