package service

import (
	"errors"

	"github.com/nasrul21/golang-unit-test/entity"
	"github.com/nasrul21/golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(ID string) (*entity.Category, error) {
	category := service.Repository.FindByID(ID)
	if category != nil {
		return category, nil
	}
	return category, errors.New("Category Not Found")
}
