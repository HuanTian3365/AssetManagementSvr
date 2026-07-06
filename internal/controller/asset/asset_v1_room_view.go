package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) RoomView(ctx context.Context, req *v1.RoomViewReq) (res *v1.RoomViewRes, err error) {
	data, err := service.Asset().RoomView(ctx, &req.RoomViewReq)
	if err != nil {
		return nil, err
	}
	return &v1.RoomViewRes{
		RoomViewRes: data,
	}, nil
}
