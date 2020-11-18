package locations_provider

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../../domain/locations"
	"../../utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

const urlGetCountry = "https://api.mercadolibre.com/countries/%s"

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	url := fmt.Sprintf(urlGetCountry, countryId)

	resp := rest.Get(url)
	if resp == nil || resp.Response == nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryId),
		}
	}

	if resp.StatusCode > 299 {
		apiErr := errors.ApiError{}
		if err := json.Unmarshal(resp.Bytes(), &apiErr); err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error interface when getting country %s", countryId),
			}
		}
		return nil, &apiErr
	}

	result := locations.Country{}
	if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
		}
	}

	return &result, nil
}
