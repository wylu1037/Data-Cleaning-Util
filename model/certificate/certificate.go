package certificate

import (
	"baas-clean/config"
	"fmt"
)

type RootCA struct {
	ID         uint64 `gorm:"primaryKey"`
	ChainId    uint64
	DeleteFlag uint8
}

func (RootCA) TableName() string {
	return "root_cert_config"
}

// FindRootCAByChainId 根据链查找根证书
func FindRootCAByChainId(chainId uint64) uint64 {
	var rootCA RootCA
	config.MySqlDB.Where("chainId = ? and deleteFlag = ?", chainId, 0).First(&rootCA)
	fmt.Printf("find rootCA by chainId = %d, return rootId = %d \n", chainId, rootCA.ID)
	return rootCA.ID
}

// DeleteRootCAByChainId 根据链删除根证书
func DeleteRootCAByChainId(chainId uint64) {
	fmt.Printf("delete rootCA by chainId = %d \n", chainId)
	config.MySqlDB.Where("chainId = ?", chainId).Delete(&RootCA{})
}

type ChildCA struct {
	ID     uint64
	rootId uint64
}

func (ChildCA) TableName() string {
	return "chain_node_cert_config"
}

// DeleteChildCAByRootId 根据根证书删除子证书
func DeleteChildCAByRootId(rootId uint64) {
	fmt.Printf("delete chain childCA by rootId = %d \n", rootId)
	config.MySqlDB.Where("rootId = ?", rootId).Delete(&ChildCA{})
}
