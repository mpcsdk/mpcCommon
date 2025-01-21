package relayerAdminServiceNats

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/mpcsdk/mpcCommon/mq"
)

type RelayerAdminServiceOpt struct {
	Url      string
	TimeOut  int
	redis    *gredis.Redis
	cacheDur int

	consumeAppCfgFn       func(ctx context.Context, data *mq.RelayerAdminAppCfgMsg) error
	consumeAssignFeeFen   func(ctx context.Context, data *mq.RelayerAdminAssignFeeMsg) error
	consumeSpecifiedGasFn func(ctx context.Context, data *mq.RelayerAdminSpecifiedGas) error
}

func NewRelayerAdminServiceBuilder() *RelayerAdminServiceOpt {
	return &RelayerAdminServiceOpt{}
}
func (s *RelayerAdminServiceOpt) check() error {
	if s.consumeAppCfgFn == nil {
		return errors.New("RelayerRpcServiceCfg consumeRelayerChannelFn is nil")
	}
	if s.consumeAssignFeeFen == nil {
		return errors.New("RelayerRpcServiceCfg consumeAssignFeeFen is nil")
	}
	if s.consumeSpecifiedGasFn == nil {
		return errors.New("RelayerRpcServiceCfg consumeSpecifiedGasFn is nil")
	}
	return nil
}
func (s *RelayerAdminServiceOpt) WithUrlTimeOut(url string, timeOut int) *RelayerAdminServiceOpt {
	s.Url = url
	s.TimeOut = timeOut
	return s
}
func (s *RelayerAdminServiceOpt) WithConsumeAppCfgFn(consumeAppCfgFn func(ctx context.Context, data *mq.RelayerAdminAppCfgMsg) error) *RelayerAdminServiceOpt {
	s.consumeAppCfgFn = consumeAppCfgFn
	return s
}
func (s *RelayerAdminServiceOpt) WithConsumeAssignFeeFen(consumeAssignFeeFen func(ctx context.Context, data *mq.RelayerAdminAssignFeeMsg) error) *RelayerAdminServiceOpt {
	s.consumeAssignFeeFen = consumeAssignFeeFen
	return s
}
func (s *RelayerAdminServiceOpt) WithConsumeSpecifiedGasFn(consumeSpecifiedGasFn func(ctx context.Context, data *mq.RelayerAdminSpecifiedGas) error) *RelayerAdminServiceOpt {
	s.consumeSpecifiedGasFn = consumeSpecifiedGasFn
	return s
}

func (s *RelayerAdminServiceOpt) WithRedis(redis *gredis.Redis, dur int) *RelayerAdminServiceOpt {
	s.redis = redis
	s.cacheDur = dur
	return s
}

// /////
func (s *RelayerAdminServiceOpt) Redis() *gredis.Redis {
	return s.redis
}

// func (s *RelayerAdminServiceOpt) ConsumeAppCfgFn() func(ctx context.Context, data *mq.RiskAdminChainMsg) error {
// 	return s.consumeAppCfgFn
// }
// func (s *RelayerAdminServiceOpt) ConsumeContractFn() func(ctx context.Context, data *mq.RiskAdminContractMsg) error {
// 	return s.consumeContractFn
// }
// func (s *RelayerAdminServiceOpt) ConsumeRiskRuleFn() func(ctx context.Context, data *mq.RiskAdminRiskRuleMsg) error {
// 	return s.consumeRiskRuleFn
// }
// func (s *RelayerAdminServiceOpt) ConsumeRiskRuleCheckRespFn() func(ctx context.Context, data *mq.RiskAdminRiskRuleCheckMsg) (*mq.RiskAdminRiskRuleCheckResp, error) {
// 	return s.consumeRiskRuleCheckRespFn
// }
