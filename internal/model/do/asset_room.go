// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetRoom is the golang structure of table asset_room for DAO operations like Where/Data.
type AssetRoom struct {
	g.Meta     `orm:"table:asset_room, do:true"`
	Id         any         // ID
	BuildingId any         // 建筑ID
	FloorId    any         // 楼层ID
	Name       any         // 房间名称
	Code       any         // 房间编码
	RoomNo     any         // 房间号
	Remark     any         // 备注
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
