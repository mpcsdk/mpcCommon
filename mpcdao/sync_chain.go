package mpcdao

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lib/pq"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type ChainTransfer struct {
	chainId  int64
	dbname   string
	dbmod    *gdb.Model
	statemod *gdb.Model
	redis    *gredis.Redis
	dur      time.Duration
}
type QueryData struct {
	ChainId  int64    `json:"chainId"`
	From     string   `json:"from"`
	To       string   `json:"to"`
	Contract string   `json:"contract"`
	Kinds    []string `json:"kinds"`
	///
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	///
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

// ///
func InitSyncChainDB(ctx context.Context, chainId int64) error {
	dbname := "sync_chain_" + gconv.String(chainId)
	///
	db := g.DB("sync_chain").Schema(dbname)
	_, err := db.Exec(ctx, "select height from state limit 1")
	if err != nil {
		gerr := err.(*gerror.Error)
		if pgerr, ok := gerr.Cause().(*pq.Error); ok {
			if pgerr.Code == "3D000" {
				_, err = dao.SyncchainChainTransfer.DB().Exec(ctx, "CREATE DATABASE "+dbname)
				if err != nil {
					return err
				}
			}
			if pgerr.Code == "42P01" || pgerr.Code == "3D000" {
				err = initStateTable(ctx, dbname)
				if err != nil {
					return err
				}
				err = initTransferTable(ctx, dbname)
				if err != nil {
					return err
				}
				_, err = db.Insert(ctx, dao.SyncchainState.Table(), entity.SyncchainState{
					ChainId:      chainId,
					CurrentBlock: 0,
				})
				if err != nil {
					return err
				}
			}
		}
	} else {
		return err
	}
	return nil
	///
}

// ///
func initStateTable(ctx context.Context, dbname string) error {
	_, err := dao.SyncchainChainTransfer.DB().Schema(dbname).Exec(ctx, `
	CREATE TABLE "public"."state" (
		"chain_id" int8 NOT NULL,
		"current_block" int8 NOT NULL,
		"createdat" timestamptz(6) NOT NULL,
		"updatedat" timestamptz(6) NOT NULL,
		CONSTRAINT "state_pkey" PRIMARY KEY ("chain_id")
	)
	;
	ALTER TABLE "public"."state" 
		OWNER TO "postgres";	
		`)
	return err
}
func initTransferTable(ctx context.Context, dbname string) error {
	_, err := dao.SyncchainChainTransfer.DB().Schema(dbname).Exec(ctx, `CREATE TABLE "public"."chain_transfer" (
		"chain_id" int8 NOT NULL,
		"height" int8 NOT NULL,
		"block_hash" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"ts" int8 NOT NULL,
		"tx_hash" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"tx_idx" int4 NOT NULL,
		"log_idx" int4 NOT NULL,
		"from" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"to" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"contract" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"value" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"gas" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"gas_price" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"nonce" int8 NOT NULL,
		"kind" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"token_id" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"removed" bool NOT NULL,
		"status" int8 NOT NULL,
		"traceTag" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
	  )
	  ;
	  
	  ALTER TABLE "public"."chain_transfer" 
		OWNER TO "postgres";
	  
	  CREATE INDEX "chain_transfer_contract_ts" ON "public"."chain_transfer" USING btree (
		"contract" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
		"ts" "pg_catalog"."int8_ops" DESC NULLS LAST
	  );
	  
	  CREATE INDEX "chain_transfer_from_kind_contract_ts_idx" ON "public"."chain_transfer" USING btree (
		"from" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
		"kind" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
		"contract" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
		"ts" "pg_catalog"."int8_ops" DESC NULLS LAST
	  );
	  
	  CREATE INDEX "chain_transfer_height_idx" ON "public"."chain_transfer" USING btree (
		"height" "pg_catalog"."int8_ops" ASC NULLS LAST
	  );
	  
	  CREATE INDEX "chain_transfer_to_kind_contract_ts_idx" ON "public"."chain_transfer" USING btree (
		"to" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
		"kind" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
		"contract" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
		"ts" "pg_catalog"."int8_ops" DESC NULLS LAST
	  );
	  
	  CREATE INDEX "chain_transfer_ts_idx" ON "public"."chain_transfer" USING btree (
		"ts" "pg_catalog"."int8_ops" ASC NULLS LAST
	  );
	  
	  CREATE UNIQUE INDEX "chain_transfer_tx_hash_tx_idx_log_idx_traceTag_token_id_idx" ON "public"."chain_transfer" USING btree (
		"tx_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
		"tx_idx" "pg_catalog"."int4_ops" ASC NULLS LAST,
		"log_idx" "pg_catalog"."int4_ops" ASC NULLS LAST,
		"traceTag" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
	  );
	  `)
	return err
}
func NewChainTransfer(chainId int64, redis *gredis.Redis, dur int) *ChainTransfer {
	dbname := "sync_chain_" + gconv.String(chainId)

	dbmod := dao.SyncchainChainTransfer.DB().Schema(dbname).Model(dao.SyncchainChainTransfer.Table()).Safe()
	statemod := dao.SyncchainChainTransfer.DB().Schema(dbname).Model(dao.SyncchainState.Table()).Safe()
	if redis != nil {
		g.DB(dao.SyncchainState.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))

	}

	return &ChainTransfer{
		dbname:   dbname,
		dbmod:    dbmod,
		statemod: statemod,
		chainId:  chainId,
		redis:    redis,
		dur:      time.Duration(dur) * time.Second,
	}
}

