package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mathsant/go-arch-hexagonal/adapters/cli"
	mock_application "github.com/mathsant/go-arch-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"
	productQuantity := 10

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetQuantity().Return(productQuantity).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice, productQuantity).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name: %s has been created with price: %f, quantity: %d and status: %s",
		productId, productName, productPrice, productQuantity, productStatus)
	result, err := cli.Run(service, "create", "", productName, productPrice, productQuantity)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled.", productName)
	result, err = cli.Run(service, "enable", productId, "", 0, 1)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled.", productName)
	result, err = cli.Run(service, "disable", productId, "", 0, 1)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\n Name: %s\n Price: %f\n Quantity: %d\n Status: %s",
		productId, productName, productPrice, productQuantity, productStatus)
	result, err = cli.Run(service, "get", productId, "", 0, 1)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
