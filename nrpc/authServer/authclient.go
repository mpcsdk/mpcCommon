package authServer

import (
	"context"
	"time"

	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/nats-io/nats.go"
)

type NRpcClient struct {
	authclient *AuthServerClient
	nc         *nats.Conn
}

func NewAuthRpcClient(natsUrl string, timeout int64) (*NRpcClient, error) {
	s := &NRpcClient{}
	nc, err := nats.Connect(natsUrl, nats.Timeout(time.Second*time.Duration(timeout)))
	if err != nil {
		return nil, err
	}
	authclient := NewAuthServerClient(nc)

	///
	s.nc = nc
	s.authclient = authclient
	return s, nil
}
func (s *NRpcClient) Flush() {
	err := s.nc.Flush()
	if err != nil {
		panic(err)
	}
	s.authclient = NewAuthServerClient(s.nc)
}
func (s *NRpcClient) AuthToken(ctx context.Context, req *AuthTokenReq) (*AuthTokenRes, error) {
	res, err := s.authclient.AuthToken(req)
	if err != nil {
		if err == nats.ErrTimeout {
			s.Flush()
			return nil, mpccode.CodeTimeOut()
		}
		return nil, err
	}
	return res, nil
}
func (s *NRpcClient) RefreshToken(ctx context.Context, req *RefreshTokenReq) (*RefreshTokenRes, error) {
	res, err := s.authclient.RefreshToken(req)
	if err != nil {
		if err == nats.ErrTimeout {
			s.Flush()
			return nil, mpccode.CodeTimeOut()
		}
		return nil, err
	}
	return res, nil
}
func (s *NRpcClient) TokenGetInfo(ctx context.Context, req *TokenGetInfoReq) (*TokenGetInfoRes, error) {
	res, err := s.authclient.TokenGetInfo(req)
	if err != nil {
		if err == nats.ErrTimeout {
			s.Flush()
			return nil, mpccode.CodeTimeOut()
		}
		return nil, err
	}
	return res, nil
}
