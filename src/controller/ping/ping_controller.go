package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping handdler rest response pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
