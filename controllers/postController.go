package controllers

import (
	"go-crud-gin-gorm/models"
	"go-crud-gin-gorm/services"
	"go-crud-gin-gorm/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var postCreate models.PostCreate
	err := c.Bind(&postCreate)
	if err != nil {
		restErr := errors.NewBadRequestError("PC17 - Invalid request")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := services.CreatePost(postCreate)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetPosts(c *gin.Context) {
	result, restErr := services.GetPosts()
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetPost(c *gin.Context) {
	idString := c.Param("id")
	id, parsingError := strconv.Atoi(idString)
	if parsingError != nil {
		restErr := errors.NewBadRequestError("PC45 - Invalid request")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := services.GetPost(id)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdatePost(c *gin.Context) {
	idString := c.Param("id")
	var postUpdate models.PostCreate
	c.Bind(&postUpdate)

	id, parsingError := strconv.Atoi(idString)
	if parsingError != nil {
		restErr := errors.NewBadRequestError("PC66 - Invalid request")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := services.UpdatePost(id, postUpdate)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeletePost(c *gin.Context) {
	idString := c.Param("id")
	id, parsingError := strconv.Atoi(idString)
	if parsingError != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// if _, err := services.GetPost(id); err != nil {
	// 	if err.Error() == "record not found" {
	// 		c.Status(http.StatusNotFound)
	// 		return
	// 	}
	// }

	_, err := services.DeletePost(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
