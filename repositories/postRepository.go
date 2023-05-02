package repositories

import (
	"go-crud/models"
	"go-crud/initializers"
)

func FindAllPosts() ([]models.Post, error) {
	var posts []models.Post

	err := initializers.DB.Find(&posts).Error

	return posts, err
}

func FindPostById(id string) (models.Post, error) {
	var post models.Post

	err := initializers.DB.First(&post, id).Error

	return post, err
}

func CreatePost(post *models.Post) error {
	return initializers.DB.Create(&post).Error
}

func UpdatePost(post *models.Post, title string, body string) error {
	return initializers.DB.Model(&post).Updates(models.Post{
		Title: title,
		Body:  body,
	}).Error
}

func DeletePostById(id string) error {
	return initializers.DB.Delete(&models.Post{}, id).Error
}