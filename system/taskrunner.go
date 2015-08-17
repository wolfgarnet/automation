package system

import (
	"container/list"
	"log"
	"fmt"
)

type Task interface {
	Process(cache map[string]interface{}, tr *TaskRunner) error
	Conclude(failed bool) error
}

type TaskRunner struct {
	tasks *list.List
	current *list.Element
	cache map[string]interface{}
}

func NewTaskRunner(tasks *list.List) *TaskRunner {
	tr := &TaskRunner{
		tasks: tasks,
		current: nil,
		cache: make(map[string]interface{}),
	}

	return tr
}

func (t *TaskRunner) Run() error {
	log.Printf("Running task for %v", t)

	if t.current == nil {
		t.current = t.tasks.Front()
	} else {
		t.current = t.current.Next()
	}

	log.Printf("Current task is %v", t.current)

	task := t.current.Value.(Task)
	if task == nil {
		return fmt.Errorf("Task is wrong type, %v", t.current.Value)
	}

	return task.Process(t.cache, t)
}

func (t *TaskRunner) EndTask(failed bool) {
	log.Printf("Task ended, %v", t.current)

	// No next, we're done
	if t.current.Next() == nil {
		log.Printf("%v is completely done", t.current)
	} else { // Otherwise run nex
		t.Run()
	}
}

func (t TaskRunner) String() string {
	return "Task runner"
}