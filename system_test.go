package main

import (
	"testing"
	"github.com/wolfgarnet/automation/tasks"
	"container/list"
	"github.com/wolfgarnet/automation/system"
)

func TestTaskRunner(t *testing.T) {
	task := tasks.Display{"test"}
	tasks := list.New()
	tasks.PushBack(task)

	tr := automation.NewTaskRunner(tasks)
	tr.Run()
}