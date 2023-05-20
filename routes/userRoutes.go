package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func UserAllRoutes(router *gin.Engine)  {
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.PATCH("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}