// +build integration

package apartment

import (
	"testing"

	"github.com/matheusvill/aircnc/storage"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	apartmentExpected := Apartment{
		Titulo:    "teste",
		Valor:     "1",
		Descricao: "testando",
		Cidade:    "floripa",
		Uf:        "SC",
	}

	storage.InsertApartment(
		apartmentExpected.Titulo,
		apartmentExpected.Valor,
		apartmentExpected.Descricao,
		apartmentExpected.Cidade,
		apartmentExpected.Uf,
	)

	Create(apartmentExpected)

	apartmentId := storage.GetApartmentByTitulo(apartmentExpected.Titulo)
	apartmentActual := Get(apartmentId)

	assert.Exactly(t, apartmentExpected, apartmentActual, "Expected to be equal.")
}
