package service

import (
	"chal8/models"
	"chal8/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// var productRepo = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
// var productService = ProductService{Repository: productRepo}

func TestProductService_GetOneProductFound(t *testing.T) {
	productRepo := &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	productService := ProductService{Repository: productRepo}

	mock_product := models.Product{
		Title:       "barang",
		Description: "yaa ini barang",
	}

	productRepo.Mock.On("FindByID", "1").Return(mock_product)

	product, err := productService.GetOneProduct("1")

	assert.NotNil(t, product)
	assert.Nil(t, err)

	assert.Equal(t, &mock_product, product, "The returned product as same as the expected product")
}

func TestProductService_GetOneProductNotFound(t *testing.T) {
	productRepo := &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	productService := ProductService{Repository: productRepo}

	productRepo.Mock.On("FindByID", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)

	assert.Equal(t, "product not found", err.Error(), "Error response has to be 'product not found'")
}

func TestProductService_GetAllProductFound(t *testing.T) {
	productRepo := &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	productService := ProductService{Repository: productRepo}

	mock_products := []models.Product{
		{
			Title:       "barang",
			Description: "yaa ini barang",
		},
		{
			Title:       "produk",
			Description: "yaa ini produk",
		},
	}

	productRepo.Mock.On("FindAll").Return(mock_products)
	products, err := productService.GetAllProduct()

	assert.NotNil(t, products)
	assert.Nil(t, err)

	assert.Equal(t, &mock_products, products, "The returned all product as same as the expected product")
}

func TestProductService_GetAllProductNotFound(t *testing.T) {
	productRepo := &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	productService := ProductService{Repository: productRepo}

	productRepo.Mock.On("FindAll").Return(nil)

	products, err := productService.GetAllProduct()

	assert.Nil(t, products)
	assert.NotNil(t, err)

	assert.Equal(t, "all product not found", err.Error(), "Error response has to be 'all product not found'")
}
