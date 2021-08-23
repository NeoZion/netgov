package tools

// Whether the element is in the array
func IsContain(items []interface{}, item interface{}) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
