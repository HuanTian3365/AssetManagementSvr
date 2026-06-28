// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AssetProduct is the golang structure for table asset_product.
type AssetProduct struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`   // ID
	CategoryId uint64      `json:"categoryId" orm:"category_id" description:"分类ID"` // 分类ID
	Name       string      `json:"name"       orm:"name"        description:"产品名称"` // 产品名称
	Code       string      `json:"code"       orm:"code"        description:"产品编码"` // 产品编码
	Brand      string      `json:"brand"      orm:"brand"       description:"品牌"`   // 品牌
	Model      string      `json:"model"      orm:"model"       description:"型号"`   // 型号
	Unit       string      `json:"unit"       orm:"unit"        description:"计量单位"` // 计量单位
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`   // 备注
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:"删除时间"` // 删除时间
}
