package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) ProductView(ctx context.Context, req *v1.ProductViewReq) (res *v1.ProductViewRes, err error) {
	data, err := service.Asset().ProductView(ctx, &req.ProductViewReq)
	if err != nil {
		return nil, err
	}
	return &v1.ProductViewRes{
		ProductViewRes: data,
	}, nil
}
