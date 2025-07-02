package routes

import (
	"github.com/Outtech105k/UoA-IES1-Ex7-Server/controllers"
	"github.com/Outtech105k/UoA-IES1-Ex7-Server/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(appCtx utils.AppContext) *gin.Engine {
	r := gin.Default()

	r.POST("/faceauth", controllers.PostFaceAuthHandler(appCtx))
	r.GET("/history", controllers.GetDetectHistory())

	return r
}
