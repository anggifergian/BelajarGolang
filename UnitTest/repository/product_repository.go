package repository

import "GolangUnitTest/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
