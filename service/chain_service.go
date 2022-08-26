package service

import (
	"baas-clean/model"
	"baas-clean/utils"
	"strconv"
)

const (
	ChainNode     = "CHAIN_NODE:"
	NodeVote      = "NODE_VOTE:"
	NodeVoteTally = "NODE_VOTE_TALLY:"

	//
)

// ChainDelete 删除链信息
func ChainDelete(chainId uint64) {

	// 删除缓存
	utils.LikeDelete(ChainNode+strconv.FormatUint(chainId, 10), 7)

	// 查找链关联的节点
	nodes := model.FindNodesByChainId(chainId)
	for _, node := range *nodes {
		utils.LikeDelete(NodeVote+strconv.FormatUint(node.ID, 10), 7)
		utils.LikeDelete(NodeVoteTally+strconv.FormatUint(node.ID, 10), 7)
	}

	// model.DeleteChainById(chainId)

	// 查找链关联的节点

	// 查找联盟成员及权限

	// 查找证书

	//
}
