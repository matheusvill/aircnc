// +build integration
package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	CreateDb()

	resultExpected := map[string]interface{}{
		"email":    "teste@teste.com",
		"password": "1234",
	}

	InsertUser(
		resultExpected["email"].(string),
		resultExpected["password"].(string),
	)

	userId := GetUserByEmail(resultExpected["email"].(string))
	resultActual := GetUser(int(userId))

	assert.Exactly(t, resultExpected, resultActual, "Expected to be equal.")

	// DropDb()
}

// func TestApartment(t *testing.T) {
// 	CreateDb()

// 	resultExpected := map[string]interface{}{
// 		"titulo":    "teste",
// 		"valor":     "1",
// 		"descricao": "testando",
// 		"cidade":    "floripa",
// 		"uf":        "SC",
// 	}

// 	InsertApartment(
// 		resultExpected["titulo"].(string),
// 		resultExpected["valor"].(string),
// 		resultExpected["descricao"].(string),
// 		resultExpected["cidade"].(string),
// 		resultExpected["uf"].(string),
// 	)

// 	apartmentId := GetApartmentByTitulo(resultExpected["titulo"].(string))
// 	resultActual := GetApartment(int(apartmentId))

// 	assert.Exactly(t, resultExpected, resultActual, "Expected to be equal.")

// 	DropDb()
// }
