package vote

import (
	"baas-clean/config"
	"fmt"
)

type NodeVote struct {
	ID     uint64
	NodeId uint64
}

func (NodeVote) TableName() string {
	return "chain_node_vote"
}

// DeleteNodeVoteByNodeId 根据节点删除节点投票主题数据
func DeleteNodeVoteByNodeId(nodeId uint64) {
	fmt.Printf("delete node vote by nodeId = %d \n", nodeId)
	config.MySqlDB.Where("nodeId = ?", nodeId).Delete(&NodeVote{})
}

type NodeVoteDetails struct {
	ID     uint64
	NodeId uint64
}

func (NodeVoteDetails) TableName() string {
	return "chain_node_vote_detail"
}

// DeleteNodeVoteDetailsByNodeId 根据节点删除节点投票详情数据
func DeleteNodeVoteDetailsByNodeId(nodeId uint64) {
	fmt.Printf("delete node vote details by nodeId = %d\n", nodeId)
	config.MySqlDB.Where("nodeId = ?", nodeId).Delete(&NodeVoteDetails{})
}
