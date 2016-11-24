// +build integration

package user

import (
	"testing"

	"github.com/matheusvill/aircnc/storage"
	"github.com/stretchr/testify/assert"
)

func TestCheckEmail(t *testing.T) {
	tests := map[string]bool{
		"valido@gmai.com": true,
		"invalido":        false,
	}

	for value, resultExpected := range tests {
		resultActual := CheckEmail(value)
		assert.Exactly(t, resultExpected, resultActual, "Expected to be equal.")
	}
}

func TestCreate(t *testing.T) {
	userExpected := User{
		Email:    "teste@teste.com",
		Password: "1234",
	}

	Create(userExpected)

	userId := storage.GetUserByEmail(userExpected.Email)
	userActual := Get(userId)

	assert.Exactly(t, userExpected, userActual, "Expected to be equal.")
}
