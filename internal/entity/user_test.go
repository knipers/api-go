package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Senhor developer", "senhorD@go.dev", "dev1234")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Senhor developer", user.Name)
	assert.Equal(t, "senhorD@go.dev", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Senhor developer", "senhorD@go.dev", "dev1234")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("dev1234"))
	assert.False(t, user.ValidatePassword("dev12345"))
	assert.NotEqual(t, "dev12345", user.Password)
}
