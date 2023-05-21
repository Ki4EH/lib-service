package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Msg string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, err string) {
	logrus.Error(err)
	c.AbortWithStatusJSON(statusCode, errorResponse{Msg: err})
}
