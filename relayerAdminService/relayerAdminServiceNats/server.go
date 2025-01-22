package relayerAdminServiceNats

import (
	"context"

	"time"

	"github.com/mpcsdk/mpcCommon/mpcdao"
	"github.com/nats-io/nats.go"
)

type RelayerAdminNatsService struct {
	ctx context.Context
	sub *nats.Subscription
	nc  *nats.Conn
	////
	relayerRepo *RelayerAdminRepo
	////
	opt *RelayerAdminServiceOpt
}

func NewRiskAdminNatsService(ctx context.Context, opt *RelayerAdminServiceOpt) (*RelayerAdminNatsService, error) {
	//
	nc, err := nats.Connect(opt.Url, nats.Timeout(time.Duration(opt.TimeOut)*time.Second))
	if err != nil {
		return nil, err
	}
	s := &RelayerAdminNatsService{
		ctx:         ctx,
		opt:         opt,
		relayerRepo: NewRelayerAdminRepo(opt.redis, opt.cacheDur),
	}
	///
	s.nc = nc
	///
	go s.runConsumeAppCfgFn()
	go s.runConsumeAssignFeeFen()
	go s.runConsumeSpecifiedGasFn()
	return s, nil
}

// //

func (s *RelayerAdminNatsService) RelayerAdminDB() *mpcdao.RelayerAdminDB {
	return s.relayerRepo.RelayerAdminDB()
}
func (s *RelayerAdminNatsService) RelayerRepo() *RelayerAdminRepo {
	return s.relayerRepo
}
