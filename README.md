# go-util

`go-util` is a simple collection of utility functions and packages for Go, designed to simplify common tasks and improve code reusability. This package includes utilities for working with slices, JSON, files, pointers, and shell commands.

## Installation

To install `go-util`, use `go get`:

```sh
go get github.com/abuabdillatief/go-util
```

## Modules

### Slices

Utilities for working with slices.

- `GetMax`: Returns the maximum value in a slice.
- `GetMin`: Returns the minimum value in a slice.
- `GetSum`: Returns the sum of values in a slice.
- `GetAverage`: Returns the average of values in a slice.
- `GetMedian`: Returns the median of values in a slice.
- `QuickSort`: Sorts a slice using the quicksort algorithm.

### JSON Modifier

Utilities for modifying JSON strings and files.

- `TransformJSONString`: Modifies a field in a JSON string using a transform function.
- `TransformJSONFile`: Modifies a field in a JSON file using a transform function.
- `TransformJSONFilesInDir`: Modifies a field in all JSON files in a directory using a transform function.
- `TransformJSONFilesInDirIfFileCondition`: Modifies a field in JSON files in a directory based on a condition.
- `GetValueAtPath`: Retrieves the value at a specified path in a JSON map.

### Files

Utilities for working with files.

- `ReadFileAsString`: Reads a file and returns its content as a string.
- `ReadFileAsJson`: Reads a JSON file and unmarshals it into a given type.
- `WriteFileAsString`: Writes a string to a file.
- `WriteJSONFileAsString`: Marshals a type into a JSON string and writes it to a file.
- `GetFilesInDir`: Returns all files in a directory.

### Pointers

Utilities for working with pointers.

- `FromVal`: Returns a pointer to a value.
- `ToVal`: Returns the value at a pointer, or a zero value if the pointer is nil.
- `FromValOrNil`: Returns a pointer if the value is not the zero value, otherwise nil.
- `FromSliceOrNilIfEmpty`: Returns a pointer to a slice or nil if empty.
- `SliceOrNilIfEmpty`: Returns nil if the slice is empty.

### Shell

Utilities for executing shell commands.

- `ExecuteCommand`: Executes a command and returns the output.
- `ExecuteCommandWait`: Executes a command and waits for it to complete, with a timeout.

### Iterator

Utilities for iterating over slices.

- `ForEach`: Applies a function to each element in a slice.
- `Map`: Applies a function to each element in a slice and returns a new slice with the results.
- `Filter`: Applies a function to each element in a slice and returns a new slice with the elements that pass the filter.
- `Reduce`: Applies a function to each element in a slice and returns a single value.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Author

Mohammad Rendra Sura Aditama 
