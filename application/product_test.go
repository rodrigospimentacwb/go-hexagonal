package application_test

import (
	"github.com/rodrigospimentacwb/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	require "github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enabled()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enabled()
	require.NotNil(t, err)
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Error(t, err, "the price must be zero in order to have the product disabled")
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater than zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{ID: uuid.NewV4().String(), Name: "Hello", Price: 10, Status: application.ENABLED}
	require.Equal(t, product.GetID(), product.ID)
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{ID: uuid.NewV4().String(), Name: "Hello", Price: 10, Status: application.ENABLED}
	require.Equal(t, product.GetName(), "Hello")
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{ID: uuid.NewV4().String(), Name: "Hello", Price: 10, Status: application.ENABLED}
	require.Equal(t, product.GetPrice(), float64(10))
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{ID: uuid.NewV4().String(), Name: "Hello", Price: 10, Status: application.ENABLED}
	require.Equal(t, product.GetStatus(), application.ENABLED)

	product.Status = application.DISABLED
	require.Equal(t, product.GetStatus(), application.DISABLED)
}
