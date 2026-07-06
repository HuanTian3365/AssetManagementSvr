package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) RoomUpdate(ctx context.Context, req *v1.RoomUpdateReq) (res *v1.RoomUpdateRes, err error) {
	err = service.Asset().RoomUpdate(ctx, &req.RoomUpdateReq)
	return
}
