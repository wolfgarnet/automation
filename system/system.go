package automation

import (
	"container/list"
	"fmt"
)

type Cache map[string]interface{}

type Task interface {
	Execute(tr *TaskRunner, cache Cache) error
	Finalize(failed bool) error
}

// TaskRunner runs a list of tasks sequentially
type TaskRunner struct {
	tasks *list.List
	cache Cache
	current *list.Element
}

func NewTaskRunner(tasks *list.List) *TaskRunner {
	tr := &TaskRunner{
		tasks: tasks,
		cache: make(Cache),
		current: nil,
	}

	return tr
}

func (tr *TaskRunner) AddTask(task Task, last bool) {
	if last {
		tr.tasks.PushBack(task)
	} else {
		tr.tasks.InsertAfter(task, tr.current)
	}
}

func (tr *TaskRunner) Run() error {
	if tr.current == nil {
		tr.current = tr.tasks.Front()
	} else {
		tr.current = tr.current.Next()
	}

	fmt.Printf("Current task is %v", tr.current.Value)

	task, ok := tr.current.Value.(Task)
	if !ok {
		return fmt.Errorf("Task is wrong type, %v", tr.current.Value)
	}

	return task.Execute(tr, tr.cache)
}

