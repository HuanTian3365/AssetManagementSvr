// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetItem is the golang structure for table asset_item.
type AssetItem struct {
	Id                uint64      `json:"id"                orm:"id"                  description:"ID"`                   // ID
	ProductId         uint64      `json:"productId"         orm:"product_id"          description:"产品ID"`                 // 产品ID
	AssetCode         string      `json:"assetCode"         orm:"asset_code"          description:"资产编码"`                 // 资产编码
	Name              string      `json:"name"              orm:"name"                description:"资产名称"`                 // 资产名称
	Status            int         `json:"status"            orm:"status"              description:"资产状态：1闲置 2在用 3维修 4报废"` // 资产状态：1闲置 2在用 3维修 4报废
	CurrentBuildingId uint64      `json:"currentBuildingId" orm:"current_building_id" description:"建筑ID"`                 // 建筑ID
	CurrentFloorId    uint64      `json:"currentFloorId"    orm:"current_floor_id"    description:"楼层ID"`                 // 楼层ID
	CurrentRoomId     uint64      `json:"currentRoomId"     orm:"current_room_id"     description:"房间ID"`                 // 房间ID
	PurchaseDate      *gtime.Time `json:"purchaseDate"      orm:"purchase_date"       description:"购买日期"`                 // 购买日期
	PurchasePrice     float64     `json:"purchasePrice"     orm:"purchase_price"      description:"购买价格"`                 // 购买价格
	Remark            string      `json:"remark"            orm:"remark"              description:"备注"`                   // 备注
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          description:"创建时间"`                 // 创建时间
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          description:"更新时间"`                 // 更新时间
	DeletedAt         *gtime.Time `json:"deletedAt"         orm:"deleted_at"          description:"删除时间"`                 // 删除时间
}
