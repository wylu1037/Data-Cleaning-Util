package controller

import "github.com/gin-gonic/gin"
import "net/http"

func HelloRoute(e *gin.Engine) {
	e.GET("/hello/say", sayHandler)
}

func sayHandler(c *gin.Context) {
	message := c.DefaultQuery("message", "unknow")
	c.JSONP(http.StatusOK, gin.H{
		"method": "say",
		"data":   message,
	})
}
