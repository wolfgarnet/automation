package tasks

import (
	"wolfgarnet/automation/system"
	"fmt"
	"log"
)

func init() {
	system.System.AddType("wait", NewWait)
}

type Wait struct {
	time int64
}

type RandomWait struct {
	min, max int64
}

func NewWait(config map[string]interface{}) (system.Task, error)  {
	_, ok := config["time"]
	if ok {
		time, err := getTimeText(config, "time")
		if err != nil {
			return nil, err
		}

		return &Wait{time}, nil
	}

	_, okmin := config["min"]
	_, okmax := config["max"]
	if okmin && okmax {
		min, err := getTimeText(config, "min")
		if err != nil {
			return nil, err
		}
		max, err := getTimeText(config, "max")
		if err != nil {
			return nil, err
		}

		return &RandomWait{min, max}, nil
	}

	return nil, fmt.Errorf("Incorrect wait configuration")
}

func (w Wait) Process(cache map[string]interface{}, tr *system.TaskRunner) error {
	log.Printf("Waiting %v ms", w.time)
	tr.EndTask(false)
	return nil
}

func (w Wait) Conclude(failed bool) error {
	return nil
}

func (w Wait) String() string {
	return fmt.Sprintf("Waiting %v ms", w.time)
}


func (w RandomWait) Process(cache map[string]interface{}, tr *system.TaskRunner) error {
	log.Printf("Waiting between %v and %v ms", w.min, w.max)
	tr.EndTask(false)
	return nil
}

func (w RandomWait) Conclude(failed bool) error {
	return nil
}

func (w RandomWait) String() string {
	return fmt.Sprintf("Waiting between %v and %v ms", w.min, w.max)
}
