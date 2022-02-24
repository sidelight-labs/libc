package numbers

import (
	"encoding/json"
	"strconv"
)

func NumberToUint64(number json.Number) (uint64, error) {
	if number.String() == "" {
		return uint64(0), nil
	}

	result, err := strconv.ParseUint(number.String(), 10, 64)
	if err != nil {
		return uint64(0), err
	}

	return result, nil
}
