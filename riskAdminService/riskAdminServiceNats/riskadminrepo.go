package riskAdminServiceNats

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpcdao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	mpcdaoutil "github.com/mpcsdk/mpcCommon/mpcdao/util"
	"github.com/mpcsdk/mpcCommon/mq"
)

// /
type RiskAdminRepo struct {
	ctx             context.Context
	riskadminDB     *mpcdao.RiskAdminDB
	chains          map[int]*mq.RiskAdminChainMsg
	chainsRWLock    sync.RWMutex
	contracts       map[int]*mq.RiskAdminContractMsg
	contractsRWLock sync.RWMutex
	riskRules       map[int]*mq.RiskAdminRiskRuleMsg
	riskRulesRWLock sync.RWMutex
	redis           *gredis.Redis
	redisDur        int
}

// //
func NewRiskAdminRepo(redis *gredis.Redis, cacheDur int) *RiskAdminRepo {
	////
	ctx := gctx.GetInitCtx()
	if redis != nil {
		_, err := redis.Conn(context.Background())
		if err != nil {
			panic(err)
		}
	}
	////
	s := &RiskAdminRepo{
		ctx:             ctx,
		chains:          make(map[int]*mq.RiskAdminChainMsg),
		contracts:       make(map[int]*mq.RiskAdminContractMsg),
		riskRules:       make(map[int]*mq.RiskAdminRiskRuleMsg),
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
		s.chains[chain.Id] = &mq.RiskAdminChainMsg{
			Msg:  mq.Msg{Version: 0},
			Data: chain,
		}
	}
	//contract
	contracts, err := s.riskadminDB.AllContractAbi(ctx)
	if err != nil {
		panic(err)
	}
	for _, contract := range contracts {
		s.contracts[contract.Id] = &mq.RiskAdminContractMsg{
			Msg:  mq.Msg{Version: 0},
			Data: contract,
		}
	}
	///riskrule
	ruls, err := s.riskadminDB.AllRiskCtrlRule(ctx)
	if err != nil {
		panic(err)
	}
	for _, rule := range ruls {
		s.riskRules[rule.Id] = &mq.RiskAdminRiskRuleMsg{
			Msg:  mq.Msg{Version: 0},
			Data: rule,
		}
	}
	return s
}
func (s *RiskAdminRepo) Dump() string {
	return fmt.Sprintf("chains: %v\ncontracts: %v\nriskRules: %v\n", s.chains, s.contracts, s.riskRules)
}
func (s *RiskAdminRepo) SetChain(id int, data *mq.RiskAdminChainMsg) {
	///
	switch data.Opt {
	case mq.OptAdd:
		s.chainsRWLock.Lock()
		defer s.chainsRWLock.Unlock()
		s.chains[data.Data.Id] = data
	case mq.OptUpdate:
		///
		s.chainsRWLock.Lock()
		defer s.chainsRWLock.Unlock()
		////
		needReload := false
		if chain, ok := s.chains[data.Data.Id]; !ok {
			needReload = true
		} else {
			if chain.Version != data.Version {
				needReload = true
			}
		}
		//////
		if needReload {
			chain, err := s.riskadminDB.GetChainCfgById(s.ctx, data.Data.Id)
			if err != nil {
				g.Log().Error(s.ctx, "GetChainCfgById err:", err)
				return
			}
			s.chains[data.Data.Id] = &mq.RiskAdminChainMsg{
				Msg: mq.Msg{
					Version: data.Version + 1,
				},
				Data: chain,
			}
			return
		}
		////mergepath
		chain := s.chains[data.Data.Id]
		orgData, err := json.Marshal(chain.Data)
		if err != nil {
			g.Log().Error(s.ctx, "SetChain json.Marshal error:", err)
			return
		}
		merged, err := jsonpatch.MergePatch(orgData, []byte(data.JsonPatch))
		if err != nil {
			g.Log().Error(s.ctx, "SetChain jsonpatch.MergePatch error:", err)
			return
		}
		mergedData := &entity.RiskadminChaincfg{}
		if err := json.Unmarshal(merged, mergedData); err != nil {
			g.Log().Error(s.ctx, "SetChain json.Unmarshal error:", err)
			return
		}
		chain.Data = mergedData
		chain.Version += 1
	case mq.OptDelete:
		s.chainsRWLock.Lock()
		defer s.chainsRWLock.Unlock()
		delete(s.chains, data.Data.Id)
	case mq.OptCheck:
	}

}
func (s *RiskAdminRepo) GetChain(chainId int64) *entity.RiskadminChaincfg {
	s.chainsRWLock.RLock()
	defer s.chainsRWLock.RUnlock()
	for _, chain := range s.chains {
		if chain.Data.ChainId == chainId {
			return chain.Data
		}
	}
	return nil
}
func (s *RiskAdminRepo) GetChainById(id int) *entity.RiskadminChaincfg {
	s.chainsRWLock.RLock()
	defer s.chainsRWLock.RUnlock()
	chain := s.chains[id]
	if chain == nil {
		return nil
	}
	return chain.Data
}
func (s *RiskAdminRepo) AllChain() map[int64]*entity.RiskadminChaincfg {
	s.chainsRWLock.RLock()
	defer s.chainsRWLock.RUnlock()
	rst := map[int64]*entity.RiskadminChaincfg{}
	for _, chain := range s.chains {
		rst[chain.Data.ChainId] = chain.Data
	}
	return rst
}

