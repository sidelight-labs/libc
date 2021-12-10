package time

import (
	"fmt"
	"time"
)

const TimeParsingError = "non RFC3339 or YYYY-MM-DD date format: %s"

func TimeToEpoch(timeString string) (int, error) {
	var (
		parsed time.Time
		err    error
	)

	if parsed, err = time.Parse("2006-01-02", timeString); err != nil {
		parsed, err = time.Parse(time.RFC3339, timeString)
		if err != nil {
			return 0, fmt.Errorf(TimeParsingError, timeString)
		}
	}

	return int(parsed.Unix()), nil
}

const (
	DateFormat  = "2006-01-02"
	InvalidTime = "non RFC3339 or YYYY-MM-DD date format: %s"
)

func FormatTime(timeString string) (string, error) {
	var (
		parsed time.Time
		err    error
	)

	if parsed, err = time.Parse("2006-01-02", timeString); err != nil {
		parsed, err = time.Parse(time.RFC3339, timeString)
		if err != nil {
			return "", fmt.Errorf(InvalidTime, timeString)
		}
	}

	return parsed.Format("Jan 2 2006"), nil
}

func TimestampToDate(input string) (string, error) {
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return "", err
	}

	return t.Format(DateFormat), nil
}
