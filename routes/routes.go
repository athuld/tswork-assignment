package routes

import (
	"tswork-mongo/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.POST("/api/v1/upload", controllers.UploadCsv())
	router.GET("/api/v1/getstock/:date", controllers.GetStockByDate())

	router.Run("localhost:5000")
}
