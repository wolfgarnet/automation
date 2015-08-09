package system

import (
	"container/list"
	"log"
)

type Task interface {
	Process(cache map[string]interface{}) error
	conclude(failed bool) error
}

type TaskRunner struct {
	tasks *list.List
	isDistributed bool
}

func NewTaskRunner(tasks *list.List, isDistributed bool) *TaskRunner {
	tr := &TaskRunner{
		tasks: tasks,
		isDistributed: isDistributed,
	}

	return tr
}

func (t TaskRunner) Run(cache map[string]interface{}) error {
	log.Printf("Running tasks")
	return nil
}

func (t TaskRunner) Distribute(task *TaskRunner) error {
	go func() {
		println("DIST")
	}()
}

func (t TaskRunner) runNext() error {
	return nil
}
