package mpcdao

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
)

type MpcContext struct {
	redis *gredis.Redis
	dur   time.Duration
}

func (s *MpcContext) ExistsMpcAddr(ctx context.Context, addr string) (bool, error) {
	cnt, err := dao.MpcContext.DB().Model(dao.MpcContext.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.MpcContext.Table() + addr,
		Force:    true,
	}).Where(dao.MpcContext.Columns().PubKey, addr).Count()
	if err != nil {
		g.Log().Error(ctx, "ExistsMpcAddr:", "addr", addr, "err", err)
		return false, mpccode.CodeInternalError()
	}

	if cnt > 0 {
		return true, nil
	}
	return false, nil
}

func NewMcpContet(redis *gredis.Redis, dur int) *MpcContext {
	// dao.MpcContext.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(redis))

	g.DB(dao.MpcContext.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	return &MpcContext{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}
