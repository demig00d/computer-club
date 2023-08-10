package time24

import (
	"time"
)

const LAYOUT_TIME = "15:04"

type Time struct {
	time.Time
}

func (t Time) Add(duration time.Duration) Time {
	return Time{t.Time.Add(duration)}
}

func (t Time) String() string {
	return t.Format(LAYOUT_TIME)
}

func Parse(s string) (Time, error) {
	t, err := time.Parse(LAYOUT_TIME, s)

	if err != nil {
		return Time{}, err
	}

	return Time{t}, nil
}
