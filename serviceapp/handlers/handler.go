package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func API() *gin.Engine {

	router := gin.New()
	router.GET("/check", Check)
	// router.Run(":8080")
	return router

}
func Check(c *gin.Context) {
	c.JSON(http.StatusOK, "OKAY11")

}
