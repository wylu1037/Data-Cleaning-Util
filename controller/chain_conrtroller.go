package controller

import (
	"baas-clean/model/certificate"
	"baas-clean/model/chain"
	"baas-clean/model/node"
	"baas-clean/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ChainRoute(e *gin.Engine) {
	e.GET("/chain/query", queryChainInfoByIdHandler)
	e.GET("/chain/delete/:chainId", deleteChainInfoHandler)
	e.GET("/chain/findNodes", findNodesByChainId)
	e.GET("/certificate/findRootCA", findRootCAByChainId)
}

// 查询链信息
func queryChainInfoByIdHandler(c *gin.Context) {
	chainIdStr := c.Query("chainId")
	num, _ := strconv.Atoi(chainIdStr)
	chainId := uint64(num)
	data, err := chain.QueryChainInfoById(chainId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "查询成功！",
			"data":    data,
		})
	}
}

// 删除链及相关信息
func deleteChainInfoHandler(c *gin.Context) {
	num, _ := strconv.Atoi(c.Param("chainId"))
	chinId := uint64(num)

	service.ChainDelete(chinId)

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
		"data":    node.FindNodesByChainId(chainId),
	})
}

func findRootCAByChainId(c *gin.Context) {
	chainIdStr := c.Query("chainId")
	num, _ := strconv.Atoi(chainIdStr)
	chainId := uint64(num)

	rootId := certificate.FindRootCAByChainId(chainId)
	c.JSON(http.StatusOK, gin.H{
		"message": "查询根证书成功！",
		"data":    rootId,
	})
}
