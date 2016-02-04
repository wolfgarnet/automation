package tasks

import (
	"github.com/wolfgarnet/automation/system"
	"fmt"
	"github.com/wolfgarnet/typeutils"
	"reflect"
)

type Display struct {
	Text string `json:text`
}

func (d *Display) Execute(tr *automation.TaskRunner, cache automation.Cache) error {
	fmt.Printf("%v\n", d.Text)

	return nil
}

func (d *Display) Finalize(failed bool) error {
	return nil
}

func (d Display) String() string {
	return "Display"
}

func init() {
	//registry.AddFactory("display", NewDisplay)
	typeutils.RegisterType("display", reflect.TypeOf(Display{}))
}