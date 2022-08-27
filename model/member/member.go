package member

import (
	"clear-chain/config"
	"github.com/sirupsen/logrus"
)

type LeagueMember struct {
	ID      uint64 // 联盟成员id
	ChainId uint64
	NodeId  uint64
}

func (LeagueMember) TableName() string {
	return "chain_league_user"
}

// FindMembersByChainId 根据链查找联盟成员
func FindMembersByChainId(chainId uint64) []LeagueMember {
	var members []LeagueMember
	config.MySqlDB.Where(&LeagueMember{ChainId: chainId}).Find(&members)
	return members
}

func DeleteMemberByChain(chainId uint64) {
	logrus.Infof("[member] DeleteMemberByChain() called with: chainId = %d", chainId)

	config.MySqlDB.Where("chainId = ?", chainId).Delete(&LeagueMember{})
}

type Permissions struct {
	UserId  uint64 // 指向联盟成员id
	ChainId uint64
}

func (Permissions) TableName() string {
	return "chain_league_user_permission"
}

// DeletePermissionsByChainId 根据链删除联盟成员权限
func DeletePermissionsByChainId(chainId uint64) {
	logrus.Infof("[member] DeletePermissionsByChainId() called with: chainId = %d", chainId)

	config.MySqlDB.Where("chainId = ?", chainId).Delete(&Permissions{})
}
