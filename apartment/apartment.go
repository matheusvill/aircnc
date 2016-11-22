package apartment

import "github.com/matheusvill/aircnc/storage"

type Apartment struct {
	Titulo    string `json:"titulo"`
	Valor     string `json:"valor"`
	Descricao string `json:"descricao"`
	Cidade    string `json:"cidade"`
	Uf        string `json:"uf"`
}

func Create(apartment Apartment) error {
	storage.InsertApartment(
		apartment.Titulo,
		apartment.Valor,
		apartment.Descricao,
		apartment.Cidade,
		apartment.Uf,
	)

	return nil
}

func Get(id int) Apartment {
	apartmentMap := storage.GetApartment(id)

	ap := Apartment{}
	if apartmentMap != nil {
		ap.Titulo = apartmentMap["titulo"].(string)
		ap.Valor = apartmentMap["valor"].(string)
		ap.Descricao = apartmentMap["descricao"].(string)
		ap.Cidade = apartmentMap["cidade"].(string)
		ap.Uf = apartmentMap["uf"].(string)
	}

	return ap
}
