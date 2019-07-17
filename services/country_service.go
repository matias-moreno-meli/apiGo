package services

import (
	"../domains"
	"../utils"
)

func GetCountryFromAPI(countryID string) (*domains.Country, *utils.ApiError) {

	country := &domains.Country{
		ID: countryID,
	}
	if err := country.Get(); err != nil {
		return nil, err
	}

	return country, nil
}
