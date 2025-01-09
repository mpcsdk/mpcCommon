package mpcdao

import (
	"context"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

// ///
// ////
type RiskAdminDB struct {
	redis *gredis.Redis
	dur   time.Duration
}

func NewRiskAdminDB(redis *gredis.Redis, dur int) *RiskAdminDB {
	if redis != nil {
		g.DB(dao.RiskadminChaincfg.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
		g.DB(dao.RiskadminContractabi.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
		g.DB(dao.RiskadminRiskcontrolRule.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	}

	return &RiskAdminDB{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}

// /contract abi
func (s *RiskAdminDB) AllContractAbi(ctx context.Context) ([]*entity.RiskadminContractabi, error) {
	model := dao.RiskadminContractabi.Ctx(ctx)
	////
	// model = model.Hook(gdb.HookHandler{
	// 	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
	// 		fmt.Println(in)
	// 		return nil, nil
	// 	},
	// })
	///
	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	// ///
	rule := []*entity.RiskadminContractabi{}
	rst.Structs(&rule)
	return rule, nil
}

func (s *RiskAdminDB) GetContractAbiBriefs(ctx context.Context, ChainId int64, kind string) ([]*entity.RiskadminContractabi, error) {
	model := dao.RiskadminContractabi.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.RiskadminContractabi.Table() + strconv.FormatInt(ChainId, 10) + kind + "briefs",
		Force:    true,
	}).FieldsEx(
		dao.RiskadminContractabi.Columns().AbiContent,
	)
	if ChainId > 0 {
		model = model.Where(dao.RiskadminContractabi.Columns().ChainId, ChainId)
	}
	if kind != "" {
		model = model.Where(dao.RiskadminContractabi.Columns().ContractKind, kind)
	}
	////
	// model = model.Hook(gdb.HookHandler{
	// 	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
	// 		fmt.Println(in)
	// 		return nil, nil
	// 	},
	// })
	///
	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	// ///
	rule := []*entity.RiskadminContractabi{}
	rst.Structs(&rule)
	return rule, nil
}

// /
func (s *RiskAdminDB) GetContractAbiById(ctx context.Context, id int) (*entity.RiskadminContractabi, error) {
	where := dao.RiskadminContractabi.Ctx(ctx).
		Where(dao.RiskadminContractabi.Columns().Id, id)

	rst, err := where.One()
	if err != nil {
		return nil, err
	}
	// /
	rule := &entity.RiskadminContractabi{}
	rst.Struct(&rule)
	return rule, nil
}
func (s *RiskAdminDB) GetContractAbi(ctx context.Context, ChainId int64, address string) (*entity.RiskadminContractabi, error) {
	where := dao.RiskadminContractabi.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.RiskadminContractabi.Table() + strconv.FormatInt(ChainId, 10) + address,
		Force:    true,
	}).
		Where(dao.RiskadminContractabi.Columns().ChainId, ChainId).
		Where(dao.RiskadminContractabi.Columns().ContractAddress, address)

	rst, err := where.One()
	if err != nil {
		return nil, err
	}
	// /
	rule := &entity.RiskadminContractabi{}
	rst.Struct(&rule)
	return rule, nil
}

// func (s *RiskAdminDB) ClearContractRuleCache(ctx context.Context,
// 	ChainId int64, kind string, address string) {
// 	cache := g.DB(dao.RiskcontrolRule.Group()).GetCache()
// 	if cache != nil {
// 		g.DB(dao.RiskcontrolRule.Group()).GetCache().Remove(ctx, "SelectCache:"+dao.Contractrule.Table()+strconv.FormatInt(ChainId, 10)+kind+address)
// 	}
// }

// // riskctrl rule
func (s *RiskAdminDB) AllRiskCtrlRule(ctx context.Context) ([]*entity.RiskadminRiskcontrolRule, error) {
	model := dao.RiskadminRiskcontrolRule.Ctx(ctx)
	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	///
	rule := []*entity.RiskadminRiskcontrolRule{}
	rst.Structs(&rule)
	return rule, nil
}

func (s *RiskAdminDB) GetRiskCtrlRuleBriefs(ctx context.Context) ([]*entity.RiskadminRiskcontrolRule, error) {
	model := dao.RiskadminRiskcontrolRule.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.RiskadminRiskcontrolRule.Table(),
		Force:    true,
	}).Fields(
		dao.RiskadminRiskcontrolRule.Columns().Id,
		dao.RiskadminRiskcontrolRule.Columns().RuleName,
		dao.RiskadminRiskcontrolRule.Columns().Salience,
	)

	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	///
	rule := []*entity.RiskadminRiskcontrolRule{}
	rst.Structs(&rule)
	return rule, nil
}

