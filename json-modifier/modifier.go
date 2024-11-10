package jsonmodifier

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/abuabdillatief/go-util/files"
)

// TransformJSONString will modify the field at the given path in the JSON string using the transform function.
// It returns the modified JSON string and an error if any.
func TransformJSONString[O, R any](jsonAsString string, path string, transformFn func(O) R) (string, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonAsString), &jsonMap)
	if err != nil {
		return jsonAsString, err
	}

	updateFieldByPath(jsonMap, strings.Split(path, "."), transformFn)

	bytes, err := json.Marshal(jsonMap)
	if err != nil {
		return jsonAsString, err
	}

	return string(bytes), nil
}

// TransformJSONFile will modify the field at the given path in the JSON file using the transform function.
// It returns an error if any.
func TransformJSONFile[O, R any](filePath, keyPath string, transformFn func(O) R) error {
	jsonAsString, err := files.ReadFileAsString(filePath)
	if err != nil {
		return err
	}

	modifiedJSONAsString, err := TransformJSONString[O, R](jsonAsString, keyPath, transformFn)
	if err != nil {
		return err
	}

	return files.WriteFileAsString(filePath, modifiedJSONAsString)
}

// TransformJSONFilesInDir will modify the field at the given path in all the JSON files in the given directory using the transform function.
// It returns an error if any.
// keyPath is the path to the field to be modified.
// example: "addresses.country"
func TransformJSONFilesInDir[O, R any](dirPath, keyPath string, transformFn func(O) R) error {
	files, err := files.GetFilesInDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := TransformJSONFile[O, R](dirPath+"/"+file, keyPath, transformFn)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateFieldByPath[O, T any](data map[string]interface{}, path []string, transformFn func(O) T) {
	if len(path) == 0 {
		return
	}

	currKey := path[0]
	if currKey == "*" {
		for _, nestedValue := range data {
			if nestedMap, ok := nestedValue.(map[string]interface{}); ok {
				updateFieldByPath(nestedMap, path[1:], transformFn)
			}
		}

		return
	}

	if len(path) == 1 {
		if arr, ok := data[currKey].([]interface{}); ok {
			for i, item := range arr {
				if val, ok := item.(O); ok {
					arr[i] = transformFn(val)
				}
			}

			return
		}

		if value, ok := data[currKey].(O); ok {
			data[currKey] = transformFn(value)
		}
		
		return
	}

	if nextMap, ok := data[currKey].(map[string]interface{}); ok {
		updateFieldByPath(nextMap, path[1:], transformFn)
	} else if nextArray, ok := data[currKey].([]interface{}); ok {
		for _, item := range nextArray {
			if itemMap, ok := item.(map[string]interface{}); ok {
				updateFieldByPath(itemMap, path[1:], transformFn)
			}
		}
	}
}

// CastToType converts a value from one type to another, preserving fields that match by name and type
func CastToType[O any](value interface{}) (O, error) {
	var output O

	// If direct type assertion works, return immediately
	if castedValue, ok := value.(O); ok {
		return castedValue, nil
	}

	// Convert both input and output to reflect.Value
	inputVal := reflect.ValueOf(value)
	outputVal := reflect.ValueOf(&output).Elem()

	// Handle pointer input
	if inputVal.Kind() == reflect.Ptr {
		inputVal = inputVal.Elem()
	}

	// Ensure we're working with structs
	if inputVal.Kind() != reflect.Struct || outputVal.Kind() != reflect.Struct {
		return output, fmt.Errorf("both input and output must be structs")
	}

	outputType := outputVal.Type()

	// Iterate through the fields of the output struct
	for i := 0; i < outputVal.NumField(); i++ {
		outputField := outputVal.Field(i)
		outputFieldType := outputType.Field(i)

		// Find corresponding field in input struct
		inputField := inputVal.FieldByName(outputFieldType.Name)

		// Skip if field doesn't exist in input
		if !inputField.IsValid() {
			continue
		}

		// Only copy if types match exactly
		if inputField.Type() == outputField.Type() && outputField.CanSet() {
			outputField.Set(inputField)
		}
	}

	return output, nil
}
