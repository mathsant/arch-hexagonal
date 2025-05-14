package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
	GetQuantity() int
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64, quantity int) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReaderInterface interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriterInterface interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReaderInterface
	ProductWriterInterface
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID       string  `valid:"uuidv4"`
	Name     string  `valid:"required"`
	Price    float64 `valid:"float,optional"`
	Status   string  `valid:"required"`
	Quantity int     `valid:"required"`
}

func NewProduct() *Product {
	product := Product{
		ID:       uuid.NewV4().String(),
		Status:   DISABLED,
		Quantity: 0,
		Price:    0,
	}
	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	if p.Price <= 0 {
		return false, errors.New("the price must be greater of zero")
	}

	if p.Quantity <= 0 {
		return false, errors.New("the quantity must be greater of zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price <= 0 || p.Quantity <= 0 {
		return errors.New("the price and the quantity must be greater than zero to enable the product")
	}
	p.Status = ENABLED
	return nil
}

func (p *Product) Disable() error {
	if p.Quantity > 0 {
		return errors.New("the quantity must be zero in order to have the product disabled")
	}
	p.Status = DISABLED
	return nil
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetQuantity() int {
	return p.Quantity
}
