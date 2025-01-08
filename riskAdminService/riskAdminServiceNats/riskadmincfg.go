package riskAdminServiceNats

import (
	"context"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpcdao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	"github.com/mpcsdk/mpcCommon/mq"
)

func contractKey(contract *entity.RiskadminContractabi) string {
	return fmt.Sprint(contract.ChainId, contract.ContractAddress, contract.ContractKind)
}

// /
type RiskAdminCfg struct {
	ctx             context.Context
	riskadminDB     *mpcdao.RiskAdminDB
	chains          map[int64]*entity.RiskadminChaincfg
	chainsRWLock    sync.RWMutex
	contracts       map[string]*entity.RiskadminContractabi
	contractsRWLock sync.RWMutex
	riskRules       map[int]*entity.RiskadminRiskcontrolRule
	riskRulesRWLock sync.RWMutex
	redis           *gredis.Redis
	redisDur        int
}

// //
func NewRiskAdminCfg(redis *gredis.Redis, cacheDur int) *RiskAdminCfg {
	////
	ctx := gctx.GetInitCtx()
	_, err := redis.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	////
	s := &RiskAdminCfg{
		ctx:             ctx,
		chains:          make(map[int64]*entity.RiskadminChaincfg),
		contracts:       make(map[string]*entity.RiskadminContractabi),
		riskRules:       make(map[int]*entity.RiskadminRiskcontrolRule),
		chainsRWLock:    sync.RWMutex{},
		contractsRWLock: sync.RWMutex{},
		riskRulesRWLock: sync.RWMutex{},
		redis:           redis,
		redisDur:        cacheDur,
	}
	s.riskadminDB = mpcdao.NewRiskAdminDB(redis, cacheDur)
	////load all cfg
	chains, err := s.riskadminDB.AllChainsCfg(ctx)
	if err != nil {
		panic(err)
	}
	for _, chain := range chains {
		s.chains[chain.ChainId] = chain
	}
	//contract
	contracts, err := s.riskadminDB.AllContractAbi(ctx)
	if err != nil {
		panic(err)
	}
	for _, contract := range contracts {
		s.contracts[contractKey(contract)] = contract
	}
	///riskrule
	ruls, err := s.riskadminDB.AllRiskCtrlRule(ctx)
	if err != nil {
		panic(err)
	}
	for _, rule := range ruls {
		s.riskRules[rule.Id] = rule
	}
	return s
}
func (s *RiskAdminCfg) Dump() string {
	return fmt.Sprintf("chains: %v\ncontracts: %v\nriskRules: %v\n", s.chains, s.contracts, s.riskRules)
}
func (s *RiskAdminCfg) SetChain(id int, data *mq.RiskAdminChainMsg) {
	///
	switch data.Opt {
	case mq.OptAdd, mq.OptUpdate:
		///
		chain, err := s.riskadminDB.GetChainCfgById(s.ctx, id)
		if err != nil {
			g.Log().Error(s.ctx, "GetChainCfgById err:", err)
			return
		}
		s.chainsRWLock.Lock()
		defer s.chainsRWLock.Unlock()
		s.chains[chain.ChainId] = chain
	case mq.OptDelete:
	case mq.OptCheck:
	}

}
func (s *RiskAdminCfg) GetChain(chainId int64) *entity.RiskadminChaincfg {
	s.chainsRWLock.RLock()
	defer s.chainsRWLock.RUnlock()
	return s.chains[chainId]
}

// //
func (s *RiskAdminCfg) SetContract(id int, data *mq.RiskAdminContractMsg) {

	switch data.Opt {
	case mq.OptAdd, mq.OptUpdate:
		///
		contract, err := s.riskadminDB.GetContractAbiById(s.ctx, id)
		if err != nil {
			g.Log().Error(s.ctx, "GetChainCfgById err:", err)
			return
		}
		s.chainsRWLock.Lock()
		defer s.chainsRWLock.Unlock()
		s.contracts[contractKey(contract)] = contract
	case mq.OptDelete:
	case mq.OptCheck:
	}
}
func (s *RiskAdminCfg) AllContract() []*entity.RiskadminContractabi {
	s.contractsRWLock.RLock()
	defer s.contractsRWLock.RUnlock()
	rst := []*entity.RiskadminContractabi{}
	for _, contract := range s.contracts {
		rst = append(rst, contract)
	}
	return rst
}

// todo: key
func (s *RiskAdminCfg) GetContract(key string) *entity.RiskadminContractabi {
	s.contractsRWLock.RLock()
	defer s.contractsRWLock.RUnlock()
	return s.contracts[key]
}

// /
func (s *RiskAdminCfg) SetRiskRule(id int, data *mq.RiskAdminRiskRuleMsg) {

	switch data.Opt {
	case mq.OptAdd, mq.OptUpdate:
		///
		riskrule, err := s.riskadminDB.GetRiskCtrlRuleById(s.ctx, id)
		if err != nil {
			g.Log().Error(s.ctx, "GetChainCfgById err:", err)
			return
		}
		s.chainsRWLock.Lock()
		defer s.chainsRWLock.Unlock()
		s.riskRules[riskrule.Id] = riskrule
	case mq.OptDelete:
	case mq.OptCheck:
	}

}
func (s *RiskAdminCfg) GetRiskRule(id int) *entity.RiskadminRiskcontrolRule {
	s.riskRulesRWLock.RLock()
	defer s.riskRulesRWLock.RUnlock()
	return s.riskRules[id]
}
func (s *RiskAdminCfg) AllRiskRule() []*entity.RiskadminRiskcontrolRule {
	s.riskRulesRWLock.RLock()
	defer s.riskRulesRWLock.RUnlock()

	rst := []*entity.RiskadminRiskcontrolRule{}
	for _, rule := range s.riskRules {
		rst = append(rst, rule)
	}
	return rst
}

///
