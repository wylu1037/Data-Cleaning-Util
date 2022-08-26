package model

import (
	"baas-clean/config"
	"fmt"
)

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
