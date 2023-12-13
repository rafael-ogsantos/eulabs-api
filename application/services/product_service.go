package services

import (
	"context"

	"github.com/rafael-ogsantos/eulabs-api/application/repositories"
	"github.com/rafael-ogsantos/eulabs-api/domain"
)

// ProductServiceInterface interface
type ProductServiceInterface interface {
	FindById(ctx context.Context, id string) (*domain.Product, error)
	Insert(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Delete(ctx context.Context, id string) error
}

// ProductService struct
type ProductService struct {
	ProductRepository repositories.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(repository repositories.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: repository}
}

// FindById returns a product by id
func (service ProductService) FindById(ctx context.Context, id string) (*domain.Product, error) {
	product, err := service.ProductRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

// Insert inserts a new product
func (service ProductService) Insert(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	p, err := service.ProductRepository.Insert(ctx, product)

	if err != nil {
		return nil, err
	}

	return p, nil
}

// Update updates a product
func (service ProductService) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	updatedProduct, err := service.ProductRepository.Update(ctx, product)

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

// Delete deletes a product
func (service ProductService) Delete(ctx context.Context, id string) error {
	err := service.ProductRepository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
