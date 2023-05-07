package service

import (
	"chal8/models"
	"chal8/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (s *ProductService) GetOneProduct(id string) (*models.Product, error) {
	product := s.Repository.FindByID(id)

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (s *ProductService) GetAllProduct() (*[]models.Product, error) {
	product := s.Repository.FindAll()

	if product == nil {
		return nil, errors.New("all product not found")
	}

	return product, nil
}
