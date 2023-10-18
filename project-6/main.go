package main

import (
	"fmt"
	"log"
	"net/http"
	"project-6/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/home/:user_id", Home)
	router.Run(":8081")
}

// func(*Context)
func Home(c *gin.Context) {

	userIdString := c.Param("user_id")
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {

		log.Println("Error: ", err)

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "not a valid number"})

		return

	}
	fmt.Println(userId)
	uData, err := models.FetchUser(userId)
	if err != nil {
		log.Println("Error: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "not a valid number"})
		return
	}

	c.JSON(http.StatusOK, uData)
	//c.String(http.StatusOK, "this is my home page")
	//using the map to send the json response
	// c.JSON(http.StatusOK, gin.H{"msg": "this is my home page"})
	// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
}
