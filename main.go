package main

import (
  "github.com/gin-gonic/gin"
  "go-crud/initializers"
  "go-crud/controllers"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

  r.GET("/posts", controllers.GetPosts)
  r.GET("/posts/:id", controllers.GetPostById)

  r.PUT("/posts/:id", controllers.UpdatePost)
  r.POST("/posts", controllers.PostsCreate)

  r.DELETE("/posts/:id", controllers.DeletePost)

  r.Run()
}