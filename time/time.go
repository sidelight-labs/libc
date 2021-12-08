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
