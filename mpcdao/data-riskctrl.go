package mpcdao

import (
	"context"

	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

func GetAll(ctx context.Context) ([]*entity.ChainData, error) {
	rst, err := dao.ChainData.Ctx(ctx).All()
	if err != nil {
		return nil, err
	}
	///
	data := []*entity.ChainData{}
	err = rst.Structs(&data)
	if err != nil {
		return nil, err
	}
	//
	return data, err
}

func InsertTx(ctx context.Context, tx *entity.ChainData) error {
	_, err := dao.ChainData.Ctx(ctx).Data(tx).Insert()
	return err
}

// /
type QueryAggNft struct {
}

func GetAggNft(ctx context.Context, query *QueryAggNft) (*entity.AggNft, error) {

	return nil, nil
}
func UpSertAggNft(ctx context.Context, tx *entity.AggNft) error {
	_, err := dao.AggNft.Ctx(ctx).
		Data(tx).
		OnConflict(
			dao.AggNft.Columns().ChainId,
			dao.AggNft.Columns().FromAddr,
			dao.AggNft.Columns().Contract,
		).Save()
	return err
}

// /
func GetAggFt(ctx context.Context) (*entity.AggNft, error) {

	return nil, nil
}
func UpSertAggFt(ctx context.Context, tx *entity.AggNft) error {
	_, err := dao.AggFt.Ctx(ctx).
		Data(tx).
		OnConflict(
			dao.AggFt.Columns().ChainId,
			dao.AggFt.Columns().FromAddr,
			dao.AggFt.Columns().Contract,
		).Save()
	return err
}

///
