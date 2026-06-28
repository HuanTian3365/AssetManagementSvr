// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetCategory is the golang structure of table asset_category for DAO operations like Where/Data.
type AssetCategory struct {
	g.Meta    `orm:"table:asset_category, do:true"`
	Id        any         // ID
	ParentId  any         // 父分类ID
	Name      any         // 分类名称
	Code      any         // 分类编码
	Sort      any         // 排序
	Remark    any         // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
