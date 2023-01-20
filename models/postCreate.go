package models

type PostCreate struct {
	Title string `binding:"required"`
	Body  string `binding:"required"`
}
