package parser

import "github.com/palavrapasse/import/internal/entity"

type LeaksParser interface {
	Parse() ([]entity.Leak, error)
}
