// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetItem is the golang structure of table asset_item for DAO operations like Where/Data.
type AssetItem struct {
	g.Meta            `orm:"table:asset_item, do:true"`
	Id                any         // ID
	ProductId         any         // 产品ID
	AssetCode         any         // 资产编码
	Name              any         // 资产名称
	Status            any         // 资产状态：1闲置 2在用 3维修 4报废
	CurrentBuildingId any         // 建筑ID
	CurrentFloorId    any         // 楼层ID
	CurrentRoomId     any         // 房间ID
	PurchaseDate      *gtime.Time // 购买日期
	PurchasePrice     any         // 购买价格
	Remark            any         // 备注
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
	DeletedAt         *gtime.Time // 删除时间
}
