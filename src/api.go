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
		Experience: []*models.Experience{
			{
				StartDate: "02 Jan 06 15:04 -0700",
				EndDate:   "02 Jan 06 15:04 -0700",
				Project:   nil,
				Company:   nil,
			},
		},
		Firstname: "Viktor",
		Surname:   "Dzundza",
		Language: []*models.CV_Language{
			{Name: "English", Level: models.Language_UPPER_INTERMEDIATE.String()},
			{Name: "Ukrainian", Level: models.Language_NATIVE.String()},
			{Name: "Russian", Level: models.Language_FLUENT.String()},
		},
	}, nil
}
