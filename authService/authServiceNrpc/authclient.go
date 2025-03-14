package authServiceNrpc

import (
	"context"
	"time"

	"github.com/franklihub/nrpc"
	"github.com/gogf/gf/v2/database/gredis"
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

func NewAuthRpcClient(r *gredis.Redis, natsUrl string, timeout int) (*AuthRpcClient, error) {
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

func (s *AuthRpcClient) TryFlush(err error) {
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
	s.authclient = NewAuthServiceClient(s.nc)
}

func (s *AuthRpcClient) AuthToken(ctx context.Context, tokenStr string) (string, error) {
	// if v, err := s.cache.Get(ctx, "AuthNrpc:AuthToken:"+tokenStr); err == nil && !v.IsEmpty() {
	// 	return v.String(), nil
	// }
	res, err := s.authclient.AuthToken(ctx, &AuthTokenReq{UserToken: tokenStr})
	if err != nil {
		s.TryFlush(err)
		return "", mpccode.FromNrcpErr(err)
	}
	//todo: expire
	// s.cache.Set(ctx, "AuthNrpc:AuthToken:"+tokenStr, res, s.cacheDur)

	return res.Token, err
}
func (s *AuthRpcClient) RefreshToken(ctx context.Context, tokenStr string) (string, error) {
	// if v, err := s.cache.Get(ctx, "AuthNrpc:RefreshToken:"+tokenStr); err == nil && !v.IsEmpty() {
	// 	return v.String(), nil
	// }
	res, err := s.authclient.RefreshToken(ctx, &RefreshTokenReq{Token: tokenStr})
	if err != nil {
		s.TryFlush(err)
		return "", mpccode.FromNrcpErr(err)
	}
	//todo: expire
	// s.cache.Set(ctx, "AuthNrpc:RefreshToken:"+tokenStr, res, s.cacheDur)

	return res.Token, err
}
func (s *AuthRpcClient) TokenInfo(ctx context.Context, tokenStr string) (*authServiceModel.MpcUserToken, error) {
	// if v, err := s.cache.Get(ctx, "AuthNrpc:TokenInfo:"+tokenStr); err == nil && !v.IsEmpty() {
	// 	var res *authServiceModel.MpcUserToken = nil
	// 	v.Struct(&res)
	// 	return res, nil
	// }
	res, err := s.authclient.TokenInfo(ctx, &TokenInfoReq{Token: tokenStr})

	if err != nil {
		s.TryFlush(err)
		return nil, mpccode.FromNrcpErr(err)
	}
	tokenInfo := &authServiceModel.MpcUserToken{
		UserInfo: authServiceModel.UserInfo{
			AppId:  res.AppId,
			UserId: res.UserId,
		},
	}
	//todo: expire
	// s.cache.Set(ctx, "AuthNrpc:TokenInfo:"+tokenStr, tokenInfo, s.cacheDur)

	return tokenInfo, err
}

// ///
// ///
func (s *AuthRpcClient) Alive(ctx context.Context) error {
	return nil
	_, err := s.authclient.Alive(ctx, &emptypb.Empty{})
	if err != nil {
		s.TryFlush(err)
		return mpccode.FromNrcpErr(err)
	}
	return err
}
