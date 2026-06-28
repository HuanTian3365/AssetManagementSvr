// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetFloor is the golang structure for table asset_floor.
type AssetFloor struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`   // ID
	BuildingId uint64      `json:"buildingId" orm:"building_id" description:"建筑ID"` // 建筑ID
	Name       string      `json:"name"       orm:"name"        description:"楼层名称"` // 楼层名称
	Code       string      `json:"code"       orm:"code"        description:"楼层编码"` // 楼层编码
	FloorNo    int         `json:"floorNo"    orm:"floor_no"    description:"楼层序号"` // 楼层序号
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`   // 备注
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:"删除时间"` // 删除时间
}
