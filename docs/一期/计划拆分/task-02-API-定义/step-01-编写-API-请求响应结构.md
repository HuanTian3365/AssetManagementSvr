# Task 2 / Step 1: 编写 API 请求响应结构

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [API 定义](README.md)

- [ ] **Step 1: 编写 API 请求响应结构**

写入 `api/asset/v1/asset.go`：

```go
package v1

import "github.com/gogf/gf/v2/frame/g"

type PageReq struct {
	Page     int `json:"page" v:"min:1#页码不能小于1" dc:"页码"`
	PageSize int `json:"pageSize" v:"max:100#每页数量不能超过100" dc:"每页数量"`
}

type IdReq struct {
	Id uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
}

type BuildingCreateReq struct {
	g.Meta  `path:"/asset/building/create" method:"post" tags:"资产建筑" summary:"创建建筑"`
	Name    string `json:"name" v:"required#建筑名称不能为空" dc:"建筑名称"`
	Code    string `json:"code" v:"required#建筑编码不能为空" dc:"建筑编码"`
	Address string `json:"address" dc:"地址"`
	Remark  string `json:"remark" dc:"备注"`
}
type BuildingCreateRes struct{ Id uint64 `json:"id"` }

type BuildingUpdateReq struct {
	g.Meta  `path:"/asset/building/update" method:"put" tags:"资产建筑" summary:"更新建筑"`
	Id      uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	Name    string `json:"name" v:"required#建筑名称不能为空" dc:"建筑名称"`
	Code    string `json:"code" v:"required#建筑编码不能为空" dc:"建筑编码"`
	Address string `json:"address" dc:"地址"`
	Remark  string `json:"remark" dc:"备注"`
}
type BuildingUpdateRes struct{}

type BuildingDeleteReq struct {
	g.Meta `path:"/asset/building/delete" method:"delete" tags:"资产建筑" summary:"删除建筑"`
	IdReq
}
type BuildingDeleteRes struct{}

type BuildingDetailReq struct {
	g.Meta `path:"/asset/building/detail" method:"get" tags:"资产建筑" summary:"建筑详情"`
	IdReq
}
type BuildingDetailRes struct{ Data any `json:"data"` }

type BuildingListReq struct {
	g.Meta `path:"/asset/building/list" method:"get" tags:"资产建筑" summary:"建筑列表"`
	PageReq
	Name string `json:"name" dc:"建筑名称"`
	Code string `json:"code" dc:"建筑编码"`
}
type BuildingListRes struct {
	List  any `json:"list"`
	Total int `json:"total"`
}

type FloorCreateReq struct {
	g.Meta     `path:"/asset/floor/create" method:"post" tags:"资产楼层" summary:"创建楼层"`
	BuildingId uint64 `json:"buildingId" v:"required|min:1#建筑不能为空|建筑不正确" dc:"建筑ID"`
	Name       string `json:"name" v:"required#楼层名称不能为空" dc:"楼层名称"`
	Code       string `json:"code" v:"required#楼层编码不能为空" dc:"楼层编码"`
	FloorNo    int    `json:"floorNo" dc:"楼层序号"`
	Remark     string `json:"remark" dc:"备注"`
}
type FloorCreateRes struct{ Id uint64 `json:"id"` }

type FloorUpdateReq struct {
	g.Meta     `path:"/asset/floor/update" method:"put" tags:"资产楼层" summary:"更新楼层"`
	Id         uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	BuildingId uint64 `json:"buildingId" v:"required|min:1#建筑不能为空|建筑不正确" dc:"建筑ID"`
	Name       string `json:"name" v:"required#楼层名称不能为空" dc:"楼层名称"`
	Code       string `json:"code" v:"required#楼层编码不能为空" dc:"楼层编码"`
	FloorNo    int    `json:"floorNo" dc:"楼层序号"`
	Remark     string `json:"remark" dc:"备注"`
}
type FloorUpdateRes struct{}

type FloorDeleteReq struct {
	g.Meta `path:"/asset/floor/delete" method:"delete" tags:"资产楼层" summary:"删除楼层"`
	IdReq
}
type FloorDeleteRes struct{}

type FloorDetailReq struct {
	g.Meta `path:"/asset/floor/detail" method:"get" tags:"资产楼层" summary:"楼层详情"`
	IdReq
}
type FloorDetailRes struct{ Data any `json:"data"` }

type FloorListReq struct {
	g.Meta     `path:"/asset/floor/list" method:"get" tags:"资产楼层" summary:"楼层列表"`
	PageReq
	BuildingId uint64 `json:"buildingId" dc:"建筑ID"`
	Name       string `json:"name" dc:"楼层名称"`
	Code       string `json:"code" dc:"楼层编码"`
}
type FloorListRes struct {
	List  any `json:"list"`
	Total int `json:"total"`
}

type RoomCreateReq struct {
	g.Meta     `path:"/asset/room/create" method:"post" tags:"资产房间" summary:"创建房间"`
	BuildingId uint64 `json:"buildingId" v:"required|min:1#建筑不能为空|建筑不正确" dc:"建筑ID"`
	FloorId    uint64 `json:"floorId" v:"required|min:1#楼层不能为空|楼层不正确" dc:"楼层ID"`
	Name       string `json:"name" v:"required#房间名称不能为空" dc:"房间名称"`
	Code       string `json:"code" v:"required#房间编码不能为空" dc:"房间编码"`
	RoomNo     string `json:"roomNo" dc:"房间号"`
	Remark     string `json:"remark" dc:"备注"`
}
type RoomCreateRes struct{ Id uint64 `json:"id"` }

type RoomUpdateReq struct {
	g.Meta     `path:"/asset/room/update" method:"put" tags:"资产房间" summary:"更新房间"`
	Id         uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	BuildingId uint64 `json:"buildingId" v:"required|min:1#建筑不能为空|建筑不正确" dc:"建筑ID"`
	FloorId    uint64 `json:"floorId" v:"required|min:1#楼层不能为空|楼层不正确" dc:"楼层ID"`
	Name       string `json:"name" v:"required#房间名称不能为空" dc:"房间名称"`
	Code       string `json:"code" v:"required#房间编码不能为空" dc:"房间编码"`
	RoomNo     string `json:"roomNo" dc:"房间号"`
	Remark     string `json:"remark" dc:"备注"`
}
type RoomUpdateRes struct{}

type RoomDeleteReq struct {
	g.Meta `path:"/asset/room/delete" method:"delete" tags:"资产房间" summary:"删除房间"`
	IdReq
}
type RoomDeleteRes struct{}

type RoomDetailReq struct {
	g.Meta `path:"/asset/room/detail" method:"get" tags:"资产房间" summary:"房间详情"`
	IdReq
}
type RoomDetailRes struct{ Data any `json:"data"` }

type RoomListReq struct {
	g.Meta     `path:"/asset/room/list" method:"get" tags:"资产房间" summary:"房间列表"`
	PageReq
	BuildingId uint64 `json:"buildingId" dc:"建筑ID"`
	FloorId    uint64 `json:"floorId" dc:"楼层ID"`
	Name       string `json:"name" dc:"房间名称"`
	Code       string `json:"code" dc:"房间编码"`
}
type RoomListRes struct {
	List  any `json:"list"`
	Total int `json:"total"`
}

type CategoryCreateReq struct {
	g.Meta  `path:"/asset/category/create" method:"post" tags:"资产分类" summary:"创建分类"`
	ParentId uint64 `json:"parentId" dc:"父级ID"`
	Name     string `json:"name" v:"required#分类名称不能为空" dc:"分类名称"`
	Code     string `json:"code" v:"required#分类编码不能为空" dc:"分类编码"`
	Sort     int    `json:"sort" dc:"排序"`
	Remark   string `json:"remark" dc:"备注"`
}
type CategoryCreateRes struct{ Id uint64 `json:"id"` }

type CategoryUpdateReq struct {
	g.Meta  `path:"/asset/category/update" method:"put" tags:"资产分类" summary:"更新分类"`
	Id       uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	ParentId uint64 `json:"parentId" dc:"父级ID"`
	Name     string `json:"name" v:"required#分类名称不能为空" dc:"分类名称"`
	Code     string `json:"code" v:"required#分类编码不能为空" dc:"分类编码"`
	Sort     int    `json:"sort" dc:"排序"`
	Remark   string `json:"remark" dc:"备注"`
}
type CategoryUpdateRes struct{}

type CategoryDeleteReq struct {
	g.Meta `path:"/asset/category/delete" method:"delete" tags:"资产分类" summary:"删除分类"`
	IdReq
}
type CategoryDeleteRes struct{}

type CategoryDetailReq struct {
	g.Meta `path:"/asset/category/detail" method:"get" tags:"资产分类" summary:"分类详情"`
	IdReq
}
type CategoryDetailRes struct{ Data any `json:"data"` }

type CategoryListReq struct {
	g.Meta `path:"/asset/category/list" method:"get" tags:"资产分类" summary:"分类列表"`
	PageReq
	ParentId uint64 `json:"parentId" dc:"父级ID"`
	Name     string `json:"name" dc:"分类名称"`
	Code     string `json:"code" dc:"分类编码"`
}
type CategoryListRes struct {
	List  any `json:"list"`
	Total int `json:"total"`
}

type ProductCreateReq struct {
	g.Meta    `path:"/asset/product/create" method:"post" tags:"资产产品" summary:"创建产品"`
	CategoryId uint64 `json:"categoryId" v:"required|min:1#分类不能为空|分类不正确" dc:"分类ID"`
	Name       string `json:"name" v:"required#产品名称不能为空" dc:"产品名称"`
	Code       string `json:"code" v:"required#产品编码不能为空" dc:"产品编码"`
	Brand      string `json:"brand" dc:"品牌"`
	Model      string `json:"model" dc:"型号"`
	Unit       string `json:"unit" dc:"单位"`
	Remark     string `json:"remark" dc:"备注"`
}
type ProductCreateRes struct{ Id uint64 `json:"id"` }

type ProductUpdateReq struct {
	g.Meta    `path:"/asset/product/update" method:"put" tags:"资产产品" summary:"更新产品"`
	Id         uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	CategoryId uint64 `json:"categoryId" v:"required|min:1#分类不能为空|分类不正确" dc:"分类ID"`
	Name       string `json:"name" v:"required#产品名称不能为空" dc:"产品名称"`
	Code       string `json:"code" v:"required#产品编码不能为空" dc:"产品编码"`
	Brand      string `json:"brand" dc:"品牌"`
	Model      string `json:"model" dc:"型号"`
	Unit       string `json:"unit" dc:"单位"`
	Remark     string `json:"remark" dc:"备注"`
}
type ProductUpdateRes struct{}

type ProductDeleteReq struct {
	g.Meta `path:"/asset/product/delete" method:"delete" tags:"资产产品" summary:"删除产品"`
	IdReq
}
type ProductDeleteRes struct{}

type ProductDetailReq struct {
	g.Meta `path:"/asset/product/detail" method:"get" tags:"资产产品" summary:"产品详情"`
	IdReq
}
type ProductDetailRes struct{ Data any `json:"data"` }

type ProductListReq struct {
	g.Meta    `path:"/asset/product/list" method:"get" tags:"资产产品" summary:"产品列表"`
	PageReq
	CategoryId uint64 `json:"categoryId" dc:"分类ID"`
	Name       string `json:"name" dc:"产品名称"`
	Code       string `json:"code" dc:"产品编码"`
}
type ProductListRes struct {
	List  any `json:"list"`
	Total int `json:"total"`
}

type ItemCreateReq struct {
	g.Meta      `path:"/asset/item/create" method:"post" tags:"资产实例" summary:"创建资产"`
	ProductId   uint64  `json:"productId" v:"required|min:1#产品不能为空|产品不正确" dc:"产品ID"`
	AssetCode   string  `json:"assetCode" dc:"资产编码"`
	Name        string  `json:"name" v:"required#资产名称不能为空" dc:"资产名称"`
	Status      int     `json:"status" v:"required|in:1,2,3,4#状态不能为空|状态不正确" dc:"状态"`
	RoomId      uint64  `json:"roomId" dc:"房间ID"`
	PurchaseDate string `json:"purchaseDate" dc:"购买日期"`
	PurchasePrice float64 `json:"purchasePrice" dc:"购买价格"`
	Remark      string  `json:"remark" dc:"备注"`
}
type ItemCreateRes struct {
	Id        uint64 `json:"id"`
	AssetCode string `json:"assetCode"`
}

type ItemUpdateReq struct {
	g.Meta      `path:"/asset/item/update" method:"put" tags:"资产实例" summary:"更新资产"`
	Id          uint64  `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	ProductId   uint64  `json:"productId" v:"required|min:1#产品不能为空|产品不正确" dc:"产品ID"`
	AssetCode   string  `json:"assetCode" v:"required#资产编码不能为空" dc:"资产编码"`
	Name        string  `json:"name" v:"required#资产名称不能为空" dc:"资产名称"`
	Status      int     `json:"status" v:"required|in:1,2,3,4#状态不能为空|状态不正确" dc:"状态"`
	PurchaseDate string `json:"purchaseDate" dc:"购买日期"`
	PurchasePrice float64 `json:"purchasePrice" dc:"购买价格"`
	Remark      string  `json:"remark" dc:"备注"`
}
type ItemUpdateRes struct{}

