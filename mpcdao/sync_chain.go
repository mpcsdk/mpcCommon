package mpcdao

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type ChainTransfer struct {
	redis *gredis.Redis
	dur   time.Duration
}
type QueryData struct {
	ChainId  int64    `json:"chainId"`
	From     string   `json:"from"`
	To       string   `json:"to"`
	Contract string   `json:"contract"`
	Kinds    []string `json:"kinds"`
	///
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	///
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func NewChainTransfer(redis *gredis.Redis, dur int) *ChainTransfer {
	// g.DB(dao.ChainTransfer.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	dao.ChainTransfer.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(redis))

	return &ChainTransfer{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}

func (s *ChainTransfer) Insert(ctx context.Context, data *entity.ChainTransfer) error {
	_, err := dao.ChainTransfer.Ctx(ctx).Insert(data)
	return err
}

func (s *ChainTransfer) InsertBatch(ctx context.Context, data []*entity.ChainTransfer) error {
	_, err := dao.ChainTransfer.Ctx(ctx).Insert(data)
	return err
}

func (s *ChainTransfer) Query(ctx context.Context, query *QueryData) ([]*entity.ChainTransfer, error) {
	if query.PageSize < 0 || query.Page < 0 {
		return nil, nil
	}
	//
	where := dao.ChainTransfer.Ctx(ctx)
	if query.ChainId != 0 {
		where = where.Where(dao.ChainTransfer.Columns().ChainId, query.ChainId)
	}
	if len(query.Kinds) > 0 {
		where = where.Where(dao.ChainTransfer.Columns().Kind, query.Kinds)
	}
	if query.From != "" {
		where = where.Where(dao.ChainTransfer.Columns().From, query.From)
	}
	if query.To != "" {
		where = where.Where(dao.ChainTransfer.Columns().To, query.To)
	}
	if query.Contract != "" {
		where = where.Where(dao.ChainTransfer.Columns().Contract, query.Contract)
	}
	///time
	if query.StartTime != 0 {
		where = where.WhereGTE(dao.ChainTransfer.Columns().Ts, query.StartTime)
	}
	if query.EndTime != 0 {
		where = where.WhereLTE(dao.ChainTransfer.Columns().Ts, query.EndTime)
	}
	///
	if query.PageSize != 0 {
		where = where.Limit(query.Page*query.PageSize, query.PageSize)
	}
	///
	result, err := where.All()
	if err != nil {
		return nil, err
	}
	data := []*entity.ChainTransfer{}
	err = result.Structs(&data)
	///
	return data, err
}
