package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) RoomCreate(ctx context.Context, req *v1.RoomCreateReq) (res *v1.RoomCreateRes, err error) {
	err = service.Asset().RoomCreate(ctx, &req.RoomCreateReq)
	return
}
