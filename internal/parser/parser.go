package parser

import "github.com/palavrapasse/import/internal/entity"

type LeakParser interface {
	Parse() (entity.LeakParse, []error)
}
