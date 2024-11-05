package authServiceNrpc

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/authService/authServiceModel"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthRpcClient struct {
	authclient *AuthServiceClient
	nc         *nats.Conn
	cache      *gcache.Cache
}

func NewAuthRpcClient(r *gredis.Redis, natsUrl string, timeout int64) (*AuthRpcClient, error) {
	s := &AuthRpcClient{}
	nc, err := nats.Connect(natsUrl, nats.Timeout(time.Second*time.Duration(timeout)))
	if err != nil {
		return nil, err
	}
	authclient := NewAuthServiceClient(nc)
	///
	_, err = r.Conn(gctx.GetInitCtx())
	if err != nil {
		return nil, err
	}
	cache := gcache.New()
	cache.SetAdapter(gcache.NewAdapterRedis(r))

	///
	s.nc = nc
	s.authclient = authclient
	s.cache = cache
	return s, nil
}
func (s *AuthRpcClient) Flush() {
	err := s.nc.Flush()
	if err != nil {
		panic(err)
	}
	s.authclient = NewAuthServiceClient(s.nc)
}
func (s *AuthRpcClient) AuthToken(ctx context.Context, tokenStr string) (string, error) {
	if v, err := s.cache.Get(ctx, "AuthNrpc:AuthToken:"+tokenStr); err == nil && !v.IsEmpty() {
		return v.String(), nil
	}
	res, err := s.authclient.AuthToken(&AuthTokenReq{UserToken: tokenStr})
	//todo: expire
	s.cache.Set(ctx, "AuthNrpc:AuthToken:"+tokenStr, res, 0)

	return res.Token, err
}
func (s *AuthRpcClient) RefreshToken(ctx context.Context, tokenStr string) (string, error) {
	if v, err := s.cache.Get(ctx, "AuthNrpc:RefreshToken:"+tokenStr); err == nil && !v.IsEmpty() {
		return v.String(), nil
	}
	res, err := s.authclient.RefreshToken(&RefreshTokenReq{Token: tokenStr})
	//todo: expire
	s.cache.Set(ctx, "AuthNrpc:RefreshToken:"+tokenStr, res, 0)

	return res.Token, err
}
func (s *AuthRpcClient) TokenInfo(ctx context.Context, tokenStr string) (*authServiceModel.MpcUserToken, error) {
	if v, err := s.cache.Get(ctx, "AuthNrpc:TokenInfo:"+tokenStr); err == nil && !v.IsEmpty() {
		var res *authServiceModel.MpcUserToken = nil
		v.Struct(&res)
		return res, nil
	}
	res, err := s.authclient.TokenInfo(&TokenInfoReq{Token: tokenStr})

	if err != nil {
		return nil, mpccode.FromNrcpErr(err)
	}
	//todo: expire
	s.cache.Set(ctx, "AuthNrpc:TokenInfo:"+tokenStr, res, 0)

	return &authServiceModel.MpcUserToken{}, err
}

// ///
var errDeadLine = errors.New("nats: timeout")

// ///
func (s *AuthRpcClient) Alive(ctx context.Context) error {
	_, err := s.authclient.Alive(&emptypb.Empty{})
	if err != nil {
		if err.Error() == errDeadLine.Error() {
			g.Log().Warning(ctx, "AuthAlive TimeOut:")
			s.Flush()
			return nil
		}
	}
	return err
}
