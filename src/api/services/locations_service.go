package services

import (
	"../domain/locations"
	"../providers/locations_provider"
	"../utils/errors"
)

type locationsService struct{}

type locationsServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var LocationsService locationsServiceInterface

func init() {
	LocationsService = &locationsService{}
}

func (s *locationsService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryId)
}
