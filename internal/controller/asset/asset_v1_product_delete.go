package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) ProductDelete(ctx context.Context, req *v1.ProductDeleteReq) (res *v1.ProductDeleteRes, err error) {
	return nil, service.Asset().ProductDelete(ctx, &req.ProductDeleteReq)
}
