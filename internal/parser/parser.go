package parser

import "github.com/palavrapasse/import/internal/entity"

type OnParseErrorCallback func(err error)

type LeakParser interface {
	Parse(ecb ...OnParseErrorCallback) (entity.LeakParse, []error)
}

func processOnParseError(err error, ecb ...OnParseErrorCallback) {
	for _, cb := range ecb {
		cb(err)
	}
}
