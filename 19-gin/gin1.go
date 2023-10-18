package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/home", Home)
	router.Run(":8080")
}

type User struct {
	FName    string `json:"f_name"` // giving the name of the field in the json ouptut
	Password string `json:"-"`      //ignore the field while creating the json
	LName    string `json:"l_name"`
	Email    string `json:"email"`
}

// func(*Context)
func Home(c *gin.Context) {
	//c.String(http.StatusOK, "this is my home page")
	//using the map to send the json response
	c.JSON(http.StatusOK, gin.H{"msg": "this is my home page"})
	u1 := User{FName: "k", Password: "j", LName: "j", Email: "gg"}
	c.JSON(http.StatusOK, u1)
	
}
