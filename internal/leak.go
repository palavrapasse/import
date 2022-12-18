package internal

import (
	"errors"
	"strings"

	"go.starlark.net/lib/time"
)

type DateSeconds int64 // Epoch time in Seconds
type context string

type Leak struct {
	LeakId      AutoGenKey
	ShareDateSC DateSeconds
	Context     context
}

func NewLeak(context string) (Leak, error) {
	currentTime := time.Now()

	l := Leak{
		ShareDateSC: currentTime.Unix(),
	}

	err := l.SetContext(context)

	return l, err
}

func (l *Leak) SetContext(c string) error {
	size := len(strings.TrimSpace(c))

	if size == 0 {
		return errors.New("leak context can not be nil or empty")
	}

	if size > 130 {
		return errors.New("leak context exceeds 130 characters")
	}

	l.Context = context(c)

	return nil
}
