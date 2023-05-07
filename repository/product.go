package repository

import "chal8/models"

type ProductRepository interface {
	FindByID(id string) *models.Product
	FindAll() *[]models.Product
}
