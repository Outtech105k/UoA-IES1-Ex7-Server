package routes

import (
	"net/http"

	"github.com/Outtech105k/UoA-IES1-Ex7-Server/controllers"
	"github.com/Outtech105k/UoA-IES1-Ex7-Server/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(appCtx utils.AppContext) *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("public/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api := r.Group("/api")
	{
		api.POST("/faceauth", controllers.PostFaceAuthHandler(appCtx))
		api.GET("/history", controllers.GetDetectHistory(appCtx))
	}

	return r
}