// //
func (s *RiskAdminRepo) SetContract(id int, data *mq.RiskAdminContractMsg) {

	switch data.Opt {
	case mq.OptAdd:
		s.contractsRWLock.Lock()
		defer s.contractsRWLock.Unlock()
		s.contracts[data.Data.Id] = data
	case mq.OptUpdate:
		///
		s.contractsRWLock.Lock()
		defer s.contractsRWLock.Unlock()
		////
		needReload := false
		if contract, ok := s.contracts[data.Data.Id]; !ok {
			needReload = true
		} else {
			if contract.Version != data.Version {
				needReload = true
			}
		}
		//////
		if needReload {
			contract, err := s.riskadminDB.GetContractAbiById(s.ctx, data.Data.Id)
			if err != nil {
				g.Log().Error(s.ctx, "GetContractAbiById err:", err)
				return
			}
			s.contracts[data.Data.Id] = &mq.RiskAdminContractMsg{
				Msg: mq.Msg{
					Version: data.Version + 1,
				},
				Data: contract,
			}
			return
		}
		////mergepath
		contract := s.contracts[data.Data.Id]
		orgData, err := json.Marshal(contract.Data)
		if err != nil {
			g.Log().Error(s.ctx, "SetContract json.Marshal error:", err)
			return
		}
		merged, err := jsonpatch.MergePatch(orgData, []byte(data.JsonPatch))
		if err != nil {
			g.Log().Error(s.ctx, "SetContract jsonpatch.MergePatch error:", err)
			return
		}
		mergedData := &entity.RiskadminContractabi{}
		if err := json.Unmarshal(merged, mergedData); err != nil {
			g.Log().Error(s.ctx, "SetContract json.Unmarshal error:", err)
			return
		}
		contract.Data = mergedData
		contract.Version += 1
	case mq.OptDelete:
		s.contractsRWLock.Lock()
		defer s.contractsRWLock.Unlock()
		delete(s.contracts, data.Data.Id)
	case mq.OptCheck:
	}
}
func (s *RiskAdminRepo) AllContract() map[string]*entity.RiskadminContractabi {
	s.contractsRWLock.RLock()
	defer s.contractsRWLock.RUnlock()
	rst := map[string]*entity.RiskadminContractabi{}
	for _, contract := range s.contracts {
		rst[mpcdaoutil.RiskadminContractabiKey(contract.Data.ChainId, contract.Data.ContractAddress)] = contract.Data
	}
	return rst
}
func (s *RiskAdminRepo) GetContractByChainId(chainId int64) []*entity.RiskadminContractabi {
	s.contractsRWLock.RLock()
	defer s.contractsRWLock.RUnlock()
	rst := []*entity.RiskadminContractabi{}
	if chainId > 0 {
		for _, contract := range s.contracts {
			if contract.Data.ChainId == chainId {
				rst = append(rst, contract.Data)
			}
		}
		return rst
	}
	return nil
}
func (s *RiskAdminRepo) GetContract(chainId int64, contractAddr string) *entity.RiskadminContractabi {
	s.contractsRWLock.RLock()
	defer s.contractsRWLock.RUnlock()
	// rst := map[string]*entity.RiskadminContractabi{}
	if chainId > 0 {
		for _, contract := range s.contracts {
			if contract.Data.ContractAddress == contractAddr && contract.Data.ChainId == chainId {
				return contract.Data
			}
		}
	} else {
		for _, contract := range s.contracts {
			if contract.Data.ContractAddress == contractAddr {
				return contract.Data
			}
		}
	}
	return nil
}
func (s *RiskAdminRepo) GetContractById(id int) *entity.RiskadminContractabi {
	s.contractsRWLock.RLock()
	defer s.contractsRWLock.RUnlock()
	contract := s.contracts[id]
	if contract == nil {
		return nil
	}
	return contract.Data
}

