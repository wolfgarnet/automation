package system

import (
	"io/ioutil"
	"errors"
	"encoding/json"
	"fmt"
)

type system struct {
	types map[string]interface{}
}

var System system

type NewInstance func() action

func init() {
	println("Initializing system")
	System = system{
		types: make(map[string]interface{}),
	}
}

func (s system) AddType(name string, f interface{}) {
	fmt.Printf("Adding %v\n", name)
	s.types[name] = f
}

func (s system) Read(path string) (map[string]interface{}, error) {
	content, err := ioutil.ReadFile(path)
	if err!=nil{
		return nil, errors.New("Unable to read " + path)
	}
	var conf map[string]interface{}
	err=json.Unmarshal(content, &conf)
	if err!=nil{
		return nil, errors.New("Failed to read json, " + err.Error())
	}

	fmt.Println(conf)

	return conf, nil
}

func (s system) NewTasks(config map[string]interface{}) error {
	tasks, ok := config["tasks"]
	if !ok {
		return errors.New("The field tasks was not found")
	}

	switch vt := tasks.(type) {
		case []interface{}:
		fmt.Printf("vt: %v\n", vt)

		for t, val := range vt {
			fmt.Printf("%v == %v\n", t, val)
			mt, ok := val.(map[string]interface{})
			if ok {
				s.NewTask(mt)
			}
		}
		default:
		return errors.New("tasks field was wrong type")
	}

	return nil
}

func (s system) NewTask(config map[string]interface{}) error {
	ftype, ok := config["type"]
	if !ok {
		return errors.New("The field type was not found")
	}

	vt, ok := ftype.(string)
	if !ok {
		return errors.New("tasks field was wrong type")
	}

	fmt.Printf("vt: %v\n", vt)

	f, err := s.getType(vt)
	if err != nil {
		return err
	}

	f()

	return nil
}

func (s system) getType(name string) (f interface{}, err error) {
	f, ok := s.types[name]
	if ok {
		return
	}

	err = errors.New("Could not find " + name)

	return
}

func (s system) Run(tr *TaskRunner) {
	fmt.Printf("Running %v\n", tr)
}
