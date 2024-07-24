package mpcdao

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/do"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type MpcContext struct {
	redis *gredis.Redis
	dur   time.Duration
}

// ///context
func (s *MpcContext) UpdateContext(ctx context.Context, userId string, data *do.MpcContext) error {
	_, err := dao.MpcContext.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.MpcContext.Table() + userId,
		Force:    false,
	}).Data(data).Where(do.MpcContext{
		UserId: userId,
	}).Update()
	return err
}

func (s *MpcContext) InertContext(ctx context.Context, userId string, data *do.MpcContext) error {
	cnt, err := dao.MpcContext.Ctx(ctx).Where(do.MpcContext{
		UserId: userId,
	}).CountColumn(dao.MpcContext.Columns().UserId)
	if err != nil {
		return err
	}
	if cnt != 0 {
		return nil
	}

	_, err = dao.MpcContext.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.MpcContext.Table() + userId,
		Force:    false,
	}).Data(data).
		Insert()

	return err
}
func (s *MpcContext) FetchContext(ctx context.Context, userId string) (*entity.MpcContext, error) {
	g.Log().Debug(ctx, "FetchContext userId:", userId)
	var data *entity.MpcContext
	if userId == "" {
		return nil, nil
	}
	rst, err := dao.MpcContext.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     dao.MpcContext.Table() + userId,
		Force:    false,
		// }).Where("user_id", 1).One()
	}).Where(do.MpcContext{
		UserId: userId,
	}).One()
	if err != nil {
		return nil, mpccode.CodeInternalError()
	}

	err = rst.Struct(&data)
	g.Log().Debug(ctx, "FetchContext data:", data)

	return data, err
}

/////

// ///walletaddr
func (s *MpcContext) InsertWalletAddr(ctx context.Context, userId string, addr string, chainId int64) error {
	_, err := dao.WalletAddr.DB().Model(dao.WalletAddr.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.WalletAddr.Table() + gconv.String(chainId) + addr,
		Force:    false,
	}).Data(&entity.WalletAddr{
		UserId:     userId,
		ChainId:    chainId,
		WalletAddr: addr,
	}).Insert()
	if err != nil {
		g.Log().Error(ctx, "ExistsMpcAddr:", "addr", addr, "err", err)
		return mpccode.CodeInternalError()
	}
	return nil
}
func (s *MpcContext) ExistsWalletAddr(ctx context.Context, addr string, chainId int64) (bool, error) {
	rst, err := dao.WalletAddr.DB().Model(dao.WalletAddr.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.WalletAddr.Table() + gconv.String(chainId) + addr,
		Force:    false,
	}).Where(dao.WalletAddr.Columns().WalletAddr, addr).One()
	if err != nil {
		g.Log().Error(ctx, "ExistsWalletAddr:", "addr", addr, "err", err)
		return false, mpccode.CodeInternalError()
	}

	if rst.IsEmpty() {
		return false, nil
	}
	return true, nil
}

func NewMcpContet(redis *gredis.Redis, dur int) *MpcContext {
	// dao.MpcContext.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	if redis != nil {
		g.DB(dao.MpcContext.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
		g.DB(dao.WalletAddr.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	}
	return &MpcContext{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}
