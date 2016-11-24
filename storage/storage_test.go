// +build integration
package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

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

}

func TestApartment(t *testing.T) {

	resultExpected := map[string]interface{}{
		"titulo":    "teste",
		"valor":     "1",
		"descricao": "testando",
		"cidade":    "floripa",
		"uf":        "SC",
	}

	InsertApartment(
		resultExpected["titulo"].(string),
		resultExpected["valor"].(string),
		resultExpected["descricao"].(string),
		resultExpected["cidade"].(string),
		resultExpected["uf"].(string),
	)

	apartmentId := GetApartmentByTitulo(resultExpected["titulo"].(string))
	resultActual := GetApartment(int(apartmentId))

	assert.Exactly(t, resultExpected, resultActual, "Expected to be equal.")

}

func TestBooking(t *testing.T) {

	userExpected := map[string]interface{}{
		"email":    "teste@teste.com",
		"password": "1234",
	}

	InsertUser(
		userExpected["email"].(string),
		userExpected["password"].(string),
	)

	apartmentExpected := map[string]interface{}{
		"titulo":    "teste",
		"valor":     "1",
		"descricao": "testando",
		"cidade":    "floripa",
		"uf":        "SC",
	}

	InsertApartment(
		apartmentExpected["titulo"].(string),
		apartmentExpected["valor"].(string),
		apartmentExpected["descricao"].(string),
		apartmentExpected["cidade"].(string),
		apartmentExpected["uf"].(string),
	)

	userId := GetUserByEmail(userExpected["email"].(string))
	apartmentId := GetApartmentByTitulo(apartmentExpected["titulo"].(string))

	bookingExpected := map[string]interface{}{
		"user":      userId,
		"apartment": apartmentId,
	}

	InsertBooking(userId, apartmentId)

	bookingId := GetBookingByUser(userId)
	bookingActual := GetBooking(bookingId)

	assert.Exactly(t, bookingExpected, bookingActual, "Expected to be equal.")

}
