// +build unit

package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gitlab.neoway.com.br/jaws/stormtroopers/catalog"
)

type AddressTestSuite struct {
	suite.Suite
	address *Address
	c       *catalog.MockCatalog
}

func (suite *AddressTestSuite) SetupTest() {
	suite.c = &catalog.MockCatalog{}
	suite.address = NewAddress(suite.c)
}
func (suite *AddressTestSuite) TestNormalizeAddressWithAllFields() {
	typeExpected := "RUA"
	nameExpected := ""
	streetExpected := "PATRICIO FARIAS"
	numberExpected := "131"
	complementExpected := "2 ANDAR"
	neighborhoodExpected := "ITACORUBI"
	cityExpected := "FLORIANOPOLIS"
	stateExpected := "SC"
	cepExpected := "88034132"
	addressHashExpected := "b1de7492f8db223860c08b4d10611b6d"
	geocodeHashExpected := "9a06baf8f94e471c1623038c5139e101"
	address := []string{"Rua Patrício Farias", "131", "2 andar", "", "Florianópolis", "SC", "88034-132"}
	suite.c.On("NormalizeAddress", address[0], address[1], address[2], address[3], address[4], address[5], address[6]).
		Return([]string{
			typeExpected,
			nameExpected,
			streetExpected,
			numberExpected,
			complementExpected,
			neighborhoodExpected,
			cityExpected,
			stateExpected,
			cepExpected,
			addressHashExpected,
			geocodeHashExpected,
		}, nil)
	addressNormalized, err := suite.address.NormalizeAddress(
		address[0], // street
		address[1], // number
		address[2], // complement
		address[3], // neighborhood
		address[4], // city
		address[5], // state
		address[6], // cep
	)
	typeAddress := addressNormalized[0]
	name := addressNormalized[1]
	street := addressNormalized[2]
	number := addressNormalized[3]
	complement := addressNormalized[4]
	neighborhood := addressNormalized[5]
	city := addressNormalized[6]
	state := addressNormalized[7]
	cep := addressNormalized[8]
	addressHash := addressNormalized[9]
	geocodeHash := addressNormalized[10]
	assert.Exactly(suite.T(), typeExpected, typeAddress, "Expected to be equal.")
	assert.Exactly(suite.T(), nameExpected, name, "Expected to be equal.")
	assert.Exactly(suite.T(), streetExpected, street, "Expected to be equal.")
	assert.Exactly(suite.T(), numberExpected, number, "Expected to be equal.")
	assert.Exactly(suite.T(), complementExpected, complement, "Expected to be equal.")
	assert.Exactly(suite.T(), neighborhoodExpected, neighborhood, "Expected to be equal.")
	assert.Exactly(suite.T(), cityExpected, city, "Expected to be equal.")
	assert.Exactly(suite.T(), stateExpected, state, "Expected to be equal.")
	assert.Exactly(suite.T(), cepExpected, cep, "Expected to be equal.")
	assert.Exactly(suite.T(), addressHashExpected, addressHash, "Expected to be equal.")
	assert.Exactly(suite.T(), geocodeHashExpected, geocodeHash, "Expected to be equal.")
	assert.Nil(suite.T(), err, "Expected to be nil.")
	suite.c.AssertExpectations(suite.T())
}
func (suite *AddressTestSuite) TestGetGeoLocationWithAllFields() {
	latitudeExpected := "27.111111"
	longitudeExpected := "30.222222"
	locationTypeExpected := "ROOF TOP"
	geoHash := "9a06baf8f94e471c1623038c5139e101"
	suite.c.On("GetGeoLocation", geoHash).Return(latitudeExpected, longitudeExpected, locationTypeExpected, nil)
	latitude, longitude, locationType, err := suite.address.GetGeoLocation(geoHash)
	assert.Exactly(suite.T(), latitudeExpected, latitude, "Expected to be equal.")
	assert.Exactly(suite.T(), longitudeExpected, longitude, "Expected to be equal.")
	assert.Exactly(suite.T(), locationTypeExpected, locationType, "Expected to be equal.")
	assert.Nil(suite.T(), err, "Expected to be nil.")
	suite.c.AssertExpectations(suite.T())
}
func TestAddressTestSuite(t *testing.T) {
	suite.Run(t, new(AddressTestSuite))
}
