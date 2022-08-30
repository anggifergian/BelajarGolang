package service

import (
	"GolangUnitTest/entity"
	"GolangUnitTest/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (ps ProductService) Get(id string) (*entity.Product, error) {
	product := ps.Repository.FindById(id)

	if product == nil {
		return nil, errors.New("Product not found")
	} else {
		return product, nil
	}
}
