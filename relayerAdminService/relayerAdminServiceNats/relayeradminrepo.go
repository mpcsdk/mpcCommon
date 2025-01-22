package relayerAdminServiceNats

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
type RelayerAdminRepo struct {
	ctx             context.Context
	realayeradminDB *mpcdao.RelayerAdminDB
	///
	apps               map[int]*mq.RelayerAdminAppCfgMsg
	appsRWLock         sync.RWMutex
	assignFees         map[int]*mq.RelayerAdminAssignFeeMsg
	assignFeesRWLock   sync.RWMutex
	specifiedGas       map[int]*mq.RelayerAdminSpecifiedGas
	specifiedGasRWLock sync.RWMutex
	redis              *gredis.Redis
	redisDur           int
}

// //
func NewRelayerAdminRepo(redis *gredis.Redis, cacheDur int) *RelayerAdminRepo {
	////
	ctx := gctx.GetInitCtx()
	////
	s := &RelayerAdminRepo{
		ctx:                ctx,
		apps:               map[int]*mq.RelayerAdminAppCfgMsg{},
		assignFees:         map[int]*mq.RelayerAdminAssignFeeMsg{},
		specifiedGas:       map[int]*mq.RelayerAdminSpecifiedGas{},
		appsRWLock:         sync.RWMutex{},
		assignFeesRWLock:   sync.RWMutex{},
		specifiedGasRWLock: sync.RWMutex{},
		redis:              redis,
		redisDur:           cacheDur,
	}
	s.realayeradminDB = mpcdao.NewRelayerAdminDB(redis, cacheDur)
	////load app cfg
	apps, err := s.realayeradminDB.AllAppCfg(ctx)
	if err != nil {
		panic(err)
	}
	for _, app := range apps {
		s.apps[app.Id] = &mq.RelayerAdminAppCfgMsg{
			Msg:  mq.Msg{Version: 0},
			Data: app,
		}
	}
	//fees
	assignFees, err := s.realayeradminDB.AllAssignFee(ctx)
	if err != nil {
		panic(err)
	}
	for _, fee := range assignFees {
		s.assignFees[fee.Id] = &mq.RelayerAdminAssignFeeMsg{
			Msg:  mq.Msg{Version: 0},
			Data: fee,
		}
	}
	///gas
	gases, err := s.realayeradminDB.AllSpecifiedGas(ctx)
	if err != nil {
		panic(err)
	}
	for _, gas := range gases {
		s.specifiedGas[gas.Id] = &mq.RelayerAdminSpecifiedGas{
			Msg:  mq.Msg{Version: 0},
			Data: gas,
		}
	}
	return s
}
func (s *RelayerAdminRepo) Dump() string {
	return fmt.Sprintf("apps: %v\nassignFees: %v\nspecifiedGas: %v\n", s.apps, s.assignFees, s.specifiedGas)
}
func (s *RelayerAdminRepo) SetApp(id int, data *mq.RelayerAdminAppCfgMsg) {
	///
	switch data.Opt {
	case mq.OptAdd:
		s.appsRWLock.Lock()
		defer s.appsRWLock.Unlock()
		s.apps[data.Data.Id] = data
	case mq.OptUpdate:
		///
		s.appsRWLock.Lock()
		defer s.appsRWLock.Unlock()
		////
		needReload := false
		if appcfg, ok := s.apps[data.Data.Id]; !ok {
			needReload = true
		} else {
			if appcfg.Version != data.Version {
				needReload = true
			}
		}
		//////
		if needReload {
			app, err := s.realayeradminDB.GetAppCfgById(s.ctx, data.Data.Id)
			if err != nil {
				g.Log().Error(s.ctx, "GetAppCfgById err:", err)
				return
			}
			s.apps[data.Data.Id] = &mq.RelayerAdminAppCfgMsg{
				Msg: mq.Msg{
					Version: data.Version + 1,
				},
				Data: app,
			}
			return
		}
		////mergepath
		app := s.apps[data.Data.Id]
		orgData, err := json.Marshal(app.Data)
		if err != nil {
			g.Log().Error(s.ctx, "SetApp json.Marshal error:", err)
			return
		}
		merged, err := jsonpatch.MergePatch(orgData, []byte(data.JsonPatch))
		if err != nil {
			g.Log().Error(s.ctx, "SetApp jsonpatch.MergePatch error:", err)
			return
		}
		mergedData := &entity.RelayeradminAppCfg{}
		if err := json.Unmarshal(merged, mergedData); err != nil {
			g.Log().Error(s.ctx, "SetChain json.Unmarshal error:", err)
			return
		}
		app.Data = mergedData
		app.Version += 1
	case mq.OptDelete:
	case mq.OptCheck:
	}

}
func (s *RelayerAdminRepo) GetApp(id int) *entity.RelayeradminAppCfg {
	s.appsRWLock.RLock()
	defer s.appsRWLock.RUnlock()
	app := s.apps[id]
	if app == nil {
		return nil
	}
	return app.Data
}
func (s *RelayerAdminRepo) AllApps() map[string]*entity.RelayeradminAppCfg {
	s.appsRWLock.RLock()
	defer s.appsRWLock.RUnlock()
	rst := map[string]*entity.RelayeradminAppCfg{}
	for _, app := range s.apps {
		rst[app.Data.AppId] = app.Data
	}
	return rst
}

