package main

import (
	"fmt"
	"github.com/wolfgarnet/automation/tasks"
	"container/list"
	"github.com/wolfgarnet/automation/system"
)

func main() {
	fmt.Printf("YEAH\n")

	task := &tasks.Display{"test"}
	tks := list.New()
	tks.PushBack(task)

	tr := automation.NewTaskRunner(tks)
	tr.Run()
}
