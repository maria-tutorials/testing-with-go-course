package locations_provider

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCountryTimeoutFromApi(t *testing.T) {
	c, err := GetCountry("AR")

	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid restclient error when getting country AR", err.Message)

}

func TestGetCountryNotFound(t *testing.T) {
	c, err := GetCountry("AR")

	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	c, err := GetCountry("AR")

	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid error response when getting country AR", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	c, err := GetCountry("AR")

	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for AR", err.Message)
}

func TestGetCountryOK(t *testing.T) {
	c, err := GetCountry("AR")

	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.EqualValues(t, "Argentina", c.Name)
	assert.EqualValues(t, "GMT-03:00", c.TimeZone)
	assert.EqualValues(t, 24, len(c.States))
}
