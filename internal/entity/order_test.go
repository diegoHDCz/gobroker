package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfGetsAnErroIfIDISBlanck(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func Test_If_It_Get_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func Test_If_It_Get_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ID: "123", Price: 123.0}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "123", Price: 123.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	order.CalculateFinalPrice()
	assert.Equal(t, 124.0, order.FinalPrice)

}