// ///
func (s *RiskAdminDB) GetRiskCtrlRuleById(ctx context.Context, id int) (*entity.RiskadminRiskcontrolRule, error) {
	rst, err := dao.RiskadminRiskcontrolRule.Ctx(ctx).
		Where(dao.RiskadminRiskcontrolRule.Columns().Id, id).One()
	if err != nil {
		return nil, err
	}
	// /
	rule := &entity.RiskadminRiskcontrolRule{}
	rst.Struct(&rule)
	return rule, nil
}

// /
func (s *RiskAdminDB) GetRiskCtrlRule(ctx context.Context, id int, flush bool) (*entity.RiskadminRiskcontrolRule, error) {
	ids := strconv.Itoa(id)
	rst, err := dao.RiskadminRiskcontrolRule.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: func() time.Duration {
			if flush {
				return -1
			} else {
				return s.dur
			}
		}(),
		Name:  dao.RiskadminRiskcontrolRule.Table() + ids,
		Force: true,
	}).Where(dao.RiskadminRiskcontrolRule.Columns().Id, ids).One()
	// Where(dao.RiskcontrolRule.Columns().RuleName, RuleName).One()
	if err != nil {
		return nil, err
	}
	// /
	rule := &entity.RiskadminRiskcontrolRule{}
	rst.Struct(&rule)
	return rule, nil
}

// ////
// ////
// /////chains
func (s *RiskAdminDB) UpdateHeigh(ctx context.Context, chainId int64, heigh int64) error {
	_, err := dao.RiskadminChaincfg.Ctx(ctx).
		Data(g.Map{dao.RiskadminChaincfg.Columns().Heigh: heigh}).
		Where(dao.RiskadminChaincfg.Columns().ChainId, chainId).
		Update()
		// OnConflict(dao.Chaincfg.Columns().ChainId).
		// Save()
	return err
}
func (s *RiskAdminDB) AllChainsCfg(ctx context.Context) ([]*entity.RiskadminChaincfg, error) {
	rst, err := dao.RiskadminChaincfg.Ctx(ctx).All()
	if err != nil {
		return nil, err
	}
	data := make([]*entity.RiskadminChaincfg, 0)
	err = rst.Structs(&data)
	return data, err
}
func (s *RiskAdminDB) GetChainCfg(ctx context.Context, chainId int64) (*entity.RiskadminChaincfg, error) {

	rst, err := dao.RiskadminChaincfg.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.RiskadminChaincfg.Table() + gconv.String(chainId),
		Force:    true,
	}).Where(dao.RiskadminChaincfg.Columns().ChainId, chainId).One()
	if err != nil {
		return nil, err
	}
	////
	data := &entity.RiskadminChaincfg{}
	err = rst.Struct(&data)
	return data, err

}

func (s *RiskAdminDB) GetChainCfgById(ctx context.Context, id int) (*entity.RiskadminChaincfg, error) {
	rst, err := dao.RiskadminChaincfg.Ctx(ctx).Where(dao.RiskadminChaincfg.Columns().Id, id).One()
	if err != nil {
		return nil, err
	}
	////
	data := &entity.RiskadminChaincfg{}
	err = rst.Struct(&data)
	return data, err

}
