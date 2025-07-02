package routes

import (
	"github.com/Outtech105k/UoA-IES1-Ex7-Server/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/faceauth", controllers.PostFaceAuthHandler())

	return r
}
