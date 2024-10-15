package authServer

import (
	"time"

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
