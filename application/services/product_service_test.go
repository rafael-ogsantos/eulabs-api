package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/rafael-ogsantos/eulabs-api/application/repositories"
	"github.com/rafael-ogsantos/eulabs-api/application/services"
	"github.com/rafael-ogsantos/eulabs-api/domain"
	"github.com/rafael-ogsantos/eulabs-api/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProductServiceInsert(t *testing.T) {
	db := database.NewDatabaseTest()
	ctx := context.Background()

	product := domain.NewProduct()
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Description = "Description 1"
	product.CreatedAt = time.Now()

	repo := repositories.ProductRepositoryDb{Db: db}
	service := services.NewProductService(&repo)
	result, err := service.Insert(ctx, product)

	require.Nil(t, err)
	require.Equal(t, result.Name, product.Name)
}

func TestProductServiceFindById(t *testing.T) {
	db := database.NewDatabaseTest()
	ctx := context.Background()

	product := domain.NewProduct()
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Description = "Description 1"
	product.CreatedAt = time.Now()

	repo := repositories.ProductRepositoryDb{Db: db}
	service := services.NewProductService(&repo)
	_, err := service.Insert(context.Background(), product)

	require.Nil(t, err)

	result, err := service.FindById(ctx, product.ID)

	require.Nil(t, err)
	require.Equal(t, result.ID, product.ID)
}

func TestProductServiceUpdate(t *testing.T) {
	db := database.NewDatabaseTest()

	product := domain.NewProduct()
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Description = "Description 1"
	product.CreatedAt = time.Now()

	repo := repositories.ProductRepositoryDb{Db: db}
	service := services.NewProductService(&repo)
	_, err := service.Insert(context.Background(), product)

	require.Nil(t, err)

	product.Name = "Product 2"
	result, err := service.Update(context.Background(), product.ID, product)

	require.Nil(t, err)
	require.Equal(t, result.Name, product.Name)
}

func TestProductServiceDelete(t *testing.T) {
	db := database.NewDatabaseTest()
	ctx := context.Background()

	product := domain.NewProduct()
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Description = "Description 1"
	product.CreatedAt = time.Now()

	repo := repositories.ProductRepositoryDb{Db: db}
	service := services.NewProductService(&repo)
	_, err := service.Insert(context.Background(), product)

	require.Nil(t, err)

	err = service.Delete(ctx, product.ID)

	require.Nil(t, err)
}
