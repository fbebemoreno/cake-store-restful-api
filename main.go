package main

import (
	"github.com/fbebemoreno/cake-store-restful-api/controllers/cakecontroller"
	"github.com/fbebemoreno/cake-store-restful-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDB()

	router.GET("/cakes", cakecontroller.Index)
	router.GET("/cakes/:id", cakecontroller.Show)
	router.POST("/cakes", cakecontroller.Create)
	router.PUT("/cakes/:id", cakecontroller.Update)
	router.DELETE("/cakes/:id", cakecontroller.Delete)

	router.Run(":5000")
}
