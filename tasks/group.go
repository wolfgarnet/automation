package tasks
import (
	"wolfgarnet/automation/system"
	"log"
	"fmt"
)

func init() {
	system.System.AddType("timedGroup", NewTimedGroup)
}

type baseGroup struct {
	subTasks []map[string]interface{}
}

func getBaseGroup(config map[string]interface{}) baseGroup {
	subTasks := config["tasks"].([]map[string]interface{})
	return baseGroup{subTasks}
}

func createAndRun(subTasks []map[string]interface{}) error {
	log.Printf("Creating and running sub tasks")

	list, err := system.System.NewTasksFromArray(subTasks)
	if err != nil {
		return err
	}

	tr := system.NewTaskRunner(list)
	return tr.Run()
}

//
//
// Group is repeated for a number of times
type Group struct {
	baseGroup
	number uint32
}

func NewGroup(config map[string]interface{}) (system.Task, error) {
	g := &Group{
		baseGroup: getBaseGroup(config),
		number: 1,
	}

	return g, nil
}

func (g Group) Process(cache map[string]interface{}, tr *system.TaskRunner) error {
	log.Printf("---->%v", g.number)


	return nil
}

func (g Group) Conclude(failed bool) error {
	return nil
}

func (g Group) String() string {
	return "Group"
}

//
//
// TinedIntervalGroup is repeated every interval in parallel for a give n duration
type TimedIntervalGroup struct {
	duration uint32
	interval uint32
}

//
//
// IntervalGroup is repeated number of times every interval
type IntervalGroup struct {
	interval uint32
	number uint32
}

//
//
// TimedGroup is repeated sequentially for a given duration
type TimedGroup struct {
	duration int64
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

func (g TimedGroup) Process(cache map[string]interface{}, tr *system.TaskRunner) error {
	log.Printf("---->%v", g.duration)
	return nil
}

func (g TimedGroup) Conclude(failed bool) error {
	return nil
}


func (g TimedGroup) String() string {
	return "Timed group"
}
