// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetCategory is the golang structure for table asset_category.
type AssetCategory struct {
	Id        uint64      `json:"id"        orm:"id"         description:"ID"`    // ID
	ParentId  uint64      `json:"parentId"  orm:"parent_id"  description:"父分类ID"` // 父分类ID
	Name      string      `json:"name"      orm:"name"       description:"分类名称"`  // 分类名称
	Code      string      `json:"code"      orm:"code"       description:"分类编码"`  // 分类编码
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`    // 排序
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`    // 备注
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`  // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`  // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`  // 删除时间
}
