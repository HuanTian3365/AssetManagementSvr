// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetProduct is the golang structure of table asset_product for DAO operations like Where/Data.
type AssetProduct struct {
	g.Meta     `orm:"table:asset_product, do:true"`
	Id         any         // ID
	CategoryId any         // 分类ID
	Name       any         // 产品名称
	Code       any         // 产品编码
	Brand      any         // 品牌
	Model      any         // 型号
	Unit       any         // 计量单位
	Remark     any         // 备注
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
