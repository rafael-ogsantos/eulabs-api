package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/rafael-ogsantos/eulabs-api/application/repositories"
	"github.com/rafael-ogsantos/eulabs-api/domain"
	"github.com/rafael-ogsantos/eulabs-api/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProductDbInsert(t *testing.T) {
	db := database.NewDatabaseTest()
	ctx := context.Background()

	product := domain.NewProduct()
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Description = "Description 1"
	product.CreatedAt = time.Now()

	repo := repositories.ProductRepositoryDb{Db: db}
	_, err := repo.Insert(ctx, product)
	if err != nil {
		t.Error("Error to insert product")
	}

	p, err := repo.FindById(ctx, product.ID)
	if err != nil {
		t.Error("Error to find product")
	}

	require.NotEmpty(t, p.ID)
	require.Nil(t, err)
	require.Equal(t, p.ID, product.ID)
}
