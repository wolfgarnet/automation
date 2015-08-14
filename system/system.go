package system

import (
	"io/ioutil"
	"errors"
	"encoding/json"
	"fmt"
	"container/list"
	"log"
)

type system struct {
	types map[string]NewInstance
}

var System system

type NewInstance func(config map[string]interface{}) (Task, error)

func init() {
	println("Initializing system")
	System = system{
		types: make(map[string]NewInstance),
	}
}

func (s system) AddType(name string, f NewInstance) {
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

func (s system) NewTasksFromConfig(config map[string]interface{}) (*list.List, error) {
	tasks, ok := config["tasks"]
	if !ok {
		return nil, errors.New("The field tasks was not found")
	}

	switch vt := tasks.(type) {
		case []map[string]interface{}:
		fmt.Printf("vt: %v\n", vt)
		/*
		mt, ok := vt.([]map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Incorrect type")
		}
		*/
		return s.NewTasksFromArray(vt)
		default:
		return nil, errors.New("tasks field was wrong type")
	}

	return nil, nil
}

func (s system) NewTasksFromArray(array []map[string]interface{}) (*list.List, error) {
	tasks := new(list.List)

	for t, val := range array {
		fmt.Printf("%v == %v\n", t, val)

		task, err := s.NewTask(val)
		if err == nil {
			log.Printf("Adding %v", task)
			tasks.PushBack(task)
		} else {
			log.Printf("Could not add %v, %v", val, err)
		}
	}

	return tasks, nil
}

func (s system) NewTask(config map[string]interface{}) (Task, error) {
	ftype, ok := config["type"]
	if !ok {
		return nil, errors.New("The field type was not found")
	}

	vt, ok := ftype.(string)
	if !ok {
		return nil, errors.New("tasks field was wrong type")
	}

	fmt.Printf("vt: %v\n", vt)

	f, err := s.getType(vt)
	if err != nil {
		return nil, err
	}

	task, err := f(config)
	if err != nil {
		log.Fatalf("Unable to create task, %v", err)
	}

	return task, nil
}

func (s system) getType(name string) (f NewInstance, err error) {
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
