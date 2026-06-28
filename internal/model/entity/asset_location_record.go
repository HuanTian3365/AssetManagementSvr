// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetLocationRecord is the golang structure for table asset_location_record.
type AssetLocationRecord struct {
	Id             uint64      `json:"id"             orm:"id"               description:"ID"`            // ID
	AssetId        uint64      `json:"assetId"        orm:"asset_id"         description:"资产ID"`          // 资产ID
	ActionType     int         `json:"actionType"     orm:"action_type"      description:"动作类型: 0入库 1转移"` // 动作类型: 0入库 1转移
	FromBuildingId uint64      `json:"fromBuildingId" orm:"from_building_id" description:"来源建筑ID"`        // 来源建筑ID
	FromFloorId    uint64      `json:"fromFloorId"    orm:"from_floor_id"    description:"来源楼层ID"`        // 来源楼层ID
	FromRoomId     uint64      `json:"fromRoomId"     orm:"from_room_id"     description:"来源房间ID"`        // 来源房间ID
	ToBuildingId   uint64      `json:"toBuildingId"   orm:"to_building_id"   description:"目标建筑ID"`        // 目标建筑ID
	ToFloorId      uint64      `json:"toFloorId"      orm:"to_floor_id"      description:"目标楼层ID"`        // 目标楼层ID
	ToRoomId       uint64      `json:"toRoomId"       orm:"to_room_id"       description:"目标房间ID"`        // 目标房间ID
	OperatedAt     *gtime.Time `json:"operatedAt"     orm:"operated_at"      description:"操作时间"`          // 操作时间
	Remark         string      `json:"remark"         orm:"remark"           description:"备注"`            // 备注
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`          // 创建时间
}
