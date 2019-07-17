package services

import (
	"../domains"
	"../utils"
)

func GetSiteFromAPI(siteID string) (*domains.Site, *utils.ApiError) {

	site := &domains.Site{
		ID: siteID,
	}
	if err := site.Get(); err != nil {
		return nil, err
	}

	return site, nil
}
