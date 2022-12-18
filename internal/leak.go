package internal

import (
	"errors"
	"strings"
	"time"
)

const (
	DateFormatLayout = "2006-01-02T15:04:05.000Z"
)

type DateInSeconds int64 // Epoch time in Seconds

type Context string

type Leak struct {
	LeakId      AutoGenKey
	ShareDateSC DateSeconds
	Context     Context
}

func NewDateSeconds(date string) (DateSeconds, err) {
	var ds DateSeconds

	t, err := time.Parse(DateLayout, date)

	if err != nil {
		ds = DateSeconds(t.Unix())
	}
	
	return ds, err

	return ds
}

func (ds DateSeconds) String() string {
	timeUnix := time.Unix(int64(ds), 0)

	return timeUnix.Format(DateLayout)
}

func NewLeak(context string, shareDateSC DateSeconds) (Leak, error) {
	var l Leak

	err := checkIfContextConstraintsAreMet(context)

	if err == nil {
		l = Leak{
			Context: Context(context),
		}
	}

	return l, err
}

func checkIfContextConstraintsAreMet(c string) error {
	size := len(strings.TrimSpace(c))

	if size == 0 {
		return errors.New("leak context can not be nil or empty")
	}

	if size > 130 {
		return errors.New("leak context constraints are not met (max 130 characters)")
	}

	return nil
}
