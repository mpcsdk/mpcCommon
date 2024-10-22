package mpcdao

import (
	"context"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type NftHolding struct {
	redis *gredis.Redis
	dur   time.Duration
}
type QueryNftHolding struct {
	ChainId   int64    `json:"chainId"`
	Address   string   `json:"address"`
	Contract  string   `json:"contract"`
	Contracts []string `json:"contracts"`
	Kinds     []string `json:"kinds"`
	///
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
type NftTransafer struct {
	ChainId     int64  `json:"chainId"`
	From        string `json:"from"`
	To          string `json:"to"`
	Contract    string `json:"contract"`
	TokenId     string `json:"tokenId"`
	Value       int64  `json:"value"`
	BlockNumber int64  `json:"blockNumber"`
	Kind        string `json:"kind"`
	Topic       string `json:"topic"`
}

func NewNftHolding() *NftHolding {

	return &NftHolding{}
}

// //
func nftHoldingKey(chainId int64, address string, contract string, tokenId string) string {
	return dao.NftHolding.Table() + strconv.FormatInt(chainId, 10) + ":" + address + ":" + contract + ":" + tokenId
}
func (s *NftHolding) UpdateTransfer1155(ctx context.Context, tx *entity.NftHolding) error {
	rst, err := dao.NftHolding.Ctx(ctx).Data(tx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     nftHoldingKey(tx.ChainId, tx.Address, tx.Contract, tx.TokenId),
		Force:    false,
	}).
		Where(dao.NftHolding.Columns().ChainId, tx.ChainId).
		Where(dao.NftHolding.Columns().Address, tx.Address).
		Where(dao.NftHolding.Columns().Contract, tx.Contract).
		Where(dao.NftHolding.Columns().TokenId, tx.TokenId).
		Increment(dao.NftHolding.Columns().Value, tx.Value)
	// OnConflict(dao.NftHolding.Columns().ChainId, dao.NftHolding.Columns().Address, dao.NftHolding.Columns().Contract, dao.NftHolding.Columns().TokenId).
	// OnDuplicate(dao.NftHolding.Columns().Value, dao.NftHolding.Columns().UpdatedAt).
	// Save()
	cnt, err := rst.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		_, err = dao.NftHolding.Ctx(ctx).Data(tx).Insert()
	}
	return err
}
func (s *NftHolding) UpdateTransfer721(ctx context.Context, tx *entity.NftHolding) error {
	rst, err := dao.NftHolding.Ctx(ctx).Data(tx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     nftHoldingKey(tx.ChainId, tx.Address, tx.Contract, tx.TokenId),
		Force:    false,
	}).
		Where(dao.NftHolding.Columns().ChainId, tx.ChainId).
		Where(dao.NftHolding.Columns().Address, tx.Address).
		Where(dao.NftHolding.Columns().Contract, tx.Contract).
		Where(dao.NftHolding.Columns().TokenId, tx.TokenId).
		OnConflict(dao.NftHolding.Columns().Address, dao.NftHolding.Columns().Contract, dao.NftHolding.Columns().TokenId, dao.NftHolding.Columns().ChainId).
		OnDuplicate(dao.NftHolding.Columns().Value, dao.NftHolding.Columns().UpdatedAt).
		Save()
	cnt, err := rst.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		_, err = dao.NftHolding.Ctx(ctx).Data(tx).Insert()
	}
	return err
}

// //
func (s *NftHolding) UpdateStat(ctx context.Context, stat *entity.NftHoldingStat) error {
	_, err := dao.NftHoldingStat.Ctx(ctx).Data(&entity.NftHoldingStat{
		ChainId:     stat.ChainId,
		BlockNumber: stat.BlockNumber,
	}).
		Where(dao.NftHolding.Columns().ChainId, stat.ChainId).
		OnConflict(dao.NftHolding.Columns().ChainId).
		Save()

	return err
}
func (s *NftHolding) GetStat(ctx context.Context, chainId int64) (*entity.NftHoldingStat, error) {
	rst, err := dao.NftHoldingStat.Ctx(ctx).
		Where(dao.NftHolding.Columns().ChainId, chainId).One()

	if err != nil {
		return nil, err
	}
	if rst.IsEmpty() {
		return nil, nil
	}
	data := &entity.NftHoldingStat{}
	err = rst.Struct(&data)

	return data, err
}

func (s *NftHolding) Query(ctx context.Context, query *QueryNftHolding) ([]*entity.NftHolding, error) {
	if query.PageSize <= 0 || query.Page < 0 {
		return nil, nil
	}
	if query.Address == "" {
		return nil, nil
	}
	//
	where := dao.NftHolding.Ctx(ctx)
	where = where.Where(dao.NftHolding.Columns().Address, query.Address)

	if query.ChainId > 0 {
		where = where.Where(dao.NftHolding.Columns().ChainId, query.ChainId)
	}
	if len(query.Kinds) > 0 {
		where = where.Where(dao.NftHolding.Columns().Kind, query.Kinds)
	}
	if len(query.Contracts) > 0 {
		where = where.Where(dao.NftHolding.Columns().Contract, query.Contracts)
	}
	where = where.WhereGT(dao.NftHolding.Columns().Value, 0)
	///
	if query.PageSize != 0 {
		where = where.Limit(query.Page*query.PageSize, query.PageSize)
	}
	///
	result, err := where.All()
	if err != nil {
		return nil, err
	}
	data := []*entity.NftHolding{}
	err = result.Structs(&data)
	///
	return data, err
}
func (s *NftHolding) QueryCount(ctx context.Context, query *QueryNftHolding) ([]*entity.NftHolding, error) {
	if query.Address == "" {
		return nil, nil
	}
	//
	where := dao.NftHolding.Ctx(ctx)
	where = where.Where(dao.NftHolding.Columns().Address, query.Address)

	if query.ChainId > 0 {
		where = where.Where(dao.NftHolding.Columns().ChainId, query.ChainId)
	}
	if query.Contract != "" {
		where = where.Where(dao.NftHolding.Columns().Contract, query.Contract)
	} else if len(query.Contracts) > 0 {
		where = where.Where(dao.NftHolding.Columns().Contract, query.Contracts)
	}
	////
	where = where.Fields(
		dao.NftHolding.Columns().Contract,
		`sum("value") as "value"`,
	).
		Group(
			dao.NftHolding.Columns().Contract,
			dao.NftHolding.Columns().Value,
		).
		Having(`"value">0`)
	///
	result, err := where.All()
	if err != nil {
		return nil, err
	}
	data := []*entity.NftHolding{}
	err = result.Structs(&data)
	///
	return data, err
}
