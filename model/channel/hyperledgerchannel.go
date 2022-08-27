package channel

import (
	"clear-chain/config"
	"github.com/sirupsen/logrus"
)

type Channel struct {
	ID          uint64
	ChannelName string
	ChainId     uint64
	DeleteFlag  uint8 `json:"_"`
}

func (Channel) TableName() string {
	return "fabric_chain_channel"
}

// FindChannels 根据链查找通道
func FindChannels(chainId uint64) []Channel {
	logrus.Infof("[hyperledgerchannel] FindChannels() called with: chainId = %d", chainId)

	var channels []Channel
	config.MySqlDB.Where("chainId = ?", chainId).Find(&channels)
	return channels
}

// DeleteChannelByChainId 根据链删除通道
func DeleteChannelByChainId(chainId uint64) {
	logrus.Infof("[hyperledgerchannel] DeleteChannelByChainId() called with: chainId = %d", chainId)

	config.MySqlDB.Where("chainId = ?", chainId).Delete(&Channel{})
}
