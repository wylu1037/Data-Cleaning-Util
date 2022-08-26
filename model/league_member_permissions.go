package model

import (
	"baas-clean/config"
	"fmt"
)

type MemberPermissions struct {
	UserId  uint64 // 指向联盟成员id
	ChainId uint64
}

func (MemberPermissions) TableName() string {
	return "chain_league_user_permission"
}

// DeletePermissionsByChainId 根据链删除联盟成员权限
func DeletePermissionsByChainId(chainId uint64) {
	fmt.Printf("delete permissions by chainId = %d \n", chainId)
	config.MySqlDB.Where("chainId = ?", chainId).Delete(&MemberPermissions{})
}
