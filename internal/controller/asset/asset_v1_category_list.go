package asset

import (
	"asset_management_svr/internal/model/common"
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) CategoryList(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error) {
	data, total, err := service.Asset().CategoryList(ctx, &req.CategoryListReq)
	if err != nil {
		return nil, err
	}
	return &v1.CategoryListRes{
		PageResult: common.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    int64(total),
		},
		Items: data,
	}, nil
}
