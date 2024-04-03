package mpcdao

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/do"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type RiskTfa struct {
	redis *gredis.Redis
	dur   time.Duration
}

func NewRiskTfa(redis *gredis.Redis, dur int) *RiskTfa {
	g.DB(dao.RiskcontrolRule.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))

	return &RiskTfa{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}
func (s *RiskTfa) FetchTfaInfo(ctx context.Context, userId string) (*entity.Tfa, error) {
	if userId == "" {
		return nil, mpccode.CodeParamInvalid()
	}

	aggdo := &do.Tfa{
		UserId: userId,
	}
	var data *entity.Tfa
	///
	rst, err := dao.Tfa.Ctx(ctx).Where(aggdo).One()
	if err != nil {
		g.Log().Error(ctx, "ExistsTfaInfo:", "userId", userId, "agg:", aggdo, "err", err)
		return nil, mpccode.CodeInternalError()
	}
	if rst.IsEmpty() {
		return nil, nil
	}
	err = rst.Struct(&data)
	if err != nil {
		return nil, mpccode.CodeInternalError()
	}

	return data, nil
}
