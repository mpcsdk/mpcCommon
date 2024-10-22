package authServiceNrpc

import (
	"time"

	"github.com/nats-io/nats.go"
)

type AuthClientNrpc struct {
	authclient *AuthServiceClient
	nc         *nats.Conn
}

func NewAuthRpcClient(natsUrl string, timeout int64) (*AuthClientNrpc, error) {
	s := &AuthClientNrpc{}
	nc, err := nats.Connect(natsUrl, nats.Timeout(time.Second*time.Duration(timeout)))
	if err != nil {
		return nil, err
	}
	authclient := NewAuthServiceClient(nc)

	///
	s.nc = nc
	s.authclient = authclient
	return s, nil
}
func (s *AuthClientNrpc) Flush() {
	err := s.nc.Flush()
	if err != nil {
		panic(err)
	}
	s.authclient = NewAuthServiceClient(s.nc)
}
