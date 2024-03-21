package mpcdao

import (
	"context"

	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

func Insert(ctx context.Context, data *entity.ChainData) error {
	_, err := dao.ChainData.Ctx(ctx).Insert(data)
	return err
}
