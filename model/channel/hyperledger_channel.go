package channel

import (
	"clear-chain/config"
	"clear-chain/util/json"
	"github.com/sirupsen/logrus"
)

type Channel struct {
	ID          uint64
	ChannelName string `gorm:"column:channelName"`
	ChainId     uint64 `gorm:"column:chainId"`
	DeleteFlag  uint8  `gorm:"column:deleteFlag"`
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

type Member struct {
	ID        uint64
	ChannelId uint64 `gorm:"column:channelId"`
}

func (Member) TableName() string {
	return "fabric_chain_channel_member"
}

func DeleteChannelMembers(channels []Channel) {
	logrus.Infof("[hyperledgerchannel] DeleteChannelMembers() called with: channels = %v", json.ToStr(channels))

	if channels == nil || len(channels) == 0 {
		return
	}

	for _, channel := range channels {
		deleteChannelMember(channel.ID)
	}
}

func deleteChannelMember(channelId uint64) {
	logrus.Infof("[hyperledgerchannel] deleteChannelMember() called with: channelId = %d", channelId)

	config.MySqlDB.Where("channelId = ?", channelId).Delete(&Member{}, channelId)
}

type AuditChannel struct {
	ID        uint64
	ChannelId uint64 `gorm:"column:channelId"`
	AccountId uint64 `gorm:"column:accountId"`
}

func (AuditChannel) TableName() string {
	return "fabric_chain_channel_audit"
}

func DeleteAuditChannels(channels []Channel) {
	logrus.Infof("[hyperledgerchannel] DeleteAuditChannels() called with: channels = %v", json.ToStr(channels))
	if channels == nil || len(channels) == 0 {
		return
	}

	for _, channel := range channels {
		deleteAuditChannel(channel.ID)
	}
}

func deleteAuditChannel(channelId uint64) {
	logrus.Infof("[hyperledgerchannel] deleteAuditChannel() called with: channelId = %d", channelId)

	config.MySqlDB.Where("channelId = ?", channelId).Delete(&AuditChannel{})
}
