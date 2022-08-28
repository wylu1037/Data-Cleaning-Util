package vote

import (
	"clear-chain/config"
	"github.com/sirupsen/logrus"
)

type NodeVote struct {
	ID     uint64
	NodeId uint64 `gorm:"column:nodeId"`
}

func (NodeVote) TableName() string {
	return "chain_node_vote"
}

// DeleteNodeVoteByNodeId 根据节点删除节点投票主题数据
func DeleteNodeVoteByNodeId(nodeId uint64) {
	logrus.Infof("[vote] DeleteNodeVoteByNodeId() called with: nodeId = %d", nodeId)

	config.MySqlDB.Where("nodeId = ?", nodeId).Delete(&NodeVote{})
}

type NodeVoteDetails struct {
	ID     uint64
	NodeId uint64 `gorm:"column:nodeId"`
}

func (NodeVoteDetails) TableName() string {
	return "chain_node_vote_detail"
}

// DeleteNodeVoteDetailsByNodeId 根据节点删除节点投票详情数据
func DeleteNodeVoteDetailsByNodeId(nodeId uint64) {
	logrus.Infof("[vote] DeleteNodeVoteDetailsByNodeId() called with: nodeId = %d", nodeId)

	config.MySqlDB.Where("nodeId = ?", nodeId).Delete(&NodeVoteDetails{})
}
