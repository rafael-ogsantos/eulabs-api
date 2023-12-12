package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Product struct {
	ID          string    `valid:"uuid"`
	Name        string    `valid:"required"`
	Description string    `valid:"required"`
	CreatedAt   time.Time `valid:"-"` // o "-" significa que não quero validar esse campo
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// o * e o & são ponteiros de memória que apontam para o endereço de memória de uma variável
// quando usado, eu altero o valor da variável original e não uma cópia dela, por exemplo se eu passar
// um ponteiro de memória para uma função, e alterar o valor da variável dentro da função, quando eu sair da função, o valor da variável original
// também terá sido alterado
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
