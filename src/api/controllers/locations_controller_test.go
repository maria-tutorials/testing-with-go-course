package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-restclient/rest"

	"../domain/locations"
	"../services"
	"../utils/errors"
)

var (
	getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

type locationsServiceMock struct{}

func (*locationsServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryId)
}

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryNotFound(t *testing.T) {

	// Mock LocationsService methods:
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{Status: http.StatusNotFound, Message: "Country not found"}
	}
	services.LocationsService = &locationsServiceMock{}

	// gin
	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, resp.Code)

	var apiErr errors.ApiError
	err := json.Unmarshal(resp.Body.Bytes(), &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}

func TestGetCountryOK(t *testing.T) {
	// Mock LocationsService methods:
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return &locations.Country{Id: "AR", Name: "Argentina"}, nil
	}
	services.LocationsService = &locationsServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusOK, response.Code)

	var country locations.Country
	err := json.Unmarshal(response.Body.Bytes(), &country)
	assert.Nil(t, err)

	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}
