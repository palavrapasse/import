package parser

import "github.com/palavrapasse/import/internal/entity"

type LeaksParser interface {
	ParseLeaks() ([]entity.Leak, error)
}
