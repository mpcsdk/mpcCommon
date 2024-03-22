package mq

import (
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/nats-io/nats.go"
)

// ContractAbi
type ContractAbiMsg struct {
	Opt             string `json:"opt"`
	Id              int64  `json:"id"`
	ContractAddress string `json:"contractAddress"`
	ChainId         string `json:"chainId"`
}

func (s *ContractAbiMsg) IsValid() bool {
	if s.Opt == "" || s.Id <= 0 || s.ContractAddress == "" || s.ChainId == "" {
		return false
	}
	return true
}

func (s *NatsServer) Sub_ContractAbi(subj string, fn func(data *ContractAbiMsg) error) {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanSubscribe(subj, ch)
	if err != nil {
		panic(err)
	}
	//
	go func() {
		for {
			select {
			case msg := <-ch:
				var data ContractAbiMsg
				if err := json.Unmarshal(msg.Data, &data); err != nil {
					g.Log().Error(s.ctx, "Sub_ContractAbi Unmarshal:", msg.Data, ",err:", err)
					continue
				}
				err = fn(&data)
				if err != nil {
					g.Log().Error(s.ctx, "Sub_ContractAbi fn:", err)
					continue
				}
			case <-s.ctx.Done():
				sub.Unsubscribe()
				close(ch)
				sub.Drain()
			}
		}
	}()
}
