package repository

import (
	"chal8/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (r *ProductRepositoryMock) FindByID(id string) *models.Product {
	arguments := r.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(models.Product)
	return &product
}

func (r *ProductRepositoryMock) FindAll() *[]models.Product {
	arguments := r.Mock.Called()

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).([]models.Product)

	return &product
}
