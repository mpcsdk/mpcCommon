package mpcdao

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type ChainData struct {
	redis *gredis.Redis
	dur   time.Duration
}
type QueryData struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Contract string `json:"contract"`
	///
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	///
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func NewChainData(redis *gredis.Redis, dur int) *ChainData {
	// g.DB(dao.ChainData.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	dao.ChainData.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(redis))

	return &ChainData{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}

func (s *ChainData) Insert(ctx context.Context, data *entity.ChainData) error {
	_, err := dao.ChainData.Ctx(ctx).Insert(data)
	return err
}

func (s *ChainData) Query(ctx context.Context, query *QueryData) ([]*entity.ChainData, error) {
	if query.PageSize < 0 || query.Page < 0 {
		return nil, nil
	}
	//
	where := dao.ChainData.Ctx(ctx)
	if query.From != "" {
		where = where.Where(dao.ChainData.Columns().From, query.From)
	}
	if query.To != "" {
		where = where.Where(dao.ChainData.Columns().To, query.To)
	}
	if query.Contract != "" {
		where = where.Where(dao.ChainData.Columns().Contract, query.Contract)
	}
	///time
	if query.StartTime != 0 {
		where = where.WhereGTE(dao.ChainData.Columns().Ts, query.StartTime)
	}
	if query.EndTime != 0 {
		where = where.WhereLTE(dao.ChainData.Columns().Ts, query.EndTime)
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
	data := []*entity.ChainData{}
	err = result.Structs(&data)
	///
	return data, err
}
