package mpcdao

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type ChainCfg struct {
}

func (s *ChainCfg) UpdateHeigh(ctx context.Context, chainId int64, heigh int64) error {
	_, err := dao.Chaincfg.Ctx(ctx).
		Data(g.Map{dao.Chaincfg.Columns().Heigh: heigh}).
		Where(dao.Chaincfg.Columns().ChainId, chainId).
		OnConflict(dao.Chaincfg.Columns().ChainId).
		Save()
	return err
}
func (s *ChainCfg) AllCfg(ctx context.Context) ([]*entity.Chaincfg, error) {
	rst, err := dao.Chaincfg.Ctx(ctx).All()
	if err != nil {
		return nil, err
	}
	data := make([]*entity.Chaincfg, 0)
	err = rst.Structs(&data)
	return data, err
}
func (s *ChainCfg) GetCfg(ctx context.Context, chainId int64) (*entity.Chaincfg, error) {

	rst, err := dao.Chaincfg.Ctx(ctx).Where(dao.Chaincfg.Columns().ChainId, chainId).One()
	if err != nil {
		return nil, err
	}
	////
	data := &entity.Chaincfg{}
	err = rst.Struct(&data)
	return data, err

}
func NewChainCfg() *ChainCfg {
	return &ChainCfg{}
}
