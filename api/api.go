package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Question struct {
	ID   string                   `json:"id"`
	Text []map[string]interface{} `json:"text"`
}

type Answer struct {
	ID         string                   `json:"id"`
	QuestionID string                   `json:"questionId"`
	Text       []map[string]interface{} `json:"text"`
}

type ApiError struct {
	Error      error  `json:"detail"`
	StatusText string `json:"message"`
}

func NewError(status int, err error) *ApiError {
	return &ApiError{
		Error:      err,
		StatusText: http.StatusText(status),
	}
}

func WriteError(ctx *gin.Context, err error, status int) {
	ctx.JSON(status, NewError(status, err))
}
