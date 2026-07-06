package asset

import (
	"asset_management_svr/internal/model/common"
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) RoomList(ctx context.Context, req *v1.RoomListReq) (res *v1.RoomListRes, err error) {
	data, total, err := service.Asset().RoomList(ctx, &req.RoomListReq)
	if err != nil {
		return nil, err
	}
	return &v1.RoomListRes{
		PageResult: common.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    int64(total),
		},
		Items: data,
	}, nil
}
