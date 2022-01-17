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

	return strings.Split(input, ",")
}

func StringToInterface(args []string) []interface{} {
	var result []interface{}

	for _, arg := range args {
		result = append(result, arg)
	}

	return result
}
