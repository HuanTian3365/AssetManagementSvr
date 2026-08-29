package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) ProductUpdate(ctx context.Context, req *v1.ProductUpdateReq) (res *v1.ProductUpdateRes, err error) {
	return nil, service.Asset().ProductUpdate(ctx, &req.ProductUpdateReq)
}
