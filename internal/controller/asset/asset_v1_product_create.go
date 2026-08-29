package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) ProductCreate(ctx context.Context, req *v1.ProductCreateReq) (res *v1.ProductCreateRes, err error) {
	return nil, service.Asset().ProductCreate(ctx, &req.ProductCreateReq)
}
