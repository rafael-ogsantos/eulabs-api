package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Product struct {
	ID          string    `valid:"uuid"`
	Name        string    `valid:"required"`
	Description string    `valid:"required"`
	CreatedAt   time.Time `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

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