func (s *ChainTransfer) Insert(ctx context.Context, data *entity.SyncchainChainTransfer) error {
	// _, err := dao.ChainTransfer.Ctx(ctx).Insert(data)
	_, err := s.dbmod.Ctx(ctx).Insert(data)
	return err
}
func (s *ChainTransfer) TruncateTransfer(ctx context.Context, chainId int64, number int64) error {
	_, err := s.dbmod.Ctx(ctx).
		Where(dao.SyncchainChainTransfer.Columns().ChainId, chainId).
		WhereLT(dao.SyncchainChainTransfer.Columns().Height, number).
		Delete()
	return err
}
func (s *ChainTransfer) DelChainBlockNumber(ctx context.Context, chainId int64, number int64) error {
	_, err := s.dbmod.Ctx(ctx).
		Where(dao.SyncchainChainTransfer.Columns().ChainId, chainId).
		Where(dao.SyncchainChainTransfer.Columns().Height, number).
		Delete()

	return err
}
func (s *ChainTransfer) InsertBatch(ctx context.Context, data []*entity.SyncchainChainTransfer) error {
	_, err := s.dbmod.Ctx(ctx).Insert(data)
	return err
}

func (s *ChainTransfer) UpTransactionMap(ctx context.Context, data map[int64][]*entity.SyncchainChainTransfer) error {
	if len(data) == 0 {
		return nil
	}
	latest := int64(0)
	for nr, _ := range data {
		if nr > latest {
			latest = nr
		}
	}
	////
	return s.dbmod.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, txs := range data {
			_, err := tx.Insert(dao.SyncchainChainTransfer.Table(), txs)
			if err != nil {
				return err
			}
		}
		///
		_, err := tx.Ctx(ctx).Model(dao.SyncchainState.Table()).
			Where(dao.SyncchainState.Columns().ChainId, s.chainId).
			Data(g.Map{
				dao.SyncchainState.Columns().CurrentBlock: latest,
			}).
			Update()

		return err
	})
}

// /
func (s *ChainTransfer) UpTransaction(ctx context.Context, data []*entity.SyncchainChainTransfer) error {
	if len(data) == 0 {
		return nil
	}
	latest := data[0].Height
	for _, d := range data {
		if d.Height > latest {
			latest = d.Height
		}
	}
	////
	return s.dbmod.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := tx.Insert(dao.SyncchainChainTransfer.Table(), data)
		if err != nil {
			return err
		}
		///
		_, err = tx.Ctx(ctx).Model(dao.SyncchainState.Table()).
			Where(dao.SyncchainState.Columns().ChainId, s.chainId).
			Data(g.Map{
				dao.SyncchainState.Columns().CurrentBlock: latest,
			}).
			Update()

		return err
	})
}
func (s *ChainTransfer) Insert_Transaction(ctx context.Context, data []*entity.SyncchainChainTransfer) error {
	err := s.dbmod.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, transfer := range data {
			// tx.SavePoint(gconv.String(i))
			_, err := tx.Insert(dao.SyncchainChainTransfer.Table(), transfer)
			if err != nil {
				g.Log().Warning(ctx, "Insert_Transaction:", err)
				return err
				// tx.RollbackTo(gconv.String(i))
			}
		}
		return nil
	})

	return err
}
func (s *ChainTransfer) Query(ctx context.Context, query *QueryData) ([]*entity.SyncchainChainTransfer, error) {
	if query.PageSize < 0 || query.Page < 0 {
		return nil, nil
	}
	//
	where := s.dbmod.Ctx(ctx)
	if query.ChainId != 0 {
		where = where.Where(dao.SyncchainChainTransfer.Columns().ChainId, query.ChainId)
	}
	if len(query.Kinds) > 0 {
		where = where.Where(dao.SyncchainChainTransfer.Columns().Kind, query.Kinds)
	}
	if query.From != "" {
		where = where.Where(dao.SyncchainChainTransfer.Columns().From, query.From)
	}
	if query.To != "" {
		where = where.Where(dao.SyncchainChainTransfer.Columns().To, query.To)
	}
	if query.Contract != "" {
		where = where.Where(dao.SyncchainChainTransfer.Columns().Contract, query.Contract)
	}
	///time
	if query.StartTime != 0 {
		where = where.WhereGTE(dao.SyncchainChainTransfer.Columns().Ts, query.StartTime)
	}
	if query.EndTime != 0 {
		where = where.WhereLTE(dao.SyncchainChainTransfer.Columns().Ts, query.EndTime)
	}
	///
	if query.PageSize != 0 {
		where = where.Limit(query.Page*query.PageSize, query.PageSize)
	}
	///
	result, err := where.All()
	if err != nil {
		return nil, err
	}
	data := []*entity.SyncchainChainTransfer{}
	err = result.Structs(&data)
	///
	return data, err
}
func (s *ChainTransfer) UpdateState(ctx context.Context, chainId int64, currentBlock int64) error {
	// _, err := dao.ChainTransfer.Ctx(ctx).Insert(data)
	_, err := s.statemod.Ctx(ctx).Where(dao.SyncchainState.Columns().ChainId, chainId).
		Data(g.Map{
			dao.SyncchainState.Columns().CurrentBlock: currentBlock,
		}).
		// OnConflict(dao.SyncchainState.Columns().ChainId).
		Update()
	return err
}
func (s *ChainTransfer) GetState(ctx context.Context, chainId int64) (*entity.SyncchainState, error) {
	rst, err := s.statemod.Ctx(ctx).
		Where(dao.SyncchainState.Columns().ChainId, chainId).
		One()
	if err != nil {
		return nil, err
	}
	var stat *entity.SyncchainState = nil
	if err := rst.Struct(&stat); err != nil {
		return nil, err
	}
	return stat, nil
}
