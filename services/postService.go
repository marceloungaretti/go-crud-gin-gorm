package services

import (
	"go-crud-gin-gorm/models"
	"go-crud-gin-gorm/repositories"
	"go-crud-gin-gorm/utils/errors"
)

func CreatePost(postCreate models.PostCreate) (*models.PostResponse, *errors.RestErr) {
	result, err := repositories.CreatePost(postCreate)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetPosts() ([]models.PostResponse, *errors.RestErr) {
	allPosts, err := repositories.GetPosts()
	if err != nil {
		return nil, err
	}

	return allPosts, nil
}

func GetPost(id int) (*models.PostResponse, *errors.RestErr) {
	post, err := repositories.GetPost(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func UpdatePost(id int, postUpdate models.PostCreate) (*models.PostResponse, *errors.RestErr) {
	post, err := repositories.UpdatePost(id, postUpdate)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func DeletePost(id int) (*models.PostResponse, *errors.RestErr) {
	post, err := repositories.DeletePost(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}
