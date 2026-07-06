package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) FloorCreate(ctx context.Context, req *v1.FloorCreateReq) (res *v1.FloorCreateRes, err error) {
	err = service.Asset().FloorCreate(ctx, &req.FloorCreateReq)
	return
}
