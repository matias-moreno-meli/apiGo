package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main() {

	router.GET("/user/:userId", controllers.GetUser)
	router.GET("/country/:countryId", controllers.GetCountry)
	router.GET("/site/:siteId", controllers.GetSite)
	router.GET("/result/:userId", controllers.GetResult)
	router.Run(port)

}
