package node

import (
	"clear-chain/config"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Node struct {
	ID      uint64 // nodeId
	Name    string
	ChainId uint64 // chainId
	IP      string
}

func (n Node) TableName() string {
	return "chain_node"
}

// FindNodesByChainId 根据链id查找节点列表
func FindNodesByChainId(chainId uint64) *[]Node {
	var nodes []Node
	config.MySqlDB.Where("chainId = ?", chainId).Find(&nodes)

	nodesData, _ := json.Marshal(nodes)
	fmt.Println(string(nodesData))

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
