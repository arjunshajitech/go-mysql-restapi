package routes

import (
	"github.com/gin-gonic/gin"
	"go/packages/handler"
	"go/packages/middleware"
)

func InitlizeRoute() {
	router := gin.Default()
	router.Use(middleware.Logger)

	router.GET("/user", handler.GetUsers)
	router.GET("/user/:id", handler.GetUser)
	router.DELETE("/user/:id", handler.DeleteUser)
	router.PATCH("/user/:id", handler.UpdateUser)
	router.POST("/user", handler.CreateUser)

	router.Run("127.0.0.1:8888")
}
