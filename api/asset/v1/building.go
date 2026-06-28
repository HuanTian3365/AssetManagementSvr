package v1

import (
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/model/common"

	"github.com/gogf/gf/v2/frame/g"
)

type BuildingCreateReq struct {
	g.Meta `path:"/building/create" method:"post" tags:"资产建筑" summary:"新增建筑" `
	model.BuildingCreateReq
}

type BuildingCreateRes struct {
}

type BuildingUpdateReq struct {
	g.Meta `path:"/building/update" method:"put" tags:"资产建筑" summary:"修改建筑" `
	model.BuildingUpdateReq
}
type BuildingUpdateRes struct {
}

type BuildingDeleteReq struct {
	g.Meta `path:"/building/delete" method:"delete" tags:"资产建筑" summary:"删除建筑" `
	model.BuildingDeleteReq
}
type BuildingDeleteRes struct {
}

type BuildingViewReq struct {
	g.Meta `path:"/building/view" method:"get" tags:"资产建筑" summary:"建筑详情" `
	model.BuildingViewReq
}

type BuildingViewRes struct {
	*model.BuildingViewRes
}

type BuildingListReq struct {
	g.Meta `path:"/building/list" method:"get" tags:"资产建筑" summary:"建筑列表" `
	model.BuildingListReq
}

type BuildingListRes struct {
	common.PageResult
	Items []*model.BuildingListRes `json:"items"   dc:"数据列表"`
}
