package repository

import "github.com/nasrul21/golang-unit-test/entity"

type CategoryRepository interface {
	FindByID(ID string) *entity.Category
}
