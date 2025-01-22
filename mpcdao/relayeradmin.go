package mpcdao

import (
	"context"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

// ///
// ////
type RelayerAdminDB struct {
	redis *gredis.Redis
	dur   time.Duration
}

func NewRelayerAdminDB(redis *gredis.Redis, dur int) *RelayerAdminDB {
	if redis != nil {
		g.DB(dao.RelayeradminAppCfg.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
		g.DB(dao.RelayeradminAssignFee.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
		g.DB(dao.RelayeradminSpecifiedGas.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	}

	return &RelayerAdminDB{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}

// // assignfee
func (s *RelayerAdminDB) AllAssignFee(ctx context.Context) ([]*entity.RelayeradminAssignFee, error) {
	model := dao.RelayeradminAssignFee.Ctx(ctx)
	///
	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	// ///
	data := []*entity.RelayeradminAssignFee{}
	rst.Structs(&data)
	return data, nil
}

func (s *RelayerAdminDB) GetAssignFee(ctx context.Context, appId string, chainId int64) ([]*entity.RelayeradminAssignFee, error) {
	model := dao.RelayeradminAssignFee.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.RelayeradminAssignFee.Table() + strconv.FormatInt(chainId, 10) + appId,
		Force:    true,
	})
	if chainId > 0 {
		model = model.Where(dao.RelayeradminAssignFee.Columns().ChainId, chainId)
	}

	///
	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	// ///
	data := []*entity.RelayeradminAssignFee{}
	rst.Structs(&data)
	return data, nil
}

// /
func (s *RelayerAdminDB) GetAssignFeeById(ctx context.Context, id int) (*entity.RelayeradminAssignFee, error) {
	where := dao.RelayeradminAssignFee.Ctx(ctx).
		Where(dao.RelayeradminAssignFee.Columns().Id, id)

	rst, err := where.One()
	if err != nil {
		return nil, err
	}
	// /
	data := &entity.RelayeradminAssignFee{}
	rst.Struct(&data)
	return data, nil
}

// // specifiedgas
func (s *RelayerAdminDB) AllSpecifiedGas(ctx context.Context) ([]*entity.RelayeradminSpecifiedGas, error) {
	model := dao.RelayeradminSpecifiedGas.Ctx(ctx)
	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	///
	rule := []*entity.RelayeradminSpecifiedGas{}
	rst.Structs(&rule)
	return rule, nil
}

func (s *RelayerAdminDB) GetSpecifiedGas(ctx context.Context, chainId int64, contractAddr string, methodSig string) ([]*entity.RelayeradminSpecifiedGas, error) {
	model := dao.RelayeradminSpecifiedGas.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.RelayeradminSpecifiedGas.Table(),
		Force:    true,
	}).Where(
		dao.RelayeradminSpecifiedGas.Columns().ChainId, chainId,
		dao.RelayeradminSpecifiedGas.Columns().ContractAddr, contractAddr,
		dao.RelayeradminSpecifiedGas.Columns().MethodSig, methodSig,
	)

	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	///
	rule := []*entity.RelayeradminSpecifiedGas{}
	rst.Structs(&rule)
	return rule, nil
}

// ///
func (s *RelayerAdminDB) GetSpecifiedGasById(ctx context.Context, id int) (*entity.RelayeradminSpecifiedGas, error) {
	rst, err := dao.RelayeradminSpecifiedGas.Ctx(ctx).
		Where(dao.RelayeradminSpecifiedGas.Columns().Id, id).One()
	if err != nil {
		return nil, err
	}
	// /
	rule := &entity.RelayeradminSpecifiedGas{}
	rst.Struct(&rule)
	return rule, nil
}

// /

// ////
// ////
// /////apps
func (s *RelayerAdminDB) AllAppCfg(ctx context.Context) ([]*entity.RelayeradminAppCfg, error) {
	rst, err := dao.RelayeradminAppCfg.Ctx(ctx).All()
	if err != nil {
		return nil, err
	}
	data := make([]*entity.RelayeradminAppCfg, 0)
	err = rst.Structs(&data)
	return data, err
}
func (s *RelayerAdminDB) GetAppCfg(ctx context.Context, appId string) (*entity.RelayeradminAppCfg, error) {

	rst, err := dao.RelayeradminAppCfg.Ctx(ctx).Where(dao.RelayeradminAppCfg.Columns().AppId, appId).One()
	if err != nil {
		return nil, err
	}
	////
	data := &entity.RelayeradminAppCfg{}
	err = rst.Struct(&data)
	return data, err

}

func (s *RelayerAdminDB) GetAppCfgById(ctx context.Context, id int) (*entity.RelayeradminAppCfg, error) {
	rst, err := dao.RelayeradminAppCfg.Ctx(ctx).Where(dao.RelayeradminAppCfg.Columns().Id, id).One()
	if err != nil {
		return nil, err
	}
	////
	data := &entity.RelayeradminAppCfg{}
	err = rst.Struct(&data)
	return data, err

}
