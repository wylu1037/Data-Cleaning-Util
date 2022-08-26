package controller

import (
	"baas-clean/model/certificate"
	"baas-clean/model/chain"
	"baas-clean/model/node"
	"baas-clean/service"
	"baas-clean/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChainRoute(e *gin.Engine) {
	e.GET("/chain/query", queryChainInfoByIdHandler)
	e.GET("/chain/delete/:chainId", deleteChainInfoHandler)
	e.GET("/chain/findNodes", findNodesByChainId)
	e.GET("/certificate/findRootCA", findRootCAByChainId)
}

// 查询链信息
func queryChainInfoByIdHandler(c *gin.Context) {
	chainId := util.Str2uint64(c.Query("chainId"))
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
	chinId := util.Str2uint64(c.Query("chainId"))

	service.ChainDelete(chinId)

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功！",
		"data":    nil,
	})
}

// 根据链id查询节点列表
func findNodesByChainId(c *gin.Context) {
	chainId := util.Str2uint64(c.Query("chainId"))

	c.JSON(http.StatusOK, gin.H{
		"message": "查询节点列表成功！",
		"data":    node.FindNodesByChainId(chainId),
	})
}

func findRootCAByChainId(c *gin.Context) {
	chainId := util.Str2uint64(c.Query("chainId"))

	rootId := certificate.FindRootCAByChainId(chainId)
	c.JSON(http.StatusOK, gin.H{
		"message": "查询根证书成功！",
		"data":    rootId,
	})
}
