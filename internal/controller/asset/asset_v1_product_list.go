package asset

import (
	"asset_management_svr/internal/model/common"
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) ProductList(ctx context.Context, req *v1.ProductListReq) (res *v1.ProductListRes, err error) {
	data, total, err := service.Asset().ProductList(ctx, &req.ProductListReq)
	if err != nil {
		return nil, err
	}
	return &v1.ProductListRes{
		PageResult: common.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    int64(total),
		},
		Items: data,
	}, nil
}
