package repositories

import (
	"go-crud-gin-gorm/initializers"
	"go-crud-gin-gorm/models"
	"go-crud-gin-gorm/utils/errors"
)

func CreatePost(postCreate models.PostCreate) (*models.PostResponse, *errors.RestErr) {
	post := models.Post{Title: postCreate.Title, Body: postCreate.Body}
	result := initializers.DB.Create(&post)
	if err := result.Error; err != nil {
		restError := errors.NewInternalServerError("PR13 - Database error")
		return nil, restError
	}
	postResponse := &models.PostResponse{ID: post.ID, Title: post.Title, Body: post.Body}
	return postResponse, nil
}

func GetPosts() ([]models.PostResponse, *errors.RestErr) {
	var posts []models.Post
	allPosts := initializers.DB.Find(&posts)
	if err := allPosts.Error; err != nil {
		restError := errors.NewInternalServerError("PR24 - Database error")
		return nil, restError
	}

	var postResponses []models.PostResponse
	for _, post := range posts {
		postResponse := models.MapPostToPostResponse(post)
		postResponses = append(postResponses, postResponse)
	}

	return postResponses, nil
}

func GetPost(id int) (*models.PostResponse, *errors.RestErr) {
	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		if err.Error() == "record not found" {
			restError := errors.NewBadRequestError("PR41 - Post not found")
			return nil, restError
		} else {
			restError := errors.NewInternalServerError("PR44 - Database error")
			return nil, restError
		}
	}
	postResponse := &models.PostResponse{ID: post.ID, Title: post.Title, Body: post.Body}
	return postResponse, nil
}

func UpdatePost(id int, postUpdate models.PostCreate) (*models.PostResponse, *errors.RestErr) {
	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		if err.Error() == "record not found" {
			restError := errors.NewBadRequestError("PR56 - Post not found")
			return nil, restError
		} else {
			restError := errors.NewInternalServerError("PR59 - Database error")
			return nil, restError
		}
	}

	initializers.DB.Model(&post).Updates(models.Post{
		Title: postUpdate.Title,
		Body:  postUpdate.Body,
	})
	postResponse := &models.PostResponse{ID: post.ID, Title: post.Title, Body: post.Body}

	return postResponse, nil
}

func DeletePost(id int) (*models.PostResponse, *errors.RestErr) {
	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		if err.Error() == "record not found" {
			restError := errors.NewBadRequestError("PR77 - Post not found")
			return nil, restError
		} else {
			restError := errors.NewInternalServerError("PR80 - Database error")
			return nil, restError
		}
	}

	initializers.DB.Delete(&models.Post{}, id)

	postResponse := &models.PostResponse{ID: post.ID, Title: post.Title, Body: post.Body}
	return postResponse, nil
}
