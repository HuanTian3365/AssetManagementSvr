package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) CategoryUpdate(ctx context.Context, req *v1.CategoryUpdateReq) (res *v1.CategoryUpdateRes, err error) {
	err = service.Asset().CategoryUpdate(ctx, &req.CategoryUpdateReq)
	return
}
