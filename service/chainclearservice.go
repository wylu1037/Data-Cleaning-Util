package service

import (
	"clear-chain/model/certificate"
	"clear-chain/model/chain"
	"clear-chain/model/member"
	"clear-chain/model/node"
	"clear-chain/model/vote"
	"clear-chain/util"
	"github.com/sirupsen/logrus"
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

// DeleteChain 删除链信息
func DeleteChain(chainId uint64) {
	logrus.Infof("[chainclearservice] DeleteChain() called with: chainId = %d", chainId)

	chainInfo, err := chain.FindChainInfoById(chainId)
	if err != nil {
		return
	}
	if chainInfo != nil && chainInfo.ChainType == 0 {
		deleteLatticeChain(chainId)
	} else {
		deleteHyperledger(chainId)
	}
}

// 删除晶格链
func deleteLatticeChain(chainId uint64) {
	logrus.Infof("[chainClearService] deleteLatticeChain() called with: chainId = %d", chainId)
	// 查找链关联的节点
	nodes := node.FindNodesByChainId(chainId)

	// 删除缓存
	deleteCache(chainId, *nodes)

	// 删除链数据
	chain.DeleteChainById(chainId)

	// 删除关联的节点数据
	node.DeleteNodes(*nodes)

	// 删除节点关联的投票数据
	for _, nodeInfo := range *nodes {
		vote.DeleteNodeVoteByNodeId(nodeInfo.ID)
		vote.DeleteNodeVoteDetailsByNodeId(nodeInfo.ID)
	}

	// 删除证书
	rootId := certificate.FindRootCAByChainId(chainId)
	certificate.DeleteChildCAByRootId(rootId)
	certificate.DeleteRootCAByChainId(chainId)

	// 查找联盟成员及权限
	member.DeleteMemberByChain(chainId)
	member.DeletePermissionsByChainId(chainId)
}

// 删除超级账本
func deleteHyperledger(chainId uint64) {
	logrus.Infof("[chainclearservice] deleteHyperledger() called with: chainId = %d", chainId)
}

// 删除缓存
func deleteCache(chainId uint64, nodes []node.Node) {
	logrus.Infof("[chainclearservice] deleteCache() called with: chainId = %d", chainId)

	// 删除链相关缓存
	var err error
	chainIdStr := strconv.FormatUint(chainId, 10)
	err = util.LikeDelete(ChainNode+chainIdStr, 7)
	if err != nil {
		logrus.Errorf("[chainclearservice] deleteCache() delete redis cache occured error, key is %s%s, error message is %v \n",
			ChainNode, chainIdStr, err)
		return
	}
	err = util.LikeDelete(BlockLtc+chainIdStr, 8)
	if err != nil {
		logrus.Errorf("[chainclearservice] deleteCache() delete redis cache occured error, key is %s%s, error message is %v \n",
			BlockLtc, chainIdStr, err)
		return
	}
	err = util.LikeDelete(InfoLtc+chainIdStr, 8)
	if err != nil {
		logrus.Errorf("[chainclearservice] deleteCache() delete redis cache occured error, key is %s%s, error message is %v \n",
			InfoLtc, chainIdStr, err)
		return
	}
	err = util.LikeDelete(BlockLtc2+chainIdStr, 8)
	if err != nil {
		logrus.Errorf("[chainclearservice] deleteCache() delete redis cache occured error, key is %s%s, error message is %v \n",
			BlockLtc2, chainIdStr, err)
		return
	}
	err = util.LikeDelete(InfoLtc2+chainIdStr, 8)
	if err != nil {
		logrus.Errorf("[chainclearservice] deleteCache() delete redis cache occured error, key is %s%s, error message is %v \n",
			InfoLtc2, chainIdStr, err)
		return
	}
	err = util.LikeDelete(ChainNode2+chainIdStr, 8)
	if err != nil {
		logrus.Errorf("[chainclearservice] deleteCache() delete redis cache occured error, key is %s%s, error message is %v \n",
			ChainNode2, chainIdStr, err)
		return
	}

	logrus.Infof("[chainclearservice] find nodes by chainId = %d, return result = %v", chainId, nodes)
	for _, nodeInfo := range nodes {
		// 删除节点相关缓存
		nodeIdStr := strconv.FormatUint(nodeInfo.ID, 10)
		err = util.LikeDelete(NodeVote+nodeIdStr, 7)
		if err != nil {
			logrus.Errorf("[chainclearservice] delete redis cache occured error, key is %s%s, error message is %v \n",
				NodeVote, nodeIdStr, err)
			return
		}
		err := util.LikeDelete(NodeVoteTally+strconv.FormatUint(nodeInfo.ID, 10), 7)
		if err != nil {
			logrus.Errorf("[chainclearservice] delete redis cache occured error, key is %s%s, error message is %v \n",
				NodeVoteTally, nodeIdStr, err)
			return
		}
	}
}
