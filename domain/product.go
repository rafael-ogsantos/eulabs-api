package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Product is a struct that represents a product
type Product struct {
	ID          string    `valid:"uuid"`
	Name        string    `valid:"required"`
	Description string    `valid:"required"`
	CreatedAt   time.Time `valid:"-"`
}

// init is a function that is called before the execution of the program
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// NewProduct is a function that returns a new product
func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Validate() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}
