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
	ShareDateSC DateInSeconds
	Context     Context
}

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

func NewLeak(context string, shareDateSC DateInSeconds) (Leak, error) {
	var l Leak

	contextTrim := strings.TrimSpace(context)
	err := checkIfContextConstraintsAreMet(contextTrim)

	if err == nil {
		l = Leak{
			Context: Context(context),
		}
	}

	return l, err
}

func checkIfContextConstraintsAreMet(c string) error {
	size := len(c)

	if size == 0 {
		return errors.New("leak context can not be empty")
	}

	if size > 130 {
		return errors.New("leak context constraints are not met (max 130 characters)")
	}

	return nil
}
