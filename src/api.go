package src

import (
	"github.com/captainkovalsky/shop/build/gen"
)
type Profile interface {
	load() (*models.CV, error)
	filterProjectsBy(technologies []string) ([]models.Project, error)
}
