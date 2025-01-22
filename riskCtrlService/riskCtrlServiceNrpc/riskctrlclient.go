package riskCtrlServiceNrpc

import (
	"context"

	"github.com/franklihub/nrpc"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpccode"
	riskctrlservicemodel "github.com/mpcsdk/mpcCommon/riskCtrlService/riskCtrlServiceModel"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RiskCtrlRpcClient struct {
	cli   *RiskCtrlServiceClient
	nc    *nats.Conn
	cache *gcache.Cache
}

func NewRiskCtrlRpcClient(r *gredis.Redis, natsUrl string, timeout int) (*RiskCtrlRpcClient, error) {
	s := &RiskCtrlRpcClient{}
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		return nil, err
	}
	cli := NewRiskCtrlServiceClient(nc)
	///
	_, err = r.Conn(gctx.GetInitCtx())
	if err != nil {
		return nil, err
	}
	cache := gcache.New()
	cache.SetAdapter(gcache.NewAdapterRedis(r))

	///
	s.nc = nc
	s.cli = cli
	s.cache = cache
	return s, nil
}
func (s *RiskCtrlRpcClient) Flush() {
	err := s.nc.Flush()
	if err != nil {
		panic(err)
	}
	s.cli = NewRiskCtrlServiceClient(s.nc)
}
func (s *RiskCtrlRpcClient) TryFlush(err error) {
	if _, ok := err.(*nrpc.Error); ok {
		return
	} else {
		if err == nats.ErrTimeout {

		} else {
			return

		}
	}
	err = s.nc.Flush()
	if err != nil {
		panic(err)
	}
	s.cli = NewRiskCtrlServiceClient(s.nc)
}

// ///
// ///
func (s *RiskCtrlRpcClient) Alive(ctx context.Context) error {
	return nil
	_, err := s.cli.Alive(ctx, &emptypb.Empty{})
	if err != nil {
		s.TryFlush(err)
		return mpccode.FromNrcpErr(err)
	}
	return nil
}
func (s *RiskCtrlRpcClient) SendPhoneCode(ctx context.Context, req *riskctrlservicemodel.SendPhoneCodeReq) error {
	_, err := s.cli.SendPhoneCode(ctx, &SendPhoneCodeReq{
		RiskSerial: req.RiskSerial,
		UserId:     req.UserId,
		Phone:      req.Phone,
	})

	if err != nil {
		s.TryFlush(err)
		return mpccode.FromNrcpErr(err)
	}
	return err
}
func (s *RiskCtrlRpcClient) SendMailCode(ctx context.Context, req *riskctrlservicemodel.SendMailCodeReq) error {
	_, err := s.cli.SendMailCode(ctx, &SendMailCodeReq{
		RiskSerial: req.RiskSerial,
		UserId:     req.UserId,
		Mail:       req.Mail,
	})

	if err != nil {
		s.TryFlush(err)
		return mpccode.FromNrcpErr(err)
	}
	return err
}
func (s *RiskCtrlRpcClient) VerifyCode(ctx context.Context, req *riskctrlservicemodel.VerifyCodeReq) error {
	_, err := s.cli.VerifyCode(ctx, &VerifyCodeReq{
		RiskSerial: req.RiskSerial,
		UserId:     req.UserId,
		PhoneCode:  req.PhoneCode,
		MailCode:   req.MailCode,
	})

	if err != nil {
		s.TryFlush(err)
		return mpccode.FromNrcpErr(err)
	}
	return err
}
func (s *RiskCtrlRpcClient) RiskTxs(ctx context.Context, req *riskctrlservicemodel.RiskTxsReq) (*riskctrlservicemodel.RiskTxsRes, error) {
	rst, err := s.cli.TxsRequest(ctx, &TxRequestReq{
		UserId:     req.UserId,
		SignTxData: req.SignData,
	})

	if err != nil {
		s.TryFlush(err)
		return nil, mpccode.FromNrcpErr(err)
	}
	return &riskctrlservicemodel.RiskTxsRes{
		Ok:       rst.Ok,
		RiskKind: rst.RiskKind,
		Msg:      rst.Msg,
	}, nil
}
