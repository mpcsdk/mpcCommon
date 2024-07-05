package mpcdao

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type ChainTransfer struct {
	chainId int64
	dbname  string
	dbmod   *gdb.Model
	redis   *gredis.Redis
	dur     time.Duration
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

func CreateChainTransferDB(ctx context.Context, chainId int64) error {
	dbname := "sync_chain_" + gconv.String(chainId)
	_, err := dao.ChainTransfer.DB().Exec(ctx, "CREATE DATABASE "+dbname)
	if err != nil {
		return err
	}
	_, err = dao.ChainTransfer.DB().Schema(dbname).Exec(ctx, `CREATE TABLE "public"."chain_transfer" (
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
	  );`)
	return err
}
func NewChainTransfer(chainId int64, redis *gredis.Redis, dur int) *ChainTransfer {
	dbname := "sync_chain_" + gconv.String(chainId)

	dbmod := dao.ChainTransfer.DB().Schema(dbname).Model(dao.ChainTransfer.Table()).Safe()
	if redis != nil {
		g.DB(dao.ChainTransfer.Group()).GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
	}

	return &ChainTransfer{
		dbname:  dbname,
		dbmod:   dbmod,
		chainId: chainId,
		redis:   redis,
		dur:     time.Duration(dur) * time.Second,
	}
}

func (s *ChainTransfer) Insert(ctx context.Context, data *entity.ChainTransfer) error {
	// _, err := dao.ChainTransfer.Ctx(ctx).Insert(data)
	_, err := s.dbmod.Ctx(ctx).Insert(data)
	return err
}
func (s *ChainTransfer) DelChainBlockNumber(ctx context.Context, chainId int64, number int64) error {
	_, err := s.dbmod.Ctx(ctx).
		Where(dao.ChainTransfer.Columns().ChainId, chainId).
		Where(dao.ChainTransfer.Columns().Height, number).
		Delete()

	return err
}
func (s *ChainTransfer) InsertBatch(ctx context.Context, data []*entity.ChainTransfer) error {
	_, err := s.dbmod.Ctx(ctx).Insert(data)
	return err
}

func (s *ChainTransfer) Query(ctx context.Context, query *QueryData) ([]*entity.ChainTransfer, error) {
	if query.PageSize < 0 || query.Page < 0 {
		return nil, nil
	}
	//
	where := s.dbmod.Ctx(ctx)
	if query.ChainId != 0 {
		where = where.Where(dao.ChainTransfer.Columns().ChainId, query.ChainId)
	}
	if len(query.Kinds) > 0 {
		where = where.Where(dao.ChainTransfer.Columns().Kind, query.Kinds)
	}
	if query.From != "" {
		where = where.Where(dao.ChainTransfer.Columns().From, query.From)
	}
	if query.To != "" {
		where = where.Where(dao.ChainTransfer.Columns().To, query.To)
	}
	if query.Contract != "" {
		where = where.Where(dao.ChainTransfer.Columns().Contract, query.Contract)
	}
	///time
	if query.StartTime != 0 {
		where = where.WhereGTE(dao.ChainTransfer.Columns().Ts, query.StartTime)
	}
	if query.EndTime != 0 {
		where = where.WhereLTE(dao.ChainTransfer.Columns().Ts, query.EndTime)
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
	data := []*entity.ChainTransfer{}
	err = result.Structs(&data)
	///
	return data, err
}
