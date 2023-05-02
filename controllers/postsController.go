package controllers

import (
	"net/http"

  "github.com/gin-gonic/gin"
	"go-crud/models"
	"go-crud/repositories"
)

func GetPosts(c *gin.Context) {
	posts, err := repositories.FindAllPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve posts!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posts found with success!",
		"data":    posts,
	})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")

	post, err := repositories.FindPostById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post found with success!",
		"data": post,
	})
}

func PostsCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	if body.Title == "" || body.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title or body can't be empty!",
		})

		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}
	
	if err := repositories.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "The post could not be created. Please try again!",
		})

		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created with success!",
		"data":  post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct{
		Title string
		Body string
	}

	c.Bind(&body)

	post, err := repositories.FindPostById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found!",
		})
		return
	}

	err = repositories.UpdatePost(&post, body.Title, body.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "The post could not be updated. Please try again!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated with success!",
		"data": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	_, err := repositories.FindPostById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found!",
		})
		return
	}

	if err := repositories.DeletePostById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "The post could not be deleted. Please try again.",
		})
		
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted with success!",
	})
}