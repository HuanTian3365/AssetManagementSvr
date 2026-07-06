package asset

import (
	"asset_management_svr/internal/model/common"
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) FloorList(ctx context.Context, req *v1.FloorListReq) (res *v1.FloorListRes, err error) {
	data, total, err := service.Asset().FloorList(ctx, &req.FloorListReq)
	if err != nil {
		return nil, err
	}
	return &v1.FloorListRes{
		PageResult: common.PageResult{Page: req.Page,
			PageSize: req.PageSize,
			Total:    int64(total)},
		Items: data,
	}, nil
}
