package application_test

import (
	"testing"

	"github.com/mathsant/go-arch-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Produto"
	product.Status = application.DISABLED
	product.Price = 10
	product.Quantity = 2

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.NotNil(t, err)
	require.Equal(t, "the price and the quantity must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Produto"
	product.Status = application.DISABLED
	product.Price = 0
	product.Quantity = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Quantity = 10
	err = product.Disable()
	require.NotNil(t, err)
	require.Equal(t, "the quantity must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Produto"
	product.Status = application.DISABLED
	product.Price = 10
	product.Quantity = 2

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.NotNil(t, err)
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -100
	_, err = product.IsValid()
	require.NotNil(t, "the price must be greater of zero", err.Error())

	product.Price = 100
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Quantity = 0
	_, err = product.IsValid()
	require.NotNil(t, "the quantity must be greater of zero", err.Error())
}
