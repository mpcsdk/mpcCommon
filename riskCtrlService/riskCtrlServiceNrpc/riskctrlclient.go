package riskCtrlServiceNrpc

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
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
	nc, err := nats.Connect(natsUrl, nats.Timeout(time.Second*time.Duration(timeout)))
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

// ///
var errDeadLine = errors.New("nats: timeout")

// ///
func (s *RiskCtrlRpcClient) Alive(ctx context.Context) error {
	_, err := s.cli.Alive(&emptypb.Empty{})
	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RiskCtrlRpcClient Alive TimeOut:")
			s.Flush()
			return nil
		}
	}
	return err
}
func (s *RiskCtrlRpcClient) SendPhoneCode(ctx context.Context, req *riskctrlservicemodel.SendPhoneCodeReq) error {
	_, err := s.cli.SendPhoneCode(&SendPhoneCodeReq{
		RiskSerial: req.RiskSerial,
		UserId:     req.UserId,
		Phone:      req.Phone,
	})

	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RiskCtrlRpcClient SendPhoneCode TimeOut:")
			s.Flush()
			return nil
		}
	}
	return err
}
func (s *RiskCtrlRpcClient) SendMailCode(ctx context.Context, req *riskctrlservicemodel.SendMailCodeReq) error {
	_, err := s.cli.SendMailCode(&SendMailCodeReq{
		RiskSerial: req.RiskSerial,
		UserId:     req.UserId,
		Mail:       req.Mail,
	})

	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RiskCtrlRpcClient SendMailCode TimeOut:")
			s.Flush()
			return nil
		}
	}
	return err
}
func (s *RiskCtrlRpcClient) VerifyCode(ctx context.Context, req *riskctrlservicemodel.VerifyCodeReq) error {
	_, err := s.cli.VerifyCode(&VerifyCodeReq{
		RiskSerial: req.RiskSerial,
		UserId:     req.UserId,
		PhoneCode:  req.PhoneCode,
		MailCode:   req.MailCode,
	})

	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RiskCtrlRpcClient VerifyCode TimeOut:")
			s.Flush()
			return nil
		}
	}
	return err
}
func (s *RiskCtrlRpcClient) RiskTxs(ctx context.Context, req *riskctrlservicemodel.RiskTxsReq) (*riskctrlservicemodel.RiskTxsRes, error) {
	rst, err := s.cli.TxsRequest(&TxRequestReq{
		UserId:     req.UserId,
		SignTxData: req.SignData,
	})

	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "RiskCtrlRpcClient VerifyCode TimeOut:")
			s.Flush()
			return nil, err
		}
	}
	return &riskctrlservicemodel.RiskTxsRes{
		Ok:       rst.Ok,
		RiskKind: rst.RiskKind,
		Msg:      rst.Msg,
	}, nil
}
