package v1

import (
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/model/common"

	"github.com/gogf/gf/v2/frame/g"
)

type FloorCreateReq struct {
	g.Meta `path:"/floor/create" method:"post" tags:"资产楼层" summary:"新增楼层"`
	model.FloorCreateReq
}
type FloorCreateRes struct{}

type FloorUpdateReq struct {
	g.Meta `path:"/floor/update" method:"put" tags:"资产楼层" summary:"修改楼层"`
	model.FloorUpdateReq
}
type FloorUpdateRes struct{}

type FloorDeleteReq struct {
	g.Meta `path:"/floor/delete" method:"delete" tags:"资产楼层" summary:"删除楼层"`
	model.FloorDeleteReq
}
type FloorDeleteRes struct{}

type FloorViewReq struct {
	g.Meta `path:"/floor/view" method:"get" tags:"资产楼层" summary:"楼层详情"`
	model.FloorViewReq
}
type FloorViewRes struct {
	*model.FloorViewRes
}
type FloorListReq struct {
	g.Meta `path:"/floor/list" method:"get" tags:"资产楼层" summary:"楼层列表"`
	model.FloorListReq
}
type FloorListRes struct {
	common.PageResult
	Items []*model.FloorListRes `json:"items" dc:"数据列表"`
}
