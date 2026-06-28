// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetLocationRecord is the golang structure of table asset_location_record for DAO operations like Where/Data.
type AssetLocationRecord struct {
	g.Meta         `orm:"table:asset_location_record, do:true"`
	Id             any         // ID
	AssetId        any         // 资产ID
	ActionType     any         // 动作类型: 0入库 1转移
	FromBuildingId any         // 来源建筑ID
	FromFloorId    any         // 来源楼层ID
	FromRoomId     any         // 来源房间ID
	ToBuildingId   any         // 目标建筑ID
	ToFloorId      any         // 目标楼层ID
	ToRoomId       any         // 目标房间ID
	OperatedAt     *gtime.Time // 操作时间
	Remark         any         // 备注
	CreatedAt      *gtime.Time // 创建时间
}
