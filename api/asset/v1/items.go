package v1

import (
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/model/common"

	"github.com/gogf/gf/v2/frame/g"
)

type ItemCreateReq struct {
	g.Meta `path:"/item/create" method:"post" tags:"物品管理" summary:"新增物品"`
	model.ItemCreateReq
}

type ItemCreateRes struct{}

type ItemUpdateReq struct {
	g.Meta `path:"/item/update" method:"put" tags:"物品管理" summary:"修改物品"`
	model.ItemUpdateReq
}
type ItemUpdateRes struct{}
type ItemDeleteReq struct {
	g.Meta `path:"/item/delete" method:"delete" tags:"物品管理" summary:"删除物品"`
	model.ItemDeleteReq
}
type ItemDeleteRes struct{}
type ItemViewReq struct {
	g.Meta `path:"/item/view" method:"get" tags:"物品管理" summary:"物品详情"`
	model.ItemViewReq
}
type ItemViewRes struct {
	*model.ItemViewRes
}
type ItemListReq struct {
	g.Meta `path:"/item/list" method:"get" tags:"物品管理" summary:"物品列表"`
	model.ItemListReq
}
type ItemListRes struct {
	common.PageResult
	Items []*model.FloorListRes `json:"items" dc:"数据列表"`
}
