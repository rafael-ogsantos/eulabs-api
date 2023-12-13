package services

import (
	"context"

	"github.com/rafael-ogsantos/eulabs-api/application/repositories"
	"github.com/rafael-ogsantos/eulabs-api/domain"
)

type ProductServiceInterface interface {
	FindById(ctx context.Context, id string) (*domain.Product, error)
	Insert(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Delete(ctx context.Context, id string) error
}

type ProductService struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(repository repositories.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: repository}
}

func (service ProductService) FindById(ctx context.Context, id string) (*domain.Product, error) {
	product, err := service.ProductRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service ProductService) Insert(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	p, err := service.ProductRepository.Insert(ctx, product)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (service ProductService) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	updatedProduct, err := service.ProductRepository.Update(ctx, product)

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (service ProductService) Delete(ctx context.Context, id string) error {
	err := service.ProductRepository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
