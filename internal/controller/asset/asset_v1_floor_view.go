package asset

import (
	"asset_management_svr/internal/service"
	"context"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) FloorView(ctx context.Context, req *v1.FloorViewReq) (res *v1.FloorViewRes, err error) {
	data, err := service.Asset().FloorView(ctx, &req.FloorViewReq)
	if err != nil {
		return nil, err
	}
	return &v1.FloorViewRes{FloorViewRes: data}, nil
}
