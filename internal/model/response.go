package model

import (
	"github.com/gin-gonic/gin"
)

type MainResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type DialogueResponse struct {
	Dialogue []DialogueView `json:"dialogue"`
}

type DialogueView struct {
	Speaker  string `json:"speaker"`
	Time     string `json:"time"`
	Sentence string `json:"sentence"`
}

func BuildHttpResponse(c *gin.Context, httpCode int, message string, data interface{}) {
	c.JSON(httpCode, MainResponse{
		Status:  httpCode,
		Message: message,
		Data:    data,
	})
}
