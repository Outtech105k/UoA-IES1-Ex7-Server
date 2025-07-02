package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostFaceAuthHandler() gin.HandlerFunc {
	// TODO: Implement it.
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}