// //
func (s *RelayerAdminRepo) SetAssignFee(id int, data *mq.RelayerAdminAssignFeeMsg) {

	switch data.Opt {
	case mq.OptAdd:
		s.assignFeesRWLock.Lock()
		defer s.assignFeesRWLock.Unlock()
		s.assignFees[data.Data.Id] = data
	case mq.OptUpdate:
		///
		s.assignFeesRWLock.Lock()
		defer s.assignFeesRWLock.Unlock()
		////
		needReload := false
		if assignFee, ok := s.assignFees[data.Data.Id]; !ok {
			needReload = true
		} else {
			if assignFee.Version != data.Version {
				needReload = true
			}
		}
		//////
		if needReload {
			assignFee, err := s.realayeradminDB.GetAssignFeeById(s.ctx, data.Data.Id)
			if err != nil {
				g.Log().Error(s.ctx, "GetAssignFeeById err:", err)
				return
			}
			s.assignFees[data.Data.Id] = &mq.RelayerAdminAssignFeeMsg{
				Msg: mq.Msg{
					Version: data.Version + 1,
				},
				Data: assignFee,
			}
			return
		}
		////mergepath
		assignFee := s.assignFees[data.Data.Id]
		orgData, err := json.Marshal(assignFee.Data)
		if err != nil {
			g.Log().Error(s.ctx, "SetAssignFee json.Marshal error:", err)
			return
		}
		merged, err := jsonpatch.MergePatch(orgData, []byte(data.JsonPatch))
		if err != nil {
			g.Log().Error(s.ctx, "SetAssignFee jsonpatch.MergePatch error:", err)
			return
		}
		mergedData := &entity.RelayeradminAssignFee{}
		if err := json.Unmarshal(merged, mergedData); err != nil {
			g.Log().Error(s.ctx, "SetContract json.Unmarshal error:", err)
			return
		}
		assignFee.Data = mergedData
		assignFee.Version += 1
	case mq.OptDelete:
	case mq.OptCheck:
	}
}
func (s *RelayerAdminRepo) AllAssignFee() map[string]*entity.RelayeradminAssignFee {
	s.assignFeesRWLock.RLock()
	defer s.assignFeesRWLock.RUnlock()
	rst := map[string]*entity.RelayeradminAssignFee{}
	for _, assignFee := range s.assignFees {
		rst[mpcdaoutil.RelayerAdminAssignFeeKey(assignFee.Data.ChainId, assignFee.Data.AppId)] = assignFee.Data
	}
	return rst
}

func (s *RelayerAdminRepo) GetAssignFee(id int) *entity.RelayeradminAssignFee {
	s.assignFeesRWLock.RLock()
	defer s.assignFeesRWLock.RUnlock()
	assignFee := s.assignFees[id]
	if assignFee == nil {
		return nil
	}
	return assignFee.Data
}

// /specified gas
func (s *RelayerAdminRepo) SetSpecifiedGas(id int, data *mq.RelayerAdminSpecifiedGas) {

	switch data.Opt {
	case mq.OptAdd:
	case mq.OptUpdate:
		///
		s.specifiedGasRWLock.Lock()
		defer s.specifiedGasRWLock.Unlock()
		////
		needReload := false
		if gasUsed, ok := s.specifiedGas[data.Data.Id]; !ok {
			needReload = true
		} else {
			if gasUsed.Version != data.Version {
				needReload = true
			}
		}
		//////
		if needReload {
			gasUsed, err := s.realayeradminDB.GetSpecifiedGasById(s.ctx, data.Data.Id)
			if err != nil {
				g.Log().Error(s.ctx, "GetSpecifiedGasById err:", err)
				return
			}
			s.specifiedGas[data.Data.Id] = &mq.RelayerAdminSpecifiedGas{
				Msg: mq.Msg{
					Version: data.Version + 1,
				},
				Data: gasUsed,
			}
			return
		}
		////mergepath
		gasused := s.specifiedGas[data.Data.Id]
		orgData, err := json.Marshal(gasused.Data)
		if err != nil {
			g.Log().Error(s.ctx, "SetSpecifiedGas json.Marshal error:", err)
			return
		}
		merged, err := jsonpatch.MergePatch(orgData, []byte(data.JsonPatch))
		if err != nil {
			g.Log().Error(s.ctx, "SetSpecifiedGas jsonpatch.MergePatch error:", err)
			return
		}
		mergedData := &entity.RelayeradminSpecifiedGas{}
		if err := json.Unmarshal(merged, mergedData); err != nil {
			g.Log().Error(s.ctx, "SetSpecifiedGas json.Unmarshal error:", err)
			return
		}
		gasused.Data = mergedData
		gasused.Version += 1
	case mq.OptDelete:
	case mq.OptCheck:
	}

}
func (s *RelayerAdminRepo) GetSpecifiedGas(id int) *entity.RelayeradminSpecifiedGas {
	s.specifiedGasRWLock.RLock()
	defer s.specifiedGasRWLock.RUnlock()
	gasused := s.specifiedGas[id]
	if gasused == nil {
		return nil
	}
	return gasused.Data
}
func (s *RelayerAdminRepo) AllSpecifiedGas() []*entity.RelayeradminSpecifiedGas {
	s.specifiedGasRWLock.RLock()
	defer s.specifiedGasRWLock.RUnlock()

	rst := []*entity.RelayeradminSpecifiedGas{}
	for _, rule := range s.specifiedGas {
		rst = append(rst, rule.Data)
	}
	return rst
}

// /
func (s *RelayerAdminRepo) RelayerAdminDB() *mpcdao.RelayerAdminDB {
	return s.realayeradminDB
}
