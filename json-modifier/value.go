package jsonmodifier

// GetValueAtPath will return the value at the given path in the JSON map.
func GetValueAtPath(jsonMap map[string]interface{}, path []string) interface{} {
	current := jsonMap
	for _, key := range path {
		if current[key] == nil {
			return nil
		}
		current = current[key].(map[string]interface{})
	}
	return current
}