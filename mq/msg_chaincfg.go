package mq

import (
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	"github.com/nats-io/nats.go"
)

// /Sub_ChainCfg
type ChainCfgMsg struct {
	Msg
	Data *entity.Chaincfg `json:"data"`
}

func (s *NatsServer) Sub_ChainCfg(subj string, fn func(data *ChainCfgMsg) error) {
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
				var data ChainCfgMsg
				var err error
				if err = json.Unmarshal(msg.Data, &data); err == nil {
					err = fn(&data)
				}
				if err != nil {
					g.Log().Error(s.ctx, "Sub_ChainCfg err:", msg.Data, ",err:", err)
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
