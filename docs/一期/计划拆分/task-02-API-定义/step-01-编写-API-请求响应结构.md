# Task 2 / Step 1: 编写 API 请求响应结构

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [API 定义](README.md)

- [ ] **Step 1: 编写 API 请求响应结构**

写入 `api/asset/v1/asset.go`。

约定：

- `api/asset/v1` 只定义 HTTP 入口壳、路由元信息和响应外壳。
- 请求业务字段统一复用 `internal/model/asset.go` 中的输入模型。
- Controller、Service、Logic 继续沿用这些 `model` 结构，避免 API 和业务层重复维护同一批字段。

```go
package v1

import (
	"asset_management_svr/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type BuildingCreateReq struct {
	g.Meta `path:"/asset/building/create" method:"post" tags:"资产建筑" summary:"创建建筑"`
	model.AssetBuildingCreateInput
}
type BuildingCreateRes struct{ Id uint64 `json:"id" dc:"建筑ID"` }

type BuildingUpdateReq struct {
	g.Meta `path:"/asset/building/update" method:"put" tags:"资产建筑" summary:"更新建筑"`
	model.AssetBuildingUpdateInput
}
type BuildingUpdateRes struct{}

type BuildingDeleteReq struct {
	g.Meta `path:"/asset/building/delete" method:"delete" tags:"资产建筑" summary:"删除建筑"`
	model.IdInput
}
type BuildingDeleteRes struct{}

type BuildingViewReq struct {
	g.Meta `path:"/asset/building/view" method:"get" tags:"资产建筑" summary:"建筑详情"`
	model.IdInput
}
type BuildingViewRes struct{ Data *model.AssetBuildingOutput `json:"data" dc:"建筑详情"` }

type BuildingListReq struct {
	g.Meta `path:"/asset/building/list" method:"get" tags:"资产建筑" summary:"建筑列表"`
	model.AssetBuildingListInput
}
type BuildingListRes struct {
	List  []*model.AssetBuildingOutput `json:"list" dc:"建筑列表"`
	Total int                          `json:"total" dc:"总条数"`
}

type FloorCreateReq struct {
	g.Meta `path:"/asset/floor/create" method:"post" tags:"资产楼层" summary:"创建楼层"`
	model.AssetFloorCreateInput
}
type FloorCreateRes struct{ Id uint64 `json:"id" dc:"楼层ID"` }

type FloorUpdateReq struct {
	g.Meta `path:"/asset/floor/update" method:"put" tags:"资产楼层" summary:"更新楼层"`
	model.AssetFloorUpdateInput
}
type FloorUpdateRes struct{}

type FloorDeleteReq struct {
	g.Meta `path:"/asset/floor/delete" method:"delete" tags:"资产楼层" summary:"删除楼层"`
	model.IdInput
}
type FloorDeleteRes struct{}

type FloorDetailReq struct {
	g.Meta `path:"/asset/floor/detail" method:"get" tags:"资产楼层" summary:"楼层详情"`
	model.IdInput
}
type FloorDetailRes struct{ Data *model.AssetFloorOutput `json:"data" dc:"楼层详情"` }

type FloorListReq struct {
	g.Meta `path:"/asset/floor/list" method:"get" tags:"资产楼层" summary:"楼层列表"`
	model.AssetFloorListInput
}
type FloorListRes struct {
	List  []*model.AssetFloorOutput `json:"list" dc:"楼层列表"`
	Total int                       `json:"total" dc:"总数"`
}

type RoomCreateReq struct {
	g.Meta `path:"/asset/room/create" method:"post" tags:"资产房间" summary:"创建房间"`
	model.AssetRoomCreateInput
}
type RoomCreateRes struct{ Id uint64 `json:"id" dc:"房间ID"` }

type RoomUpdateReq struct {
	g.Meta `path:"/asset/room/update" method:"put" tags:"资产房间" summary:"更新房间"`
	model.AssetRoomUpdateInput
}
type RoomUpdateRes struct{}

type RoomDeleteReq struct {
	g.Meta `path:"/asset/room/delete" method:"delete" tags:"资产房间" summary:"删除房间"`
	model.IdInput
}
type RoomDeleteRes struct{}

type RoomDetailReq struct {
	g.Meta `path:"/asset/room/detail" method:"get" tags:"资产房间" summary:"房间详情"`
	model.IdInput
}
type RoomDetailRes struct{ Data *model.AssetRoomOutput `json:"data" dc:"房间详情"` }

type RoomListReq struct {
	g.Meta `path:"/asset/room/list" method:"get" tags:"资产房间" summary:"房间列表"`
	model.AssetRoomListInput
}
type RoomListRes struct {
	List  []*model.AssetRoomOutput `json:"list" dc:"房间列表"`
	Total int                      `json:"total" dc:"总数"`
}

type CategoryCreateReq struct {
	g.Meta `path:"/asset/category/create" method:"post" tags:"资产分类" summary:"创建分类"`
	model.AssetCategoryCreateInput
}
type CategoryCreateRes struct{ Id uint64 `json:"id" dc:"分类ID"` }

type CategoryUpdateReq struct {
	g.Meta `path:"/asset/category/update" method:"put" tags:"资产分类" summary:"更新分类"`
	model.AssetCategoryUpdateInput
}
type CategoryUpdateRes struct{}

type CategoryDeleteReq struct {
	g.Meta `path:"/asset/category/delete" method:"delete" tags:"资产分类" summary:"删除分类"`
	model.IdInput
}
type CategoryDeleteRes struct{}

type CategoryDetailReq struct {
	g.Meta `path:"/asset/category/detail" method:"get" tags:"资产分类" summary:"分类详情"`
	model.IdInput
}
type CategoryDetailRes struct{ Data *model.AssetCategoryOutput `json:"data" dc:"分类详情"` }

type CategoryListReq struct {
	g.Meta `path:"/asset/category/list" method:"get" tags:"资产分类" summary:"分类列表"`
	model.AssetCategoryListInput
}
type CategoryListRes struct {
	List  []*model.AssetCategoryOutput `json:"list" dc:"分类列表"`
	Total int                          `json:"total" dc:"总数"`
}

type ProductCreateReq struct {
	g.Meta `path:"/asset/product/create" method:"post" tags:"资产产品" summary:"创建产品"`
	model.AssetProductCreateInput
}
type ProductCreateRes struct{ Id uint64 `json:"id" dc:"产品ID"` }

type ProductUpdateReq struct {
	g.Meta `path:"/asset/product/update" method:"put" tags:"资产产品" summary:"更新产品"`
	model.AssetProductUpdateInput
}
type ProductUpdateRes struct{}

type ProductDeleteReq struct {
	g.Meta `path:"/asset/product/delete" method:"delete" tags:"资产产品" summary:"删除产品"`
	model.IdInput
}
type ProductDeleteRes struct{}

type ProductDetailReq struct {
	g.Meta `path:"/asset/product/detail" method:"get" tags:"资产产品" summary:"产品详情"`
	model.IdInput
}
type ProductDetailRes struct{ Data *model.AssetProductOutput `json:"data" dc:"产品详情"` }

type ProductListReq struct {
	g.Meta `path:"/asset/product/list" method:"get" tags:"资产产品" summary:"产品列表"`
	model.AssetProductListInput
}
type ProductListRes struct {
	List  []*model.AssetProductOutput `json:"list" dc:"产品列表"`
	Total int                         `json:"total" dc:"总数"`
}

type ItemCreateReq struct {
	g.Meta `path:"/asset/item/create" method:"post" tags:"资产实例" summary:"创建资产"`
	model.AssetItemCreateInput
}
type ItemCreateRes struct {
	Id        uint64 `json:"id" dc:"资产ID"`
	AssetCode string `json:"assetCode" dc:"资产编码"`
}

type ItemUpdateReq struct {
	g.Meta `path:"/asset/item/update" method:"put" tags:"资产实例" summary:"更新资产"`
	model.AssetItemUpdateInput
}
type ItemUpdateRes struct{}

type ItemDeleteReq struct {
	g.Meta `path:"/asset/item/delete" method:"delete" tags:"资产实例" summary:"删除资产"`
	model.IdInput
}
type ItemDeleteRes struct{}

type ItemDetailReq struct {
	g.Meta `path:"/asset/item/detail" method:"get" tags:"资产实例" summary:"资产详情"`
	model.IdInput
}
type ItemDetailRes struct{ Data *model.AssetItemOutput `json:"data" dc:"资产详情"` }

type ItemListReq struct {
	g.Meta `path:"/asset/item/list" method:"get" tags:"资产实例" summary:"资产列表"`
	model.AssetItemListInput
}
type ItemListRes struct {
	List  []*model.AssetItemOutput `json:"list" dc:"资产列表"`
	Total int                      `json:"total" dc:"总数"`
}

type ItemInboundReq struct {
	g.Meta `path:"/asset/item/inbound" method:"post" tags:"资产实例" summary:"资产入库"`
	model.AssetItemInboundInput
}
type ItemInboundRes struct{}

type ItemTransferReq struct {
	g.Meta `path:"/asset/item/transfer" method:"post" tags:"资产实例" summary:"资产转移"`
	model.AssetItemTransferInput
}
type ItemTransferRes struct{}

type LocationRecordListReq struct {
	g.Meta `path:"/asset/item/location-record/list" method:"get" tags:"资产实例" summary:"资产位置记录"`
	model.AssetLocationRecordListInput
}
type LocationRecordListRes struct {
	List  []*model.AssetLocationRecordOutput `json:"list" dc:"资产位置记录列表"`
	Total int                                `json:"total" dc:"总数"`
}
```
