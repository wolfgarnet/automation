package tasks

import (
	"wolfgarnet/automation/system"
	"fmt"
)

type Display struct {

}

func NewDisplay(config map[string]interface{}) *Display {
	d := &Display{

	}

	fmt.Printf("New display")

	return d
}

func init() {
	system.System.AddType("display", NewDisplay)
}