package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) CategoryView(ctx context.Context, req *v1.CategoryViewReq) (res *v1.CategoryViewRes, err error) {
	data, err := service.Asset().CategoryView(ctx, &req.CategoryViewReq)
	if err != nil {
		return nil, err
	}
	return &v1.CategoryViewRes{
		CategoryViewRes: data,
	}, nil
}
