package controllers

import (
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
)

const paraSiteID = "siteId"

func GetSite(c *gin.Context) {
	siteID := c.Param(paraSiteID)

	//id, err := strconv.ParseInt(userID, 10, 64)
	//if err != nil {
	//	apiError := &utils.ApiError{
	//		Message: err.Error(),
	//		Status:  http.StatusBadRequest,
	//	}
	//
	//	c.JSON(apiError.Status, apiError)
	//	return
	//}

	site, apiError := services.GetSiteFromAPI(siteID)

	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, site)

}
