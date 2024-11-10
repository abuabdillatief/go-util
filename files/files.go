package files

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/abuabdillatief/util/pointer"
)

// ReadFileAsString will read the file at the given path and return its content as a string.
// It returns the string and an error if any.
func ReadFileAsString(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ReadFileAsJson will read the json file at the given path and unmarshal it into the given type.
// It returns the unmarshalled type and an error if any.
func ReadFileAsJson[T any](path string) (*T, error) {
	jsonAsString, err := ReadFileAsString(path)
	if err != nil {
		return nil, err
	}

	var result T
	err = json.Unmarshal([]byte(jsonAsString), &result)
	if err != nil {
		return pointer.FromVal(result), err
	}
	return pointer.FromVal(result), nil
}


// WriteFileAsString will write the given content to the file at the given path.
// It returns an error if any.
func WriteFileAsString(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// WriteJSONFileAsString will marshal the given type into a json string and write it to the file at the given path.
// It returns an error if any.
func WriteJSONFileAsString[T any](path string, content *T) error {
	jsonAsString, err := json.Marshal(content)
	if err != nil {
		return err
	}
	return WriteFileAsString(path, string(jsonAsString))
}

// GetFilesInDir will return all the files in the given directory.
// It returns the files and an error if any.
func GetFilesInDir(path string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return nil, err
	}

	var result []string
	for _, file := range files {
		if info, err := os.Stat(file); err == nil && !info.IsDir() {
			result = append(result, info.Name())
		}
	}

	return result, nil
}
