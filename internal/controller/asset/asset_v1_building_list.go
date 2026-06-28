package asset

import (
	"asset_management_svr/internal/model/common"
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) BuildingList(ctx context.Context, req *v1.BuildingListReq) (res *v1.BuildingListRes, err error) {
	data, total, err := service.Asset().BuildingList(ctx, &req.BuildingListReq)
	if err != nil {
		return nil, err
	}
	return &v1.BuildingListRes{
		PageResult: common.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    int64(total),
		},
		Items: data,
	}, nil
}
