package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/packages/connection"
	"go/packages/models"
	"net/http"
)

func GetUsers(context *gin.Context) {
	var users []models.User
	connection.DB.Find(&users)
	context.IndentedJSON(200, users)
}

func GetUser(context *gin.Context) {
	id := context.Param("id")
	var user models.User
	result := connection.DB.First(&user, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		context.IndentedJSON(400, gin.H{"msg": "User not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, user)
}

func DeleteUser(context *gin.Context) {
	id := context.Param("id")
	var user models.User
	result := connection.DB.First(&user, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "User not found"})
		return
	}
	connection.DB.Delete(&user, id)
	context.IndentedJSON(http.StatusOK, gin.H{"msg": "User deleted"})
}

func UpdateUser(context *gin.Context) {
	id := context.Param("id")
	var user models.User
	result := connection.DB.First(&user, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "User not found"})
		return
	}
	
	err := context.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Bad Request"})
		return
	}

	connection.DB.Save(user)
	context.IndentedJSON(http.StatusOK, gin.H{"msg": "User updated"})
}

func CreateUser(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Bad Request"})
		return
	}
	connection.DB.Create(&user)
	context.IndentedJSON(http.StatusCreated, gin.H{"msg": "User created"})
}
