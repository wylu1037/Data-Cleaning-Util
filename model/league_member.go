package model

import (
	"baas-clean/config"
	"fmt"
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
	fmt.Printf("delete member by chainId = %d \n", chainId)
	config.MySqlDB.Where("chainId = ?", chainId).Delete(&LeagueMember{})
}
