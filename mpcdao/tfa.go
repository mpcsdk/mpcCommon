package mpcdao

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/do"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type RiskTfa struct {
	redis *gredis.Redis
	dur   time.Duration
}

func NewRiskTfa(redis *gredis.Redis, dur int) *RiskTfa {

	return &RiskTfa{
		redis: redis,
		dur:   time.Duration(dur) * time.Second,
	}
}
func (s *RiskTfa) FetchTfaInfo(ctx context.Context, userId string) (*entity.Tfa, error) {
	if userId == "" {
		return nil, mpccode.CodeParamInvalid()
	}

	aggdo := &do.Tfa{
		UserId: userId,
	}
	var data *entity.Tfa
	///
	rst, err := dao.Tfa.Ctx(ctx).Where(aggdo).One()
	if err != nil {
		g.Log().Error(ctx, "ExistsTfaInfo:", "userId", userId, "agg:", aggdo, "err", err)
		return nil, mpccode.CodeInternalError()
	}
	if rst.IsEmpty() {
		return nil, nil
	}
	err = rst.Struct(&data)
	if err != nil {
		g.Log().Error(ctx, "ExistsTfaInfo:", "userId", userId, "agg:", aggdo, "err", err)
		return nil, mpccode.CodeInternalError()
	}
	///
	// if data.MailUpdatedAt != nil {
	// 	g.Log().Debug(ctx, "##time1:", data.MailUpdatedAt.String(), data.MailUpdatedAt.Timestamp(), data.MailUpdatedAt.UTC().Timestamp())
	// 	// data.MailUpdatedAt = gtime.New(data.MailUpdatedAt.String())
	// 	g.Log().Debug(ctx, "##time11:", data.MailUpdatedAt.String(), data.MailUpdatedAt.Timestamp(), data.MailUpdatedAt.UTC().Timestamp())
	// }
	return data, nil
}
func (s *RiskTfa) TfaMailNotExists(ctx context.Context, mail string) (bool, error) {
	rst, err := dao.Tfa.Ctx(ctx).Where(do.Tfa{
		Mail: mail,
	}).Count()
	if err != nil {
		g.Log().Error(ctx, "TfaMailNotExists:", "mail", mail, "err", err)
		return false, mpccode.CodeInternalError()
	}
	if rst > 0 {
		return false, nil
	}
	return true, nil
}
func (s *RiskTfa) TfaPhoneNotExists(ctx context.Context, phone string) (bool, error) {
	rst, err := dao.Tfa.Ctx(ctx).Where(do.Tfa{
		Phone: phone,
	}).CountColumn(dao.Tfa.Columns().Phone)
	if err != nil {
		g.Log().Error(ctx, "TfaPhoneNotExists:", "phone", phone, "err", err)
		return false, mpccode.CodeInternalError()
	}
	if rst > 0 {
		return false, nil
	}
	return true, nil
}
func (s *RiskTfa) InsertTfaInfo(ctx context.Context, userId string, data *do.Tfa) error {
	cnt, err := dao.Tfa.Ctx(ctx).Where(do.Tfa{
		UserId: data.UserId,
	}).CountColumn(dao.Tfa.Columns().UserId)

	if err != nil {
		g.Log().Error(ctx, "InsertTfaInfo:", "userId", userId, "data:", data, "err", err)
		return mpccode.CodeInternalError()
	}
	if cnt != 0 {
		return nil
	}

	_, err = dao.Tfa.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.Tfa.Table() + userId,
		Force:    false,
	}).Data(data).Insert()
	if err != nil {
		g.Log().Error(ctx, "InsertTfaInfo:", "userId", userId, "data:", data, "err", err)
		return mpccode.CodeInternalError()
	}

	return nil
}
func (s *RiskTfa) UpdateTfaInfo(ctx context.Context, userId string, data *do.Tfa) error {
	///todo:
	// if data != nil && data.MailUpdatedAt != nil {
	// 	g.Log().Debug(ctx, "##timeup:", data.MailUpdatedAt.String(), data.MailUpdatedAt.Timestamp(), data.MailUpdatedAt.UTC().Timestamp())
	// }

	_, err := dao.Tfa.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.Tfa.Table() + userId,
		Force:    false,
	}).Data(data).Where(do.Tfa{
		UserId: data.UserId,
	}).Update()
	if err != nil {
		g.Log().Error(ctx, "UpdateTfaInfo:", "userId", userId, "data:", data, "err", err)
		return mpccode.CodeInternalError()
	}
	return nil
}
func (s *RiskTfa) ExistsTfaInfo(ctx context.Context, userId string) (bool, error) {
	if userId == "" {
		return false, mpccode.CodeParamInvalid()
	}
	cnt, err := dao.Tfa.Ctx(ctx).Where(do.Tfa{
		UserId: userId,
	}).CountColumn(dao.Tfa.Columns().UserId)

	if err != nil {
		g.Log().Error(ctx, "ExistsTfaInfo:", "userId", userId, "err", err)
		return false, mpccode.CodeInternalError()
	}
	if cnt != 0 {
		return true, nil
	}
	return false, nil
}
