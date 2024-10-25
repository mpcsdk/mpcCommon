package mpcdao

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type RelayerAdmin struct {
	redis *gredis.Redis
	dur   time.Duration
	/////

}

func NewRelayerAdmin(redis *gredis.Redis, dur int) *RelayerAdmin {
	if redis != nil {
		g.DB(dao.RelayerChannel.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
		g.DB(dao.RelayerFee.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	}

	return &RelayerAdmin{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}
func (s *RelayerAdmin) RelayerAdminChannel(ctx context.Context, appId string) (*entity.RelayerChannel, error) {
	rst, err := dao.RelayerChannel.Ctx(ctx).Where(dao.RelayerChannel.Columns().ChannelId, appId).One()
	if err != nil {
		return nil, err
	}
	var entity *entity.RelayerChannel = nil
	err = rst.Struct(&entity)
	if err != nil {
		return nil, err
	}
	////
	return entity, nil
}
func (s *RelayerAdmin) RelayerAdminChannelFee(ctx context.Context, appId string) (*entity.RelayerFee, error) {
	rst, err := dao.RelayerFee.Ctx(ctx).Where(dao.RelayerFee.Columns().ChannelId, appId).One()
	if err != nil {
		return nil, err
	}
	var entity *entity.RelayerFee = nil
	err = rst.Struct(&entity)
	if err != nil {
		return nil, err
	}
	////
	return entity, nil
}
