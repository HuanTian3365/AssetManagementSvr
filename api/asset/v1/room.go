package v1

import (
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/model/common"

	"github.com/gogf/gf/v2/frame/g"
)

type RoomCreateReq struct {
	g.Meta `path:"/room/create" method:"post" tags:"资产房间" summary:"新增房间"`
	model.RoomCreateReq
}
type RoomCreateRes struct {
}

type RoomUpdateReq struct {
	g.Meta `path:"/room/update" method:"put" tags:"资产房间" summary:"修改房间"`
	model.RoomUpdateReq
}
type RoomUpdateRes struct{}
type RoomDeleteReq struct {
	g.Meta `path:"/room/delete" method:"delete" tags:"资产房间" summary:"删除房间"`
	model.RoomDeleteReq
}
type RoomDeleteRes struct {
}

type RoomViewReq struct {
	g.Meta `path:"/room/view" method:"get" tags:"资产房间" summary:"房间详情"`
	model.RoomViewReq
}
type RoomViewRes struct {
	*model.RoomViewRes
}

type RoomListReq struct {
	g.Meta `path:"/room/list" method:"get" tags:"资产房间" summary:"房间列表"`
	model.RoomListReq
}
type RoomListRes struct {
	common.PageResult
	Items []*model.RoomListRes `json:"items"   dc:"数据列表"`
}
