package repositories

import (
	"context"
	"fmt"

	"github.com/rafael-ogsantos/eulabs-api/domain"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ProductRepository interface
type ProductRepository interface {
	FindById(ctx context.Context, id string) (*domain.Product, error)
	Insert(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Delete(ctx context.Context, id string) error
}

// ProductRepositoryDb struct
type ProductRepositoryDb struct {
	Db *gorm.DB
}

// NewProductRepositoryDb creates a new product repository
func NewProductRepositoryDb(db *gorm.DB) *ProductRepositoryDb {
	return &ProductRepositoryDb{Db: db}
}

// FindById returns a product by id
func (repo ProductRepositoryDb) Insert(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	if product.ID == "" {
		product.ID = uuid.NewV4().String()
	}

	err := repo.Db.WithContext(ctx).Create(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

// FindById returns a product by id
func (repo ProductRepositoryDb) FindById(ctx context.Context, id string) (*domain.Product, error) {
	var product domain.Product

	repo.Db.WithContext(ctx).First(&product, "id = ?", id)

	if product.ID == "" {
		return nil, fmt.Errorf("product does not exist")
	}

	return &product, nil
}

// Update updates a product
func (repo ProductRepositoryDb) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	err := repo.Db.WithContext(ctx).Save(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

// Delete deletes a product
func (repo ProductRepositoryDb) Delete(ctx context.Context, id string) error {
	err := repo.Db.WithContext(ctx).Delete(&domain.Product{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
