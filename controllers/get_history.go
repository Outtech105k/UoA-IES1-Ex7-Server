package controllers

import (
	"net/http"

	"github.com/Outtech105k/UoA-IES1-Ex7-Server/repos"
	"github.com/Outtech105k/UoA-IES1-Ex7-Server/utils"
	"github.com/gin-gonic/gin"
)

func GetDetectHistory(appCtx utils.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var histories []repos.Histories
		err := appCtx.DB.Select(&histories, "SELECT * FROM histories ORDER BY detected DESC")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, histories)
	}
}
