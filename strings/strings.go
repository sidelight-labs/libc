package strings

import "strings"

func Shorten(input string, size int) string {
	if len(input) > size {
		return input[:size]
	}
	return input
}

func StringToArray(input string) []string {
	if input == "" {
		return []string{}
	}

	result := []string{input}

	if strings.Contains(input, ",") {
		return strings.Split(input, ",")
	}

	return result
}

func StringToInterface(args []string) []interface{} {
	var result []interface{}

	for _, arg := range args {
		result = append(result, arg)
	}

	return result
}
