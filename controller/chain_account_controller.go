package controller

import "github.com/gin-gonic/gin"

func ChainAccountRoute(e *gin.Engine) {
	e.POST("/chain/account/importAdmin", importAdminAccountHandler)
}

func importAdminAccountHandler(c *gin.Context) {

}
