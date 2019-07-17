package controllers

import (
	"../services"
	"../utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const paraUserID = "userId"

func GetUser(c *gin.Context) {
	userID := c.Param(paraUserID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiError := &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}

		c.JSON(apiError.Status, apiError)
		return
	}

	user, apiError := services.GetUserFromAPI(id)

	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, user)

}
