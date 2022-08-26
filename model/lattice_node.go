package model

import (
	"baas-clean/config"
	"encoding/json"
	"fmt"
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
