package tasks

import (
	"github.com/wolfgarnet/typeutils"
	"reflect"
	"github.com/wolfgarnet/automation/system"
)

// IntervalGroup is repeated number of times every interval
type IntervalGroup struct {
	Interval uint32 `json:interval`
	Number uint32 `json:number`
	taskunner *automation.TaskRunner
	tasks []automation.Task
}

func (d *IntervalGroup) Execute(tr *automation.TaskRunner, cache automation.Cache) error {

	return nil
}

func (d *IntervalGroup) Finalize(failed bool) error {
	return nil
}

func (d IntervalGroup) String() string {
	return "Display"
}

func init() {
	typeutils.RegisterType("intervalgroup", reflect.TypeOf(IntervalGroup{}))
}