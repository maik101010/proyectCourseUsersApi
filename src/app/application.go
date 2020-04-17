package app

import (
	"github.com/gin-gonic/gin"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/logger"
)

var (
	router = gin.Default()
)

//StartApplication methods
func StartApplication() {
	logger.Info("start the application ")
	mapUrls()
	router.Run(":8081")
}
