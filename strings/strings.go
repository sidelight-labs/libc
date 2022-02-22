package strings

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func AminoToHash(input string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}

	sum := sha256.New()
	sum.Write(bytes)

	return fmt.Sprintf("%x", sum.Sum(nil)), nil
}

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
