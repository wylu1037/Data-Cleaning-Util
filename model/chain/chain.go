package chain

import (
	"baas-clean/config"
	"baas-clean/util"
	"fmt"
)

type Chain struct {
	ID           uint64 `gorm:"primaryKey"`
	BlockChainId string
	Version      string
	Name         string
	ChainType    uint8 // 链类型：0-晶格链
}

// TableName 实现接口重写表名
func (c Chain) TableName() string {
	return "chain_info"
}

// QueryChainInfoById 查询链信息
func QueryChainInfoById(chainId uint64) (*Chain, error) {
	var chain Chain
	config.MySqlDB.Find(&chain, chainId)
	fmt.Println(util.ToStr(chain))
	return &chain, nil
}

// DeleteChainById 删除链数据
func DeleteChainById(chainId uint64) {
	fmt.Printf("delete chain by id = %d \n", chainId)
	config.MySqlDB.Delete(&Chain{}, chainId).Limit(1)
}
