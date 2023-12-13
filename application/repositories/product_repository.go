package repositories

import (
	"context"
	"fmt"

	"github.com/rafael-ogsantos/eulabs-api/domain"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindById(ctx context.Context, id string) (*domain.Product, error)
	Insert(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Delete(ctx context.Context, id string) error
}

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func NewProductRepositoryDb(db *gorm.DB) *ProductRepositoryDb {
	return &ProductRepositoryDb{Db: db}
}

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

func (repo ProductRepositoryDb) FindById(ctx context.Context, id string) (*domain.Product, error) {
	var product domain.Product

	repo.Db.WithContext(ctx).First(&product, "id = ?", id)

	if product.ID == "" {
		return nil, fmt.Errorf("product does not exist")
	}

	return &product, nil
}

func (repo ProductRepositoryDb) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	err := repo.Db.WithContext(ctx).Save(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repo ProductRepositoryDb) Delete(ctx context.Context, id string) error {
	err := repo.Db.WithContext(ctx).Delete(&domain.Product{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
