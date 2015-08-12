package system

import (
	"container/list"
	"log"
	"fmt"
)

type Task interface {
	Process(cache map[string]interface{}) error
	Conclude(failed bool) error
}

type TaskRunner struct {
	tasks *list.List
	current *list.Element
}

func NewTaskRunner(tasks *list.List) *TaskRunner {
	tr := &TaskRunner{
		tasks: tasks,
		current: nil,
	}

	return tr
}

func (t TaskRunner) Run(cache map[string]interface{}) error {
	log.Printf("Running tasks")
	if t.current == nil {
		t.current = t.tasks.Front()
	} else {
		t.current = t.current.Next()
	}
	task := t.current.Value.(Task)
	if task == nil {
		return fmt.Errorf("Task is wrong type, %v", t.current.Value)
	}

	task.Process()

	return nil
}

func (t TaskRunner) Distribute(task *TaskRunner) error {
	go func() {
		println("DIST")
	}()

	return nil
}

func (t TaskRunner) runNext() error {
	return nil
}
