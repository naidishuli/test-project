package utils

import (
	"time"
)

func ParseDateLI(d string) (time.Time, error) {
	if d == "" {
		return time.Time{}, nil
	}
	if len(d) == 4 {
		return time.Parse("2006", d)
	}

	if len(d) > 8 {
		return time.Parse("January 2006", d)
	}

	return time.Parse("Jan 2006", d)
}
