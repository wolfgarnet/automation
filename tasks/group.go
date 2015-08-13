package tasks
import (
	"wolfgarnet/automation/system"
	"log"
	"fmt"
)

// TimedGroup is repeated sequentially for a given duration
type TimedGroup struct {
	duration int64
}

func (g TimedGroup) Process(cache map[string]interface{}, tr *system.TaskRunner) error {
	log.Printf("---->%v", g.duration)
	return nil
}

func (g TimedGroup) Conclude(failed bool) error {
	return nil
}

// Group is repeated for a number of times
type Group struct {
	number uint32
}

func (g Group) Process(cache map[string]interface{}, tr *system.TaskRunner) error {
	log.Printf("---->%v", g.number)
	return nil
}

func (g Group) Conclude(failed bool) error {
	return nil
}

// TinedIntervalGroup is repeated every interval in parallel for a give n duration
type TimedIntervalGroup struct {
	duration uint32
	interval uint32
}

// IntervalGroup is repeated number of times every interval
type IntervalGroup struct {
	interval uint32
	number uint32
}

func NewGroup(config map[string]interface{}) (system.Task, error) {
	/*
	ms, err := getTimeText(config, "number")
	if err != nil {
		return nil, fmt.Errorf("The field, number, is not a valid time string, %v", err)
	}
	*/

	g := &Group{
		number: 1,
	}

	return g, nil
}

func NewTimedGroup(config map[string]interface{}) (system.Task, error) {
	log.Printf("Creating timed group")

	ms, err := getTimeText(config, "duration")
	if err != nil {
		return nil, fmt.Errorf("The field, number, is not a valid time string, %v", err)
	}
	g := &TimedGroup{
		duration: ms,
	}

	return g, nil
}

func init() {
	system.System.AddType("timedGroup", NewTimedGroup)
}
