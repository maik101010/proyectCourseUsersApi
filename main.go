package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maik101010/proyectCourseUsersApi/app"
)

var (
	router = gin.Default()
)

func main() {  
	app.StartApplication()
}