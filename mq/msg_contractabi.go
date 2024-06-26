package mq

import (
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	"github.com/nats-io/nats.go"
)

// ContractAbi
type ContractAbiMsg struct {
	Msg
	Data *entity.Contractabi `json:"data"`
}

// func (s *ContractAbiMsg) IsValid() bool {
// 	if s.Opt == "" || s.Id <= 0 || s.ContractAddress == "" || s.ChainId == "" {
// 		return false
// 	}
// 	return true
// }

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
				var err error
				if err = json.Unmarshal(msg.Data, &data); err == nil {
					err = fn(&data)
				}
				if err != nil {
					g.Log().Error(s.ctx, "Sub_ContractAbi err:", msg.Data, ",err:", err)
				}
				msg.Ack()
			case <-s.ctx.Done():
				sub.Unsubscribe()
				close(ch)
				sub.Drain()
			}
		}
	}()
}
