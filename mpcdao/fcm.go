package mpcdao

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type Fcm struct {
	redis *gredis.Redis
	dur   time.Duration
}
type QueryFcmToken struct {
	FcmToken string `json:"fcm_token"`
	Token    string `json:"token"`
	Address  string `json:"address"`
}

func NewFcm(redis *gredis.Redis, dur int) *Fcm {
	// g.DB(dao.ChainTransfer.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	if redis != nil {
		dao.FcmToken.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	}
	return &Fcm{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}

func (s *Fcm) ExistsFcmToken(ctx context.Context, address string, fcmToken string) (bool, error) {
	result, err := dao.FcmToken.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.FcmToken.Table() + address,
		Force:    true,
	}).Where(dao.FcmToken.Columns().Address, address).
		Where(dao.FcmToken.Columns().FcmToken, fcmToken).Count()
	if err != nil {
		return false, err
	}

	if result == 0 {
		return false, nil
	}
	return true, nil
}

// /////
func (s *Fcm) InsertFcmToken(ctx context.Context, data *entity.FcmToken) error {
	_, err := dao.FcmToken.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.FcmToken.Table() + data.Address,
		Force:    true,
	}).Insert(data)
	return err
}
func (s *Fcm) InsertPushErr(ctx context.Context, data *entity.PushErr) error {
	_, err := dao.PushErr.Ctx(ctx).Insert(data)
	return err
}
func (s *Fcm) InsertFcmOfflineMsg(ctx context.Context, data *entity.FcmOfflineMsg) error {
	_, err := dao.FcmOfflineMsg.Ctx(ctx).Insert(data)
	return err
}
func (s *Fcm) DeleteOfflineMsgs(ctx context.Context, ids []string) error {
	_, err := dao.FcmOfflineMsg.Ctx(ctx).Where(dao.FcmOfflineMsg.Columns().Id, ids).Delete()
	return err
}

type PosFcmToken *entity.FcmToken

func (s *Fcm) QueryFcmTokenAll(ctx context.Context, pos PosFcmToken, limit int) ([]*entity.FcmToken, PosFcmToken, error) {
	where := dao.FcmToken.Ctx(ctx)
	if pos != nil {
		where = where.WhereGT(dao.FcmToken.Columns().Address, pos.Address)
		where = where.WhereGT(dao.FcmToken.Columns().FcmToken, pos.FcmToken)
		where = where.OrderAsc(dao.FcmToken.Columns().Address)
		where = where.OrderAsc(dao.FcmToken.Columns().FcmToken)
	}
	///
	if limit > 0 {
		where = where.Limit(limit)
	}
	///
	where.Fields()
	result, err := where.All()
	if result.Len() == 0 {
		return nil, nil, nil
	}
	if err != nil {
		return nil, nil, err
	}
	data := []*entity.FcmToken{}
	err = result.Structs(&data)
	///
	return data, data[len(data)-1], err
}

type PosFcmOffline *entity.FcmOfflineMsg

func (s *Fcm) QueryFcmOfflineMsg(ctx context.Context, pos PosFcmOffline, limit int) ([]*entity.FcmOfflineMsg, PosFcmOffline, error) {
	where := dao.FcmOfflineMsg.Ctx(ctx)
	if pos != nil {
		where = where.WhereGTE(dao.FcmToken.Columns().Address, pos.Address)
	}
	///
	if limit > 0 {
		where = where.Limit(limit)
	}
	///
	result, err := where.All()
	if err != nil {
		return nil, nil, err
	}
	data := []*entity.FcmOfflineMsg{}
	err = result.Structs(&data)
	///
	return data, data[len(data)-1], err
}

func (s *Fcm) QueryFcmToken(ctx context.Context, query *QueryFcmToken) (*entity.FcmToken, error) {
	if query == nil {
		return nil, nil
	}
	if query.Address == "" {
		return nil, nil
	}
	//
	where := dao.FcmToken.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.FcmToken.Table() + query.Address,
		Force:    true,
	})
	where = where.Where(dao.FcmToken.Columns().Address, query.Address)
	if query.FcmToken != "" {
		where = where.Where(dao.FcmToken.Columns().FcmToken, query.FcmToken)
	}
	///

	///
	rst, err := where.One()
	if err != nil {
		return nil, err
	}
	data := &entity.FcmToken{}
	err = rst.Struct(&data)
	///
	return data, err
}

func (s *Fcm) DelFcmToken(ctx context.Context, address string, fcmToken string) error {
	//
	where := dao.FcmToken.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.FcmToken.Table() + address,
		Force:    true,
	})
	where = where.Where(dao.FcmToken.Columns().Address, address)
	where = where.Where(dao.FcmToken.Columns().FcmToken, fcmToken)
	///

	///
	_, err := where.Delete()
	return err
}
func (s *Fcm) UpdateFcmToken(ctx context.Context, address string, data *entity.FcmToken) error {
	//
	where := dao.FcmToken.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.FcmToken.Table() + address,
		Force:    true,
	})
	_, err := where.Data(data).Where(dao.FcmToken.Columns().Address, address).Update()
	// OnConflict(dao.Chaincfg.Columns().ChainId).
	// Save()
	return err
}
