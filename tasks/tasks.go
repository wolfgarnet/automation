package tasks

import (
	"fmt"
	"regexp"
	"log"
	"strconv"
)

func Do() {
	println("Initializing tasks")
}

var timeStringRX *regexp.Regexp

func init() {
	var err error
	timeStringRX, err = regexp.Compile(`(\d+)(\S+)`)
	if err != nil {
		log.Fatal("Unable to compile regexp")
	}
}

func getTextField(config map[string]interface{}, name string) (string, error) {
	field, ok := config[name]
	if !ok {
		return "", fmt.Errorf("The field %v was not found", field)
	}

	text, ok := field.(string)
	if !ok {
		return "", fmt.Errorf("text field, %v, was wrong type", field)
	}

	return text, nil
}

/*
func getUint32Field(config map[string]interface{}, name string) (int64) {
	field, ok := config[name]
	if !ok {
		return 0, fmt.Errorf("The field %v was not found", field)
	}

	text, ok := field.(uint32)
	if !ok {
		return 0, fmt.Errorf("text field, %v, was wrong type", field)
	}

	return
}
*/

func getTimeFromString(number, time string) (int64, error) {
	n, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return -1, err
	}

	switch time {
	case "m":
		return n * 1000 * 60, nil
	case "s":
		return n * 1000, nil
	case "ms":
		return n, nil
	default:
		return 0, nil
	}
}

func getTimeText(config map[string]interface{}, name string) (int64, error) {
	val, err := getTextField(config, name)
	if err != nil {
		return -1, fmt.Errorf("Failed to get time text, %v", err)
	}

	i, err := strconv.ParseInt(val, 10, 64)
	if err == nil {
		log.Printf("VALUE IS %v", i)
		return i, nil
	}

	log.Printf("Time string: %v", val)

	res := timeStringRX.FindAllStringSubmatch(val, -1)
	var ms int64 = 0
	for i, v := range res {
		log.Printf("%v, %v", i, v)
		t, err := getTimeFromString(v[1], v[2])
		if err != nil {
			return 0, err
		}
		ms += t
	}

	log.Printf("MS: %v", ms)

	return ms, nil
}