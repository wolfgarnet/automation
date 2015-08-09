package tasks

import (
	"wolfgarnet/automation/system"
	"fmt"
	"errors"
	"log"
)

type Display struct {
	text string
}

func NewDisplay(config map[string]interface{}) (system.Task, error) {
	tf, ok := config["text"]
	if !ok {
		return nil, errors.New("The field text was not found")
	}

	text, ok := tf.(string)
	if !ok {
		return nil, errors.New("text field was wrong type")
	}

	d := &Display{
		text: text,
	}

	fmt.Printf("New display")

	return d, nil
}

func (d Display) String() string {
	return fmt.Sprintf("Display[%v]", d.text)
}

func init() {
	system.System.AddType("display", NewDisplay)
}

func (d Display) Process(cache map[string]interface{}) error {
	log.Printf(d.text)
	return nil
}