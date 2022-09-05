package controller

import (
	"clear-chain/model/chain"
	"clear-chain/model/node"
	"clear-chain/service"
	"clear-chain/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"net/http"
)

func ChainRoute(e *gin.Engine) {
	e.GET("/chain/query", queryChainInfoByIdHandler)
	e.GET("/chain/delete/:chainId", deleteChainInfoHandler)
	e.GET("/chain/findNodes", findNodesByChainId)
	e.GET("/chain/range/delete/:begin/:end", deleteRangeChainInfoHandler)
	e.GET("/chain/getAll", findAllChainHandler)
	e.POST("/chain/findPageChainList", findPageChainListHandler)
}

// 查询链信息
func queryChainInfoByIdHandler(c *gin.Context) {
	chainId := util.Str2uint64(c.Query("chainId"))
	data, err := chain.FindChainInfoById(chainId)
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
	chainId := util.Str2uint64(c.Param("chainId"))

	logrus.Infof("[chainconrtroller] deleteChainInfoHandler() called with: chainId = %d", chainId)

	service.DeleteChain(chainId)

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功！",
		"data":    nil,
	})
}

// 根据链id查询节点列表
func findNodesByChainId(c *gin.Context) {
	chainId := util.Str2uint64(c.Query("chainId"))

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "查询节点列表成功！",
		"data":    node.FindNodesByChainId(chainId),
	})
}

func deleteRangeChainInfoHandler(c *gin.Context) {
	begin := util.Str2uint64(c.Param("begin"))
	end := util.Str2uint64(c.Param("end"))

	service.DeleteRangeChain(begin, end)

	c.JSON(http.StatusOK, gin.H{
		"message": "删除范围内的链成功！",
		"data":    nil,
	})
}

func findAllChainHandler(c *gin.Context) {
	chains := chain.FindRangeChainInfo(1, math.MaxInt64)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "查找所有链信息成功！",
		"data":    chains,
	})
}

func findPageChainListHandler(c *gin.Context) {
	var req chain.PageChainListReq
	if err := c.ShouldBind(&req); err == nil {
		data := chain.FindPageChainList(req)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    data,
			"message": "查询列表成功！",
		})
	} else {
		c.JSON(http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "error": err.Error()})
	}
}
