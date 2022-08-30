package repository

import (
	"GolangUnitTest/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arg := repository.Mock.Called(id)

	if arg.Get(0) == nil {
		return nil
	} else {
		category := arg.Get(0).(entity.Category)
		return &category
	}
}
