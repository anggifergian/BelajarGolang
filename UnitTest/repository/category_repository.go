package repository

import "GolangUnitTest/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
