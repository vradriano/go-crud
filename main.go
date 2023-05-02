package main

import (
	"github.com/gin-gonic/gin"
	"go-crud/initializers"
	"go-crud/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

  routes.PostRoutes(r)

	r.Run()
}