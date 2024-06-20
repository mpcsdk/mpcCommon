package mpcdao

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type EnhancedRiskCtrl struct {
	redis *gredis.Redis
	dur   time.Duration
}
type QueryEnhancedRiskCtrlRes struct {
	ChainId  int64
	From     string
	Contract string
	TokenId  string
	StartTs  int64
	EndTs    int64
}

func (s *EnhancedRiskCtrl) Clear(ctx context.Context, res QueryEnhancedRiskCtrlRes) error {
	if res.EndTs <= 0 {
		return nil
	}
	key := aggKey(res.ChainId, res.From, res.Contract, res.TokenId)
	_, err := s.redis.Do(ctx, "ZREMRANGEBYSCORE", key, 0, res.EndTs)
	return err
}

func (s *EnhancedRiskCtrl) GetAgg(ctx context.Context, res QueryEnhancedRiskCtrlRes) ([]*entity.ChainTx, error) {
	if res.EndTs == 0 {
		res.EndTs = math.MaxInt64
	}
	key := aggKey(res.ChainId, res.From, res.Contract, res.TokenId)
	v, err := s.redis.Do(ctx, "ZRANGEBYSCORE", key, res.StartTs, res.EndTs)
	if err != nil {
		return nil, err
	}
	//
	data := []*entity.ChainTx{}
	v.Struct(&data)
	///

	return data, nil
}

func (s *EnhancedRiskCtrl) GetAggSum(ctx context.Context, res QueryEnhancedRiskCtrlRes) (*big.Int, error) {
	if res.EndTs == 0 {
		res.EndTs = math.MaxInt64
	}
	key := aggKey(res.ChainId, res.From, res.Contract, res.TokenId)
	v, err := s.redis.Do(ctx, "ZRANGEBYSCORE", key, res.StartTs, res.EndTs)
	if err != nil {
		return nil, err
	}
	//
	data := []*entity.ChainTx{}
	v.Struct(&data)
	///
	sum := big.NewInt(0)

	for _, tx := range data {
		i := big.NewInt(0)
		i.SetString(tx.Value, 10)
		sum = sum.Add(sum, i)
	}
	return sum, nil
}
func (s *EnhancedRiskCtrl) GetAggCnt(ctx context.Context, res QueryEnhancedRiskCtrlRes) (int64, error) {
	if res.EndTs == 0 {
		res.EndTs = math.MaxInt64
	}
	key := aggKey(res.ChainId, res.From, res.Contract, res.TokenId)
	v, err := s.redis.Do(ctx, "Zcount", key, res.StartTs, res.EndTs)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}
func aggKey(chainId int64, from, contract string, tokenId string) string {
	return fmt.Sprintf("%d_%s_%s_%s", chainId, from, contract, tokenId)
}
func (s *EnhancedRiskCtrl) AggTx(ctx context.Context, tx *entity.ChainTx) error {
	_, err := s.redis.Do(ctx, "Zadd", aggKey(tx.ChainId, tx.From, tx.Contract, tx.TokenId), tx.Ts, tx)
	if err != nil {
		return err
	}
	///
	return nil
}

// //
// /
func (s *EnhancedRiskCtrl) InsertTx(ctx context.Context, tx *entity.ChainTx) error {
	_, err := dao.ChainTx.Ctx(ctx).Insert(tx)
	return err
}

// //
type QueryTx struct {
	ChainId  int64  `json:"chainId"`
	From     string `json:"from"`
	To       string `json:"to"`
	Contract string `json:"contract"`
	Kind     string `json:"kind"`
	///
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	///
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func (s *EnhancedRiskCtrl) Query(ctx context.Context, query *QueryTx) ([]*entity.ChainTx, error) {
	if query.PageSize < 0 || query.Page < 0 {
		return nil, nil
	}
	//
	where := dao.ChainTx.Ctx(ctx)
	if query.ChainId != 0 {
		where = where.Where(dao.ChainTx.Columns().ChainId, query.ChainId)
	}
	if query.Kind != "" {
		where = where.Where(dao.ChainTx.Columns().Kind, query.Kind)
	}
	if query.From != "" {
		where = where.Where(dao.ChainTx.Columns().From, query.From)
	}
	if query.To != "" {
		where = where.Where(dao.ChainTx.Columns().To, query.To)
	}
	if query.Contract != "" {
		where = where.Where(dao.ChainTx.Columns().Contract, query.Contract)
	}
	///time
	if query.StartTime != 0 {
		where = where.WhereGTE(dao.ChainTx.Columns().Ts, query.StartTime)
	}
	if query.EndTime != 0 {
		where = where.WhereLTE(dao.ChainTx.Columns().Ts, query.EndTime)
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
	data := []*entity.ChainTx{}
	err = result.Structs(&data)
	///
	return data, err
}

// //
// /
func NewEnhancedRiskCtrl(redis *gredis.Redis, dur int) *EnhancedRiskCtrl {
	g.DB(dao.ChainTx.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))

	return &EnhancedRiskCtrl{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}
