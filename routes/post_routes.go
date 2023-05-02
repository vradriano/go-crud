package routes

import (
	"github.com/gin-gonic/gin"
	"go-crud/controllers"
)

func PostRoutes(r *gin.Engine) {
	postRoutes := r.Group("/posts")
	{
		postRoutes.GET("/", controllers.GetPosts)
		postRoutes.GET("/:id", controllers.GetPostById)
		postRoutes.PUT("/:id", controllers.UpdatePost)
		postRoutes.POST("", controllers.PostsCreate)
		postRoutes.DELETE("/:id", controllers.DeletePost)
	}
}
