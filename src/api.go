package src

import (
	"github.com/captainkovalsky/shop/src/models"
)

type Profile interface {
	Load() (*models.CV, error)
}

type CvManager struct {
}

func (c CvManager) Load() (*models.CV, error) {
	return &models.CV{}, nil
}
