package mpcdao

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type Fcm struct {
	redis *gredis.Redis
	dur   time.Duration
}
type QueryFcmToken struct {
	FcmToken string `json:"fcm_token"`
	Token    string `json:"token"`
	Address  string `json:"address"`
}

func NewFcm(redis *gredis.Redis, dur int) *Fcm {
	// g.DB(dao.ChainTransfer.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	dao.FcmToken.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(redis))

	return &Fcm{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}

func (s *Fcm) InsertFcmToken(ctx context.Context, data *entity.FcmToken) error {
	_, err := dao.FcmToken.Ctx(ctx).Insert(data)
	return err
}
func (s *Fcm) InsertPushErr(ctx context.Context, data *entity.PushErr) error {
	return nil
}

type PosFcmToken *entity.FcmToken

func (s *Fcm) QueryFcmTokenAll(ctx context.Context, pos PosFcmToken, limit int) ([]*entity.FcmToken, PosFcmToken, error) {
	where := dao.ChainTransfer.Ctx(ctx)
	if pos != nil {
		where = where.WhereGTE(dao.FcmToken.Columns().Address, pos.Address)
		where = where.WhereGTE(dao.FcmToken.Columns().FcmToken, pos.FcmToken)
	}
	///
	if limit > 0 {
		where = where.Limit(limit)
	}
	///
	result, err := where.All()
	if err != nil {
		return nil, nil, err
	}
	data := []*entity.FcmToken{}
	err = result.Structs(&data)
	///
	return data, data[len(data)-1], err
}
func (s *Fcm) QueryFcmToken(ctx context.Context, query *QueryFcmToken) ([]*entity.FcmToken, error) {
	//
	where := dao.ChainTransfer.Ctx(ctx)
	where = where.WhereGTE(dao.FcmToken.Columns().Address, query.Address)
	///

	///
	result, err := where.All()
	if err != nil {
		return nil, err
	}
	data := []*entity.FcmToken{}
	err = result.Structs(&data)
	///
	return data, err
}
