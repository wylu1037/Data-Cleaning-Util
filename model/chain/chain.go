package chain

import (
	"clear-chain/config"
	"clear-chain/util/json"
	"github.com/sirupsen/logrus"
)

type Chain struct {
	ID           uint64
	BlockChainId string `gorm:"column:blockchainId"`
	Version      string
	Name         string
	ChainType    int8   `gorm:"column:chainType"`
	AccountId    uint64 `gorm:"column:accountId"`
}

// TableName 实现接口重写表名
func (c Chain) TableName() string {
	return "chain_info"
}

// FindChainInfoById 查询链信息
func FindChainInfoById(chainId uint64) (*Chain, error) {
	logrus.Infof("[chain] FindChainInfoById() called with: chainId = %d", chainId)

	var chain Chain
	config.MySqlDB.First(&chain, chainId)

	return &chain, nil
}

// DeleteChainById 删除链数据
func DeleteChainById(chainId uint64) {
	logrus.Infof("[chain] DeleteChainById() called with: chainId = %d", chainId)

	config.MySqlDB.Delete(&Chain{}, chainId).Limit(1)
}

type HyperledgerUpChainRecord struct {
	ID      uint64
	ChainId uint64 `gorm:"column:chainId"`
}

func (HyperledgerUpChainRecord) TableName() string {
	return "fabric_chain_record"
}

func DeleteHyperledgerUpChainRecordByChainId(chainId uint64) {
	logrus.Infof("[chain] DeleteHyperledgerUpChainRecordByChainId() called with: chainId = %d", chainId)

	config.MySqlDB.Where("chainId = ?", chainId).Delete(&HyperledgerUpChainRecord{})
}

// FindRangeChainInfo 查询指定范围内的链信息
func FindRangeChainInfo(begin, end uint64) []Chain {
	logrus.Infof("[chain] FindRangeChainInfo() called with: begin = %d, end = %d", begin, end)

	var chains []Chain
	config.MySqlDB.Where("id >= ? and id <= ?", begin, end).Find(&chains)
	return chains
}

type PageChainListReq struct {
	Name string
	Page uint64
	Size uint64
}

type PageChainListResult struct {
	Total uint64
	List  []Chain
}

// FindPageChainList 分页查询查询链列表
func FindPageChainList(req PageChainListReq) PageChainListResult {
	logrus.Infof("[chain] FindPageChainList() called with: req = %v", json.ToStr(req))

	result := PageChainListResult{}
	total := findChainTotal()
	result.Total = total
	if total == 0 {
		result.List = nil
		return result
	}

	var chainList []Chain
	config.MySqlDB.Limit(req.Size).Offset(req.Page - 1).Find(&chainList)
	result.List = chainList
	return result
}

// 查询总数
func findChainTotal() uint64 {
	var total uint64
	config.MySqlDB.Model(&Chain{}).Where("id > ?", 0).Count(&total)
	return total
}
