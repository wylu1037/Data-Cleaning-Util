package service

import (
	"baas-clean/model"
	"baas-clean/model/vote"
	"baas-clean/utils"
	"fmt"
	"strconv"
)

const (
	ChainNode     = "CHAIN_NODE:" // 以下是db7
	NodeVote      = "NODE_VOTE:"
	NodeVoteTally = "NODE_VOTE_TALLY:"

	BlockLtc = "BLOCK-LTC:ZLTC" // 以下是db8
	InfoLtc  = "INFO-LTC:ZLTC"

	BlockLtc2  = "BLOCK-LTC:ZLTC" // 以下是db9
	InfoLtc2   = "INFO-LTC:ZLTC"
	ChainNode2 = "CHAIN-NODE:CHAIN:NODE:STATUS"
)

// ChainDelete 删除链信息
func ChainDelete(chainId uint64) {

	// 删除链相关缓存
	var err error
	chainIdStr := strconv.FormatUint(chainId, 10)
	err = utils.LikeDelete(ChainNode+chainIdStr, 7)
	if err != nil {
		fmt.Printf("chain_service delete redis cache occured error, key is %s%s, error message is %v \n",
			ChainNode, chainIdStr, err)
		return
	}
	err = utils.LikeDelete(BlockLtc+chainIdStr, 8)
	if err != nil {
		fmt.Printf("chain_service delete redis cache occured error, key is %s%s, error message is %v \n",
			BlockLtc, chainIdStr, err)
		return
	}
	err = utils.LikeDelete(InfoLtc+chainIdStr, 8)
	if err != nil {
		fmt.Printf("chain_service delete redis cache occured error, key is %s%s, error message is %v \n",
			InfoLtc, chainIdStr, err)
		return
	}
	err = utils.LikeDelete(BlockLtc2+chainIdStr, 8)
	if err != nil {
		fmt.Printf("chain_service delete redis cache occured error, key is %s%s, error message is %v \n",
			BlockLtc2, chainIdStr, err)
		return
	}
	err = utils.LikeDelete(InfoLtc2+chainIdStr, 8)
	if err != nil {
		fmt.Printf("chain_service delete redis cache occured error, key is %s%s, error message is %v \n",
			InfoLtc2, chainIdStr, err)
		return
	}
	err = utils.LikeDelete(ChainNode2+chainIdStr, 8)
	if err != nil {
		fmt.Printf("chain_service delete redis cache occured error, key is %s%s, error message is %v \n",
			ChainNode2, chainIdStr, err)
		return
	}

	// 查找链关联的节点
	nodes := model.FindNodesByChainId(chainId)
	fmt.Printf("Find nodes by chainId = %d, return result: %v \n", chainId, nodes)
	for _, node := range *nodes {
		// 删除节点相关缓存
		nodeIdStr := strconv.FormatUint(node.ID, 10)
		err = utils.LikeDelete(NodeVote+nodeIdStr, 7)
		if err != nil {
			fmt.Printf("chain_service delete redis cache occured error, key is %s%s, error message is %v \n",
				NodeVote, nodeIdStr, err)
			return
		}
		err := utils.LikeDelete(NodeVoteTally+strconv.FormatUint(node.ID, 10), 7)
		if err != nil {
			fmt.Printf("chain_service delete redis cache occured error, key is %s%s, error message is %v \n",
				NodeVoteTally, nodeIdStr, err)
			return
		}
	}

	// 删除链数据
	model.DeleteChainById(chainId)

	// 删除关联的节点数据
	model.DeleteNodes(*nodes)

	// 删除节点关联的投票数据
	for _, node := range *nodes {
		vote.DeleteNodeVoteByNodeId(node.ID)
		vote.DeleteNodeVoteDetailsByNodeId(node.ID)
	}

	// 删除证书
	rootId := model.FindRootCAByChainId(chainId)
	model.DeleteChildCAByRootId(rootId)
	model.DeleteRootCAByChainId(chainId)

	// 查找联盟成员及权限
	model.DeleteMemberByChain(chainId)
	model.DeletePermissionsByChainId(chainId)
}
