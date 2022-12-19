package internal

import (
	"time"
)

const (
	DateFormatLayout = "2006-01-02T15:04:05.000Z"
)

type DateInSeconds int64 // Epoch time in Seconds

func NewDateSeconds(date string) (DateInSeconds, error) {
	var ds DateInSeconds

	t, err := time.Parse(DateFormatLayout, date)

	if err != nil {
		ds = DateInSeconds(t.Unix())
	}

	return ds, err
}

func (ds DateInSeconds) String() string {
	timeUnix := time.Unix(int64(ds), 0)

	return timeUnix.Format(DateFormatLayout)
}
