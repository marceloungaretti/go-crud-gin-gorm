package models

type PostResponse struct {
	ID    uint
	Title string
	Body  string `json:"body,omitempty"`
}

func MapPostToPostResponse(post Post) PostResponse {
	return PostResponse{ID: post.ID, Title: post.Title, Body: post.Body}
}
