package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDetectHistory() gin.HandlerFunc {
	// TODO: Implement it.
	return func(ctx *gin.Context) {
		ctx.String(http.StatusServiceUnavailable, "Now featuring")
	}
}
