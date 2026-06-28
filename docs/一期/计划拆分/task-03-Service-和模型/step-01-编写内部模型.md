# Task 3 / Step 1: 编写内部模型

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Service 和模型](README.md)

- [ ] **Step 1: 编写内部模型**

写入 `internal/model/asset.go`。

本文件统一承载 API、Controller、Service、Logic 共用的业务输入和输出结构。`api/asset/v1` 只定义 HTTP 路由壳和响应壳，不重复声明业务字段。

```go
package model

import "asset_management_svr/internal/model/entity"

const (
	AssetStatusIdle    = 1
	AssetStatusInUse   = 2
	AssetStatusRepair  = 3
	AssetStatusScraped = 4

	AssetLocationActionInbound  = 0
	AssetLocationActionTransfer = 1
)

type PageInput struct {
	Page     int `json:"page" v:"min:1#页码不能小于1" dc:"页码"`
	PageSize int `json:"pageSize" v:"max:100#每页数量不能超过100" dc:"每页数量"`
}

type IdInput struct {
	Id uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
}

func (in PageInput) Limit() int {
	if in.PageSize <= 0 {
		return 20
	}
	if in.PageSize > 100 {
		return 100
	}
	return in.PageSize
}

func (in PageInput) Offset() int {
	page := in.Page
	if page <= 0 {
		page = 1
	}
	return (page - 1) * in.Limit()
}

type AssetBuildingCreateInput struct {
	Name    string `json:"name" v:"required#建筑名称不能为空" dc:"建筑名称"`
	Code    string `json:"code" v:"required#建筑编码不能为空" dc:"建筑编码"`
	Address string `json:"address" dc:"地址"`
	Remark  string `json:"remark" dc:"备注"`
}

type AssetBuildingUpdateInput struct {
	Id uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	AssetBuildingCreateInput
}

type AssetBuildingListInput struct {
	PageInput
	Name string `json:"name" dc:"建筑名称"`
	Code string `json:"code" dc:"建筑编码"`
}

type AssetFloorCreateInput struct {
	BuildingId uint64 `json:"buildingId" v:"required|min:1#建筑不能为空|建筑不正确" dc:"建筑ID"`
	Name       string `json:"name" v:"required#楼层名称不能为空" dc:"楼层名称"`
	Code       string `json:"code" v:"required#楼层编码不能为空" dc:"楼层编码"`
	FloorNo    int    `json:"floorNo" v:"required#楼层序号不能为空" dc:"楼层序号"`
	Remark     string `json:"remark" dc:"备注"`
}

type AssetFloorUpdateInput struct {
	Id uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	AssetFloorCreateInput
}

type AssetFloorListInput struct {
	PageInput
	BuildingId uint64 `json:"buildingId" dc:"建筑ID"`
	Name       string `json:"name" dc:"楼层名称"`
	Code       string `json:"code" dc:"楼层编码"`
}

type AssetRoomCreateInput struct {
	BuildingId uint64 `json:"buildingId" v:"required|min:1#建筑不能为空|建筑不正确" dc:"建筑ID"`
	FloorId    uint64 `json:"floorId" v:"required|min:1#楼层不能为空|楼层不正确" dc:"楼层ID"`
	Name       string `json:"name" v:"required#房间名称不能为空" dc:"房间名称"`
	Code       string `json:"code" v:"required#房间编码不能为空" dc:"房间编码"`
	RoomNo     string `json:"roomNo" dc:"房间号"`
	Remark     string `json:"remark" dc:"备注"`
}

type AssetRoomUpdateInput struct {
	Id uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	AssetRoomCreateInput
}

type AssetRoomListInput struct {
	PageInput
	BuildingId uint64 `json:"buildingId" dc:"建筑ID"`
	FloorId    uint64 `json:"floorId" dc:"楼层ID"`
	Name       string `json:"name" dc:"房间名称"`
	Code       string `json:"code" dc:"房间编码"`
}

type AssetCategoryCreateInput struct {
	ParentId uint64 `json:"parentId" dc:"父级ID"`
	Name     string `json:"name" v:"required#分类名称不能为空" dc:"分类名称"`
	Code     string `json:"code" v:"required#分类编码不能为空" dc:"分类编码"`
	Sort     int    `json:"sort" dc:"排序"`
	Remark   string `json:"remark" dc:"备注"`
}

type AssetCategoryUpdateInput struct {
	Id uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	AssetCategoryCreateInput
}

type AssetCategoryListInput struct {
	PageInput
	ParentId uint64 `json:"parentId" dc:"父级ID"`
	Name     string `json:"name" dc:"分类名称"`
	Code     string `json:"code" dc:"分类编码"`
}

type AssetProductCreateInput struct {
	CategoryId uint64 `json:"categoryId" v:"required|min:1#分类不能为空|分类不正确" dc:"分类ID"`
	Name       string `json:"name" v:"required#产品名称不能为空" dc:"产品名称"`
	Code       string `json:"code" v:"required#产品编码不能为空" dc:"产品编码"`
	Brand      string `json:"brand" dc:"品牌"`
	Model      string `json:"model" dc:"型号"`
	Unit       string `json:"unit" dc:"计量单位"`
	Remark     string `json:"remark" dc:"备注"`
}

type AssetProductUpdateInput struct {
	Id uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	AssetProductCreateInput
}

type AssetProductListInput struct {
	PageInput
	CategoryId uint64 `json:"categoryId" dc:"分类ID"`
	Name       string `json:"name" dc:"产品名称"`
	Code       string `json:"code" dc:"产品编码"`
}

type AssetItemCreateInput struct {
	ProductId     uint64  `json:"productId" v:"required|min:1#产品不能为空|产品不正确" dc:"产品ID"`
	AssetCode     string  `json:"assetCode" dc:"资产编码"`
	Name          string  `json:"name" v:"required#资产名称不能为空" dc:"资产名称"`
	Status        int     `json:"status" v:"required|in:1,2,3,4#状态不能为空|状态不正确" dc:"资产状态：1闲置 2在用 3维修 4报废"`
	RoomId        uint64  `json:"roomId" dc:"房间ID"`
	PurchaseDate  string  `json:"purchaseDate" dc:"购买日期"`
	PurchasePrice float64 `json:"purchasePrice" dc:"购买价格"`
	Remark        string  `json:"remark" dc:"备注"`
}

type AssetItemCreateOutput struct {
	Id        uint64 `json:"id" dc:"资产ID"`
	AssetCode string `json:"assetCode" dc:"资产编码"`
}

type AssetItemUpdateInput struct {
	Id uint64 `json:"id" v:"required|min:1#ID不能为空|ID不正确" dc:"ID"`
	AssetItemCreateInput
}

type AssetItemListInput struct {
	PageInput
	ProductId uint64 `json:"productId" dc:"产品ID"`
	RoomId    uint64 `json:"roomId" dc:"房间ID"`
	Name      string `json:"name" dc:"资产名称"`
	AssetCode string `json:"assetCode" dc:"资产编码"`
	Status    int    `json:"status" dc:"资产状态"`
}

type AssetItemInboundInput struct {
	AssetId uint64 `json:"assetId" v:"required|min:1#资产不能为空|资产不正确" dc:"资产ID"`
	RoomId  uint64 `json:"roomId" v:"required|min:1#房间不能为空|房间不正确" dc:"房间ID"`
	Remark  string `json:"remark" dc:"备注"`
}

type AssetItemTransferInput struct {
	AssetId uint64 `json:"assetId" v:"required|min:1#资产不能为空|资产不正确" dc:"资产ID"`
	RoomId  uint64 `json:"roomId" v:"required|min:1#房间不能为空|房间不正确" dc:"目标房间ID"`
	Remark  string `json:"remark" dc:"备注"`
}

type AssetLocationRecordListInput struct {
	PageInput
	AssetId uint64 `json:"assetId" v:"required|min:1#资产不能为空|资产不正确" dc:"资产ID"`
}

type AssetBuildingOutput = entity.AssetBuilding
type AssetFloorOutput = entity.AssetFloor
type AssetRoomOutput = entity.AssetRoom
type AssetCategoryOutput = entity.AssetCategory
type AssetProductOutput = entity.AssetProduct
type AssetItemOutput = entity.AssetItem
type AssetLocationRecordOutput = entity.AssetLocationRecord
```
