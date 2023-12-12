package domain_test

import (
	"testing"
	"time"

	"github.com/rafael-ogsantos/eulabs-api/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfProductIsEmpty(t *testing.T) {
	product := domain.NewProduct()
	err := product.Validate()

	require.Error(t, err)
}

func TestProductIsNotAUuid(t *testing.T) {
	product := domain.NewProduct()
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Description = "Description 1"
	product.CreatedAt = time.Now()

	err := product.Validate()
	require.Nil(t, err)
}
