package controller

import (
	"baas-clean/model"
	"baas-clean/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ChainRoute(e *gin.Engine) {
	e.GET("/chain/delete/:chainId", deleteChainInfoHandler)
	e.GET("/chain/findNodes", findNodesByChainId)
}

// 删除链及相关信息
func deleteChainInfoHandler(c *gin.Context) {
	num, _ := strconv.Atoi(c.Param("chainId"))
	chinId := uint64(num)

	service.DeleteChain(chinId)

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功！",
		"data":    nil,
	})
}

// 根据链id查询节点列表
func findNodesByChainId(c *gin.Context) {
	chainIdStr := c.Query("chainId")
	num, _ := strconv.Atoi(chainIdStr)
	chainId := uint64(num)

	c.JSON(http.StatusOK, gin.H{
		"message": "查询节点列表成功！",
		"data":    model.FindNodesByChainId(chainId),
	})
}