type ItemDeleteReq struct {
	g.Meta `path:"/asset/item/delete" method:"delete" tags:"资产实例" summary:"删除资产"`
	IdReq
}
type ItemDeleteRes struct{}

type ItemDetailReq struct {
	g.Meta `path:"/asset/item/detail" method:"get" tags:"资产实例" summary:"资产详情"`
	IdReq
}
type ItemDetailRes struct{ Data any `json:"data"` }

type ItemListReq struct {
	g.Meta    `path:"/asset/item/list" method:"get" tags:"资产实例" summary:"资产列表"`
	PageReq
	ProductId uint64 `json:"productId" dc:"产品ID"`
	RoomId    uint64 `json:"roomId" dc:"房间ID"`
	Name      string `json:"name" dc:"资产名称"`
	AssetCode string `json:"assetCode" dc:"资产编码"`
	Status    int    `json:"status" dc:"状态"`
}
type ItemListRes struct {
	List  any `json:"list"`
	Total int `json:"total"`
}

type ItemInboundReq struct {
	g.Meta `path:"/asset/item/inbound" method:"post" tags:"资产实例" summary:"资产入库"`
	AssetId uint64 `json:"assetId" v:"required|min:1#资产不能为空|资产不正确" dc:"资产ID"`
	RoomId  uint64 `json:"roomId" v:"required|min:1#房间不能为空|房间不正确" dc:"房间ID"`
	Remark  string `json:"remark" dc:"备注"`
}
type ItemInboundRes struct{}

type ItemTransferReq struct {
	g.Meta `path:"/asset/item/transfer" method:"post" tags:"资产实例" summary:"资产转移"`
	AssetId uint64 `json:"assetId" v:"required|min:1#资产不能为空|资产不正确" dc:"资产ID"`
	RoomId  uint64 `json:"roomId" v:"required|min:1#房间不能为空|房间不正确" dc:"目标房间ID"`
	Remark  string `json:"remark" dc:"备注"`
}
type ItemTransferRes struct{}

type LocationRecordListReq struct {
	g.Meta `path:"/asset/item/location-record/list" method:"get" tags:"资产实例" summary:"资产位置记录"`
	PageReq
	AssetId uint64 `json:"assetId" v:"required|min:1#资产不能为空|资产不正确" dc:"资产ID"`
}
type LocationRecordListRes struct {
	List  any `json:"list"`
	Total int `json:"total"`
}
```
