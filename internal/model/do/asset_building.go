// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetBuilding is the golang structure of table asset_building for DAO operations like Where/Data.
type AssetBuilding struct {
	g.Meta    `orm:"table:asset_building, do:true"`
	Id        any         // ID
	Name      any         // 建筑名称
	Code      any         // 建筑编码
	Address   any         // 地址
	Remark    any         // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
