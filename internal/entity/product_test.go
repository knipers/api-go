package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Mouse", 39.5)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Mouse", p.Name)
	assert.Equal(t, 39.5, p.Price)
}

func TestNewProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 39.5)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestNewProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Mouse", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestNewProductWhenInvalidPrice(t *testing.T) {
	p, err := NewProduct("Mouse", -39.5)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidade(t *testing.T) {
	p, err := NewProduct("Mouse", 39.5)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotNil(t, p.Validate())
}
