package src

import (
	"github.com/captainkovalsky/shop/build/gen"
)

type Profile interface {
	Load() (*models.CV, error)
	filterProjectsBy(technologies []string) ([]models.Project, error)
}

type Cv struct {
}

func (c Cv) Load() (*models.CV, error) {
	return &models.CV{
		Experience: nil,
		Firstname:  "Viktor",
		Surname:    "Dzundza",
		Language: []*models.Language{
			{Name: "English", Level: models.Language_UPPER_INTERMEDIATE},
			{Name: "Ukrainian", Level: models.Language_NATIVE},
			{Name: "Russian", Level: models.Language_FLUENT},
		},
	}, nil
}
