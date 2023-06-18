package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var listUsers = []User{
	User{Id: 1, Name: "Alex", Age: 25},
	User{Id: 2, Name: "Paulo", Age: 30},
}

func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Ok",
	})
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, listUsers)
}

func postUser(c *gin.Context) {
	var user User

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ErrorMessage": "Ocorreu um erro",
		})
		return
	}

	user.Id = (int16)(len(listUsers) + 1)
	listUsers = append(listUsers, user)

	c.JSON(http.StatusOK, gin.H{
		"result": "Insert user with successfuly",
	})
}

func main() {
	service := gin.Default()

	getRoutes(service)

	service.Run()

}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/health", getHealth)

	userController := c.Group("/user")
	userController.GET("/", getUsers)
	userController.POST("/", postUser)

	return c
}
