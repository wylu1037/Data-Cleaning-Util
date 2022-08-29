package certificate

import (
	"clear-chain/config"
	"github.com/sirupsen/logrus"
)

type RootCA struct {
	ID         uint64 `gorm:"primaryKey"`
	ChainId    uint64 `gorm:"column:chainId"`
	DeleteFlag uint8  `gorm:"column:deleteFlag"`
}

func (RootCA) TableName() string {
	return "root_cert_config"
}

// FindRootCAByChainId 根据链查找根证书
func FindRootCAByChainId(chainId uint64) uint64 {
	var rootCA RootCA
	config.MySqlDB.Where("chainId = ? and deleteFlag = ?", chainId, 0).First(&rootCA)
	return rootCA.ID
}

// DeleteRootCAByChainId 根据链删除根证书
func DeleteRootCAByChainId(chainId uint64) {
	logrus.Infof("[certificate] DeleteRootCAByChainId() called with: chainId = %d", chainId)

	config.MySqlDB.Where("chainId = ?", chainId).Delete(&RootCA{})
}

type ChildCA struct {
	ID     uint64
	RootId uint64 `gorm:"column:rootId"`
}

func (ChildCA) TableName() string {
	return "chain_node_cert_config"
}

// DeleteChildCAByRootId 根据根证书删除子证书
func DeleteChildCAByRootId(rootId uint64) {
	logrus.Infof("[certificate] DeleteChildCAByRootId() called with: rootId = %d", rootId)

	config.MySqlDB.Where("rootId = ?", rootId).Delete(&ChildCA{})
}

type HyperledgerCertificate struct {
	ID      uint64
	ChainId uint64 `gorm:"column:chainId"`
}

func (HyperledgerCertificate) TableName() string {
	return "fabric_node_cert"
}

func DeleteHyperledgerCertificateByChainId(chainId uint64) {
	logrus.Infof("[certificate] DeleteHyperledgerCertificateByChainId() called with: chainId = %d", chainId)

	config.MySqlDB.Where("chainId = ?", chainId).Delete(&HyperledgerCertificate{})
}
