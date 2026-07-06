package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) RoomDelete(ctx context.Context, req *v1.RoomDeleteReq) (res *v1.RoomDeleteRes, err error) {
	err = service.Asset().RoomDelete(ctx, &req.RoomDeleteReq)
	return
}
