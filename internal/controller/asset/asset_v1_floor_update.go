package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) FloorUpdate(ctx context.Context, req *v1.FloorUpdateReq) (res *v1.FloorUpdateRes, err error) {
	err = service.Asset().FloorUpdate(ctx, &req.FloorUpdateReq)
	return
}
