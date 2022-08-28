package contract

import (
	"clear-chain/config"
	"clear-chain/model/channel"
	"clear-chain/util/json"
	"github.com/sirupsen/logrus"
)

type HyperledgerContract struct {
	ID        uint64
	ChannelId uint64
}

func (HyperledgerContract) TableName() string {
	return "fabric_contract_info"
}

func DeleteHyperledgerContracts(channels []channel.Channel) {
	logrus.Infof("[hyperledgercontract] DeleteHyperledgerContracts() called with: channels = %v", json.ToStr(channels))

	if channels == nil || len(channels) == 0 {
		return
	}

	for _, c := range channels {
		deleteHyperledgerContractByChannelId(c.ID)
	}
}

func deleteHyperledgerContractByChannelId(channelId uint64) {
	logrus.Infof("[hyperledgercontract] deleteHyperledgerContractByChannelId() called with: channelId = %d",
		channelId)

	config.MySqlDB.Where("channelId = ?", channelId).Delete(&HyperledgerContract{})
}
