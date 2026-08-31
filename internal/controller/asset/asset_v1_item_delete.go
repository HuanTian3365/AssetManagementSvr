package asset

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"asset_management_svr/api/asset/v1"
)

func (c *ControllerV1) ItemDelete(ctx context.Context, req *v1.ItemDeleteReq) (res *v1.ItemDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
