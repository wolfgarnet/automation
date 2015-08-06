package main

import "container/list"

type Task interface {
	Process(cache map[string]interface{})
	IsDistributed() bool
}

type TaskRunner struct {
	tasks *list.List
	isDistributed bool
}

func NewTaskRunner() *TaskRunner {
	tr := &TaskRunner{
		tasks: list.New(),
		isDistributed: false,
	}

	return tr
}

func (t TaskRunner) Run(cache map[string]interface{}) error {
	return nil
}

func (t TaskRunner) runNext() error {
	return nil
}