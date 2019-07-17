package domains

import (
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Country struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Locale             string `json:"locale"`
	CurrencyId         string `json:"currency_id"`
	DecimalSeparator   string `json:"decimal_separator"`
	thousandsSeparator string `json:"thousands_separator"`
	TimeZone           string `json:"time_zone"`
	GeoInformation     struct {
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"states"`
}

func (country *Country) Get() *utils.ApiError {
	if len(country.ID) == 0 {
		return &utils.ApiError{
			Message: "El id de usuario es 0",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlCountry, country.ID)

	response, err := http.Get(url)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &country); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil

}
