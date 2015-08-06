package main

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

func init() {
	System = system{
		types: make(map[string]interface{}),
	}
}

func (s system) AddType(name string, f interface{}) {
	fmt.Printf("Adding %v\n", name)
	s.types[name] = f
}

func (s system) read(path string) (map[string]interface{}, error) {
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

	switch vt := ftype.(type) {
		case string:
		fmt.Printf("vt: %v\n", vt)

		for t, val := range vt {
			fmt.Printf("%v == %v\n", t, val)
		}
		default:
		return errors.New("tasks field was wrong type")
	}

	return nil
}
