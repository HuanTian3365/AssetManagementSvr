// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetFloor is the golang structure of table asset_floor for DAO operations like Where/Data.
type AssetFloor struct {
	g.Meta     `orm:"table:asset_floor, do:true"`
	Id         any         // ID
	BuildingId any         // 建筑ID
	Name       any         // 楼层名称
	Code       any         // 楼层编码
	FloorNo    any         // 楼层序号
	Remark     any         // 备注
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
