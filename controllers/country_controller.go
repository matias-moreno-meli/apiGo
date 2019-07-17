package controllers

import (
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
)

const paraCountryID = "countryId"

func GetCountry(c *gin.Context) {
	countryID := c.Param(paraCountryID)

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

	country, apiError := services.GetCountryFromAPI(countryID)

	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, country)

}
