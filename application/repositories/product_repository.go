package repositories

import (
	"fmt"

	"github.com/rafael-ogsantos/eulabs-api/domain"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindById(id string) (*domain.Product, error)
	Insert(product *domain.Product) (*domain.Product, error)
	Update(product *domain.Product) (*domain.Product, error)
	Delete(id string) error
}

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func NewProductRepositoryDb(db *gorm.DB) *ProductRepositoryDb {
	return &ProductRepositoryDb{Db: db}
}

func (repo ProductRepositoryDb) Insert(product *domain.Product) (*domain.Product, error) {
	if product.ID == "" {
		product.ID = uuid.NewV4().String()
	}

	err := repo.Db.Create(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repo ProductRepositoryDb) FindById(id string) (*domain.Product, error) {
	var product domain.Product

	repo.Db.First(&product, "id = ?", id)

	if product.ID == "" {
		return nil, fmt.Errorf("product does not exist")
	}

	return &product, nil
}
