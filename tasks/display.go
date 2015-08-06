package tasks

import (
	"wolfgarnet/automation"
)

type Display struct {

}

func NewDisplay(config map[string]interface{}) *Display {
	d := &Display{

	}

	return d
}

func init() {
	println("HEJ")
	automation.System.AddType("display", NewDisplay)
}