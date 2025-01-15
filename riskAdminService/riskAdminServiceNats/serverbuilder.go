package riskAdminServiceNats

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/mpcsdk/mpcCommon/mq"
)

type RiskAdminServiceOpt struct {
	Url      string
	TimeOut  int64
	redis    *gredis.Redis
	cacheDur int

	consumeChainFn             func(ctx context.Context, data *mq.RiskAdminChainMsg) error
	consumeContractFn          func(ctx context.Context, data *mq.RiskAdminContractMsg) error
	consumeRiskRuleFn          func(ctx context.Context, data *mq.RiskAdminRiskRuleMsg) error
	consumeRiskRuleCheckRespFn func(ctx context.Context, data *mq.RiskAdminRiskRuleCheckMsg) (*mq.RiskAdminRiskRuleCheckResp, error)
}

func RiskAdminServiceCfgCfgBuilder() *RiskAdminServiceOpt {
	return &RiskAdminServiceOpt{}
}
func (s *RiskAdminServiceOpt) check() error {
	if s.consumeChainFn == nil {
		return errors.New("RelayerRpcServiceCfg consumeRelayerChannelFn is nil")
	}
	if s.consumeContractFn == nil {
		return errors.New("RelayerRpcServiceCfg consumeRelayerFeeFn is nil")
	}
	if s.consumeRiskRuleFn == nil {
		return errors.New("RelayerRpcServiceCfg consumeRelayerFeeFn is nil")
	}
	if s.consumeRiskRuleCheckRespFn == nil {
		return errors.New("RelayerRpcServiceCfg consumeRelayerFeeFn is nil")
	}
	return nil
}
func (s *RiskAdminServiceOpt) WithUrlTimeOut(url string, timeOut int64) *RiskAdminServiceOpt {
	s.Url = url
	s.TimeOut = timeOut
	return s
}
func (s *RiskAdminServiceOpt) WithConsumeChainFn(consumeChainFn func(ctx context.Context, data *mq.RiskAdminChainMsg) error) *RiskAdminServiceOpt {
	s.consumeChainFn = consumeChainFn
	return s
}
func (s *RiskAdminServiceOpt) WithConsumeContractFn(consumeContractFn func(ctx context.Context, data *mq.RiskAdminContractMsg) error) *RiskAdminServiceOpt {
	s.consumeContractFn = consumeContractFn
	return s
}
func (s *RiskAdminServiceOpt) WithConsumeRiskRuleFn(consumeRiskRuleFn func(ctx context.Context, data *mq.RiskAdminRiskRuleMsg) error) *RiskAdminServiceOpt {
	s.consumeRiskRuleFn = consumeRiskRuleFn
	return s
}
func (s *RiskAdminServiceOpt) WithConsumeRiskRuleCheckRespFn(consumeRiskRuleCheckRespFn func(ctx context.Context, data *mq.RiskAdminRiskRuleCheckMsg) (*mq.RiskAdminRiskRuleCheckResp, error)) *RiskAdminServiceOpt {
	s.consumeRiskRuleCheckRespFn = consumeRiskRuleCheckRespFn
	return s
}
func (s *RiskAdminServiceOpt) WithRedis(redis *gredis.Redis, dur int) *RiskAdminServiceOpt {
	s.redis = redis
	s.cacheDur = dur
	return s
}

// /////
func (s *RiskAdminServiceOpt) Redis() *gredis.Redis {
	return s.redis
}
func (s *RiskAdminServiceOpt) ConsumeChainFn() func(ctx context.Context, data *mq.RiskAdminChainMsg) error {
	return s.consumeChainFn
}
func (s *RiskAdminServiceOpt) ConsumeContractFn() func(ctx context.Context, data *mq.RiskAdminContractMsg) error {
	return s.consumeContractFn
}
func (s *RiskAdminServiceOpt) ConsumeRiskRuleFn() func(ctx context.Context, data *mq.RiskAdminRiskRuleMsg) error {
	return s.consumeRiskRuleFn
}
func (s *RiskAdminServiceOpt) ConsumeRiskRuleCheckRespFn() func(ctx context.Context, data *mq.RiskAdminRiskRuleCheckMsg) (*mq.RiskAdminRiskRuleCheckResp, error) {
	return s.consumeRiskRuleCheckRespFn
}
