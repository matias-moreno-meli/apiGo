package services

import (
	"../domains"
	"../utils"
	"sync"
)

var user *domains.User
var country *domains.Country
var site *domains.Site
var wg sync.WaitGroup

func GetResultFromAPI(userID int64) (*domains.Result, *utils.ApiError) {

	wg.Add(1)

	go getUser(userID)
	wg.Wait()

	if user != nil {
		wg.Add(2)
		go getCountry(user.CountryID)
		go getSite(user.SiteID)
	}

	wg.Wait()

	result := &domains.Result{
		User:    user,
		Country: country,
		Site:    site,
	}

	return result, nil
}

func getUser(userID int64) *utils.ApiError {
	defer wg.Done()

	var err *utils.ApiError
	user, err = GetUserFromAPI(userID)
	if err != nil {
		return err
	}

	return nil

}

func getCountry(countryID string) *utils.ApiError {
	defer wg.Done()

	var err *utils.ApiError
	country, err = GetCountryFromAPI(countryID)
	if err != nil {
		return err
	}

	return nil

}

func getSite(siteID string) *utils.ApiError {
	defer wg.Done()
	var err *utils.ApiError
	site, err = GetSiteFromAPI(siteID)
	if err != nil {
		return err
	}

	return nil
}
