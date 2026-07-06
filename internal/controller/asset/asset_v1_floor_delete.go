package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) FloorDelete(ctx context.Context, req *v1.FloorDeleteReq) (res *v1.FloorDeleteRes, err error) {
	err = service.Asset().FloorDelete(ctx, &req.FloorDeleteReq)
	return
}
