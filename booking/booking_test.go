// +build integration

package booking

import (
	"testing"

	"github.com/matheusvill/aircnc/apartment"
	"github.com/matheusvill/aircnc/storage"
	"github.com/matheusvill/aircnc/user"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	userExpected := map[string]interface{}{
		"email":    "teste@teste.com",
		"password": "1234",
	}

	storage.InsertUser(
		userExpected["email"].(string),
		userExpected["password"].(string),
	)

	userId := storage.GetUserByEmail(userExpected["email"].(string))

	apartmentExpected := map[string]interface{}{
		"titulo":    "teste",
		"valor":     "1",
		"descricao": "testando",
		"cidade":    "floripa",
		"uf":        "SC",
	}

	storage.InsertApartment(
		apartmentExpected["titulo"].(string),
		apartmentExpected["valor"].(string),
		apartmentExpected["descricao"].(string),
		apartmentExpected["cidade"].(string),
		apartmentExpected["uf"].(string),
	)

	apartmentId := storage.GetApartmentByTitulo(apartmentExpected["titulo"].(string))

	b := Booking{
		User:      userId,
		Apartment: apartmentId,
	}

	Create(b)

	bookingExpected := map[string]interface{}{
		"user": user.User{
			Email:    "teste@teste.com",
			Password: "1234",
		},
		"apartment": apartment.Apartment{
			Titulo:    "teste",
			Valor:     "1",
			Descricao: "testando",
			Cidade:    "floripa",
			Uf:        "SC",
		},
	}

	bookingId := storage.GetBookingByUser(userId)
	bookingActual := Get(bookingId)

	assert.Exactly(t, bookingExpected, bookingActual, "Expected to be equal.")
}
