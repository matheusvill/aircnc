// +build unit

package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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
