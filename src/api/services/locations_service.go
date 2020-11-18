package services

import (
	"../domain/locations"
	"../providers/locations_provider"
	"../utils/errors"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryId)
}
