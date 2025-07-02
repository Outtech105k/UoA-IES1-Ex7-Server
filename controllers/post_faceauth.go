package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Outtech105k/UoA-IES1-Ex7-Server/models"
	"github.com/Outtech105k/UoA-IES1-Ex7-Server/utils"
	"github.com/gin-gonic/gin"
)

func PostFaceAuthHandler(appCtx utils.AppContext) gin.HandlerFunc {
	// TODO: Implement it.
	return func(ctx *gin.Context) {
		var req models.PostFaceAuthRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Request Params.",
			})
			return
		}

		slackText := "Error (不明な認証ステータスが送信されました)"
		switch req.Status {
		case "approved":
			slackText = "入室を許可しました"
		case "rejected":
			slackText = "入室を拒否しました"
		}

		body, err := json.Marshal(map[string]string{
			"text": slackText,
		})
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Sorry, server error occured.")
			log.Println("SERVER ERROR: ", err.Error())
			return
		}

		if res, err := http.Post(
			appCtx.Config.Slack.Webhook.Url,
			"application/json",
			bytes.NewBuffer(body),
		); res.StatusCode != http.StatusOK || err != nil {
			ctx.String(http.StatusInternalServerError, "Sorry, server error occured.")
			log.Printf("SERVER ERROR with status %d: %s\n", res.StatusCode, err.Error())
			return
		}

		// TODO: SQLITE Storeing

		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}
