// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetBuilding is the golang structure for table asset_building.
type AssetBuilding struct {
	Id        uint64      `json:"id"        orm:"id"         description:"ID"`   // ID
	Name      string      `json:"name"      orm:"name"       description:"建筑名称"` // 建筑名称
	Code      string      `json:"code"      orm:"code"       description:"建筑编码"` // 建筑编码
	Address   string      `json:"address"   orm:"address"    description:"地址"`   // 地址
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`   // 备注
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"` // 删除时间
}
