package mpcdao

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type EnhancedRiskCtrl struct {
	redis *gredis.Redis
}
type QueryEnhancedRiskCtrlRes struct {
	ChainId  int64
	From     string
	Contract string
	StartTs  int64
	EndTs    int64
}

func (s *EnhancedRiskCtrl) AggSum(ctx context.Context, res QueryEnhancedRiskCtrlRes) (*big.Int, error) {
	if res.EndTs == 0 {
		res.EndTs = math.MaxInt64
	}
	key := aggKey(res.ChainId, res.From, res.Contract)
	v, err := s.redis.Do(ctx, "ZRANGE", key, res.StartTs, res.EndTs)
	if err != nil {
		return nil, err
	}
	//
	data := []*entity.ChainData{}
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
	key := aggKey(res.ChainId, res.From, res.Contract)
	v, err := s.redis.Do(ctx, "ZCARD", key, res.StartTs, res.EndTs)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}
func aggKey(chainId int64, from, contract string) string {
	return fmt.Sprintf("%d_%s_%s", chainId, from, contract)
}
func (s *EnhancedRiskCtrl) AggTx(ctx context.Context, tx *entity.ChainTx) error {
	_, err := s.redis.Do(ctx, "Zadd", aggKey(tx.ChainId, tx.From, tx.Contract), tx.Ts, tx)
	if err != nil {
		return err
	}
	return nil
}
func InsertTx(ctx context.Context, tx *entity.ChainTx) error {
	_, err := dao.ChainTx.Ctx(ctx).Insert(tx)
	return err
}

func NewEnhancedRiskCtrl(redis *gredis.Redis) *EnhancedRiskCtrl {

	return &EnhancedRiskCtrl{
		redis: redis,
	}
}
