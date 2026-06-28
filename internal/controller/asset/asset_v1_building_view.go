package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) BuildingView(ctx context.Context, req *v1.BuildingViewReq) (res *v1.BuildingViewRes, err error) {
	data, err := service.Asset().BuildingView(ctx, &req.BuildingViewReq)
	if err != nil {
		return
	}
	return &v1.BuildingViewRes{
		BuildingViewRes: data,
	}, nil
}
