package riskAdminServiceNats

import (
	"context"

	"time"

	"github.com/nats-io/nats.go"
)

type RiskAdminNatsService struct {
	ctx context.Context
	sub *nats.Subscription
	nc  *nats.Conn
	////
	riskadminCfg *RiskAdminCfg
	////
	opt *RiskAdminServiceOpt
}

func NewRiskAdminNatsService(ctx context.Context, opt *RiskAdminServiceOpt) (*RiskAdminNatsService, error) {
	//
	nc, err := nats.Connect(opt.Url, nats.Timeout(time.Duration(opt.TimeOut)*time.Second))
	if err != nil {
		return nil, err
	}
	s := &RiskAdminNatsService{
		ctx:          ctx,
		opt:          opt,
		riskadminCfg: NewRiskAdminCfg(opt.redis, opt.cacheDur),
	}
	///
	s.nc = nc
	///
	go s.runConsumeChainFn()
	go s.runConsumeContractFn()
	go s.runConsumeRiskCtrlRuleFn()
	// go s.runConsumeRiskCtrlRuleCheckRespFn()
	return s, nil
}

// //

func (s *RiskAdminNatsService) RiskAdminCfg() *RiskAdminCfg {
	return s.riskadminCfg
}