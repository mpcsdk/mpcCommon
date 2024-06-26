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

type RiskCtrlRule struct {
	redis *gredis.Redis
	dur   time.Duration
}

func NewRiskCtrlRule(redis *gredis.Redis, dur int) *RiskCtrlRule {
	if redis != nil {
		g.DB(dao.RiskcontrolRule.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	}

	return &RiskCtrlRule{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}

// /
func (s *RiskCtrlRule) GetContractAbiBriefs(ctx context.Context, ChainId int64, kind string) ([]*entity.Contractabi, error) {
	// model := g.Model(dao.Contractabi.Table()).Ctx(ctx).Cache(gdb.CacheOption{
	model := dao.Contractabi.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.Contractabi.Table() + strconv.FormatInt(ChainId, 10) + kind,
		Force:    true,
	}).Fields(
		dao.Contractabi.Columns().ChainId,
		dao.Contractabi.Columns().ContractAddress,
		dao.Contractabi.Columns().ContractName,
		dao.Contractabi.Columns().ContractKind,
		dao.Contractabi.Columns().Decimal,
	)
	if ChainId > 0 {
		model = model.Where(dao.Contractabi.Columns().ChainId, ChainId)
	}
	if kind != "" {
		model = model.Where(dao.Contractabi.Columns().ContractKind, kind)
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
	rule := []*entity.Contractabi{}
	rst.Structs(&rule)
	return rule, nil
}

// /

func (s *RiskCtrlRule) GetContractAbi(ctx context.Context, ChainId int64, address string, detail bool) (*entity.Contractabi, error) {
	where := dao.Contractabi.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.Contractabi.Table() + strconv.FormatInt(ChainId, 10) + address + gconv.String(detail),
		Force:    true,
	}).
		Where(dao.Contractabi.Columns().ChainId, ChainId).
		Where(dao.Contractabi.Columns().ContractAddress, address)

	if !detail {
		where = where.FieldsEx(dao.Contractabi.Columns().AbiContent)
	}
	rst, err := where.One()
	if err != nil {
		return nil, err
	}
	// /
	rule := &entity.Contractabi{}
	rst.Struct(&rule)
	return rule, nil
}
func (s *RiskCtrlRule) ClearContractRuleCache(ctx context.Context,
	ChainId int64, kind string, address string) {
	cache := g.DB(dao.RiskcontrolRule.Group()).GetCache()
	if cache != nil {
		g.DB(dao.RiskcontrolRule.Group()).GetCache().Remove(ctx, "SelectCache:"+dao.Contractrule.Table()+strconv.FormatInt(ChainId, 10)+kind+address)
	}
}
func (s *RiskCtrlRule) GetContractRuleBriefs(ctx context.Context, ChainId int64, kind string) ([]*entity.Contractrule, error) {
	model := dao.Contractrule.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.Contractrule.Table() + strconv.FormatInt(ChainId, 10) + kind,
		Force:    false,
	}).Fields(
		dao.Contractrule.Columns().ChainId,
		dao.Contractrule.Columns().ContractAddress,
		dao.Contractrule.Columns().ContractName,
		dao.Contractrule.Columns().ContractKind,
	)
	if ChainId > 0 {
		model = model.Where(dao.Contractrule.Columns().ChainId, ChainId)
	}
	if kind != "" {
		model = model.Where(dao.Contractrule.Columns().ContractKind, kind)
	}
	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	///
	rule := []*entity.Contractrule{}
	rst.Structs(&rule)
	return rule, nil
}

// /
func (s *RiskCtrlRule) GetContractRule(ctx context.Context, ChainId int64, address string, flush bool) (*entity.Contractrule, error) {
	rst, err := dao.Contractrule.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: func() time.Duration {
			if flush {
				return -1
			} else {
				return s.dur
			}
		}(),
		Name:  dao.Contractrule.Table() + strconv.FormatInt(ChainId, 10) + address,
		Force: false,
	}).Where(dao.Contractrule.Columns().ChainId, ChainId).
		Where(dao.Contractrule.Columns().ContractAddress, address).One()
	if err != nil {
		return nil, err
	}
	// /
	rule := &entity.Contractrule{}
	rst.Struct(&rule)
	return rule, nil
}
func (s *RiskCtrlRule) GetRiskCtrlRuleBriefs(ctx context.Context) ([]*entity.RiskcontrolRule, error) {
	model := dao.RiskcontrolRule.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dur,
		Name:     dao.RiskcontrolRule.Table(),
		Force:    true,
	}).Fields(
		dao.RiskcontrolRule.Columns().Id,
		dao.RiskcontrolRule.Columns().RuleName,
		dao.RiskcontrolRule.Columns().Salience,
	)

	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	///
	rule := []*entity.RiskcontrolRule{}
	rst.Structs(&rule)
	return rule, nil
}

// /
func (s *RiskCtrlRule) GetRiskCtrlRule(ctx context.Context, id int, flush bool) (*entity.RiskcontrolRule, error) {
	ids := strconv.Itoa(id)
	rst, err := dao.RiskcontrolRule.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: func() time.Duration {
			if flush {
				return -1
			} else {
				return s.dur
			}
		}(),
		Name:  dao.RiskcontrolRule.Table() + ids,
		Force: true,
	}).Where(dao.RiskcontrolRule.Columns().Id, ids).One()
	// Where(dao.RiskcontrolRule.Columns().RuleName, RuleName).One()
	if err != nil {
		return nil, err
	}
	// /
	rule := &entity.RiskcontrolRule{}
	rst.Struct(&rule)
	return rule, nil
}
