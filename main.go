package main

import (
	"go-crud/initializers"
	"go-crud/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	router := gin.Default()

	routes.UserAllRoutes(router)

	router.Run(":" + os.Getenv("LOCAL_PORT"))
}