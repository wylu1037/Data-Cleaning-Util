package node

import (
	"clear-chain/config"
	"github.com/sirupsen/logrus"
)

type Node struct {
	ID      uint64
	Name    string
	ChainId uint64
	IP      string
}

func (n Node) TableName() string {
	return "chain_node"
}

// FindNodesByChainId 根据链id查找节点列表
func FindNodesByChainId(chainId uint64) *[]Node {
	var nodes []Node
	config.MySqlDB.Where("chainId = ?", chainId).Find(&nodes)

	return &nodes
}

// DeleteNodes 删除节点
func DeleteNodes(nodes []Node) {
	for _, node := range nodes {
		deleteNodeById(node.ID)
	}
}

// 根据主键删除节点
func deleteNodeById(nodeId uint64) {
	logrus.Infof("[node] deleteNodeById() called with: nodeId = %d", nodeId)

	config.MySqlDB.Delete(&Node{}, nodeId).Limit(1)
}

type HyperledgerNode struct {
	ID         uint64
	ChainId    uint64
	OrgId      uint64
	HostId     uint64
	NodeType   uint8
	NodeUserId uint64
}

func (HyperledgerNode) TableName() string {
	return "fabric_chain_node"
}

// FindHyperledgerNodesByChainId 根据链查找超级账本节点
func FindHyperledgerNodesByChainId(chainId uint64) []HyperledgerNode {
	logrus.Infof("[node] FindHyperledgerNodesByChainId() called with: chainId = %d", chainId)

	var nodes []HyperledgerNode
	config.MySqlDB.Where("chainId = ? and deleteFlag = ?", chainId, 0).Find(&nodes)

	return nodes
}

// DeleteHyperledgerNodes 批量删除超级账本节点
func DeleteHyperledgerNodes(nodes []HyperledgerNode) {
	for _, nodeInfo := range nodes {
		deleteHyperledgerNode(nodeInfo.ID)
	}
}

func deleteHyperledgerNode(nodeId uint64) {
	logrus.Infof("[node] deleteHyperledgerNode() called with: nodeId = %d", nodeId)

	config.MySqlDB.Delete(&HyperledgerNode{}, nodeId)
}
