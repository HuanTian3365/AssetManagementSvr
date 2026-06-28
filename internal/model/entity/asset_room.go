// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetRoom is the golang structure for table asset_room.
type AssetRoom struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`   // ID
	BuildingId uint64      `json:"buildingId" orm:"building_id" description:"建筑ID"` // 建筑ID
	FloorId    uint64      `json:"floorId"    orm:"floor_id"    description:"楼层ID"` // 楼层ID
	Name       string      `json:"name"       orm:"name"        description:"房间名称"` // 房间名称
	Code       string      `json:"code"       orm:"code"        description:"房间编码"` // 房间编码
	RoomNo     string      `json:"roomNo"     orm:"room_no"     description:"房间号"`  // 房间号
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`   // 备注
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:"删除时间"` // 删除时间
}
