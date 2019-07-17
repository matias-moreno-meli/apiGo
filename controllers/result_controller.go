package controllers

import (
	"../services"
	"../utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const id = "userId"

func GetResult(c *gin.Context) {
	userID := c.Param(id)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiError := &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}

		c.JSON(apiError.Status, apiError)
		return
	}

	result, apiError := services.GetResultFromAPI(id)

	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, result)

}
