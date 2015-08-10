package tasks

import (
	"fmt"
)

func Do() {
	println("Initializing tasks")
}

func getTextField(config map[string]interface{}, field string) {
	field, ok := config[field]
	if !ok {
		return "", fmt.Errorf("The field %v was not found", field)
	}

	text, ok := field.(string)
	if !ok {
		return "", fmt.Errorf("text field, %v, was wrong type", field)
	}

	return text
}


func getUint32Field(config map[string]interface{}, field string) {
	field, ok := config[field]
	if !ok {
		return "", fmt.Errorf("The field %v was not found", field)
	}

	text, ok := field.(uint32)
	if !ok {
		return "", fmt.Errorf("text field, %v, was wrong type", field)
	}

	return text
}

func getTimeText(config map[string]interface{}, field string) (int64, error) {
	val, err := getTextField(config, field)
	if err != nil {
		return -1, fmt.Errorf("Failed to get time text, %v", err)
	}

	
}