// /
func (s *RiskAdminRepo) SetRiskRule(id int, data *mq.RiskAdminRiskRuleMsg) {

	switch data.Opt {
	case mq.OptAdd:
	case mq.OptUpdate:
		///
		s.riskRulesRWLock.Lock()
		defer s.riskRulesRWLock.Unlock()
		////
		needReload := false
		if rule, ok := s.riskRules[data.Data.Id]; !ok {
			needReload = true
		} else {
			if rule.Version != data.Version {
				needReload = true
			}
		}
		//////
		if needReload {
			riskRule, err := s.riskadminDB.GetRiskCtrlRuleById(s.ctx, data.Data.Id)
			if err != nil {
				g.Log().Error(s.ctx, "GetRiskCtrlRuleById err:", err)
				return
			}
			s.riskRules[data.Data.Id] = &mq.RiskAdminRiskRuleMsg{
				Msg: mq.Msg{
					Version: data.Version + 1,
				},
				Data: riskRule,
			}
			return
		}
		////mergepath
		rule := s.riskRules[data.Data.Id]
		orgData, err := json.Marshal(rule.Data)
		if err != nil {
			g.Log().Error(s.ctx, "SetRiskRule json.Marshal error:", err)
			return
		}
		merged, err := jsonpatch.MergePatch(orgData, []byte(data.JsonPatch))
		if err != nil {
			g.Log().Error(s.ctx, "SetRiskRule jsonpatch.MergePatch error:", err)
			return
		}
		mergedData := &entity.RiskadminRiskcontrolRule{}
		if err := json.Unmarshal(merged, mergedData); err != nil {
			g.Log().Error(s.ctx, "SetRiskRule json.Unmarshal error:", err)
			return
		}
		rule.Data = mergedData
		rule.Version += 1
	case mq.OptDelete:
		s.riskRulesRWLock.Lock()
		defer s.riskRulesRWLock.Unlock()
		delete(s.riskRules, data.Data.Id)
	case mq.OptCheck:
	}

}
func (s *RiskAdminRepo) GetRiskRule(id int) *entity.RiskadminRiskcontrolRule {
	s.riskRulesRWLock.RLock()
	defer s.riskRulesRWLock.RUnlock()
	rule := s.riskRules[id]
	if rule == nil {
		return nil
	}
	return rule.Data
}
func (s *RiskAdminRepo) AllRiskRule() []*entity.RiskadminRiskcontrolRule {
	s.riskRulesRWLock.RLock()
	defer s.riskRulesRWLock.RUnlock()

	rst := []*entity.RiskadminRiskcontrolRule{}
	for _, rule := range s.riskRules {
		rst = append(rst, rule.Data)
	}
	return rst
}

// /
func (s *RiskAdminRepo) RiskadminDB() *mpcdao.RiskAdminDB {
	return s.riskadminDB
}
