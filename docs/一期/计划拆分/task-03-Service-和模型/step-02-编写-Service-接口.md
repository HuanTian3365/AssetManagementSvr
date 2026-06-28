# Task 3 / Step 2: 编写 Service 接口

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Service 和模型](README.md)

- [ ] **Step 2: 编写 Service 接口**

写入 `internal/service/asset.go`：

```go
package service

import (
	"context"

	"asset_management_svr/internal/model"
)

type IAsset interface {
	CreateBuilding(ctx context.Context, in *model.AssetBuildingCreateInput) (uint64, error)
	UpdateBuilding(ctx context.Context, in *model.AssetBuildingUpdateInput) error
	DeleteBuilding(ctx context.Context, id uint64) error
	BuildingView(ctx context.Context, id uint64) (*model.AssetBuildingOutput, error)
	ListBuilding(ctx context.Context, in *model.AssetBuildingListInput) ([]*model.AssetBuildingOutput, int, error)

	CreateFloor(ctx context.Context, in *model.AssetFloorCreateInput) (uint64, error)
	UpdateFloor(ctx context.Context, in *model.AssetFloorUpdateInput) error
	DeleteFloor(ctx context.Context, id uint64) error
	DetailFloor(ctx context.Context, id uint64) (*model.AssetFloorOutput, error)
	ListFloor(ctx context.Context, in *model.AssetFloorListInput) ([]*model.AssetFloorOutput, int, error)

	CreateRoom(ctx context.Context, in *model.AssetRoomCreateInput) (uint64, error)
	UpdateRoom(ctx context.Context, in *model.AssetRoomUpdateInput) error
	DeleteRoom(ctx context.Context, id uint64) error
	DetailRoom(ctx context.Context, id uint64) (*model.AssetRoomOutput, error)
	ListRoom(ctx context.Context, in *model.AssetRoomListInput) ([]*model.AssetRoomOutput, int, error)

	CreateCategory(ctx context.Context, in *model.AssetCategoryCreateInput) (uint64, error)
	UpdateCategory(ctx context.Context, in *model.AssetCategoryUpdateInput) error
	DeleteCategory(ctx context.Context, id uint64) error
	DetailCategory(ctx context.Context, id uint64) (*model.AssetCategoryOutput, error)
	ListCategory(ctx context.Context, in *model.AssetCategoryListInput) ([]*model.AssetCategoryOutput, int, error)

	CreateProduct(ctx context.Context, in *model.AssetProductCreateInput) (uint64, error)
	UpdateProduct(ctx context.Context, in *model.AssetProductUpdateInput) error
	DeleteProduct(ctx context.Context, id uint64) error
	DetailProduct(ctx context.Context, id uint64) (*model.AssetProductOutput, error)
	ListProduct(ctx context.Context, in *model.AssetProductListInput) ([]*model.AssetProductOutput, int, error)

	CreateItem(ctx context.Context, in *model.AssetItemCreateInput) (*model.AssetItemCreateOutput, error)
	UpdateItem(ctx context.Context, in *model.AssetItemUpdateInput) error
	DeleteItem(ctx context.Context, id uint64) error
	DetailItem(ctx context.Context, id uint64) (*model.AssetItemOutput, error)
	ListItem(ctx context.Context, in *model.AssetItemListInput) ([]*model.AssetItemOutput, int, error)
	InboundItem(ctx context.Context, in *model.AssetItemInboundInput) error
	TransferItem(ctx context.Context, in *model.AssetItemTransferInput) error
	ListLocationRecord(ctx context.Context, in *model.AssetLocationRecordListInput) ([]*model.AssetLocationRecordOutput, int, error)
}

var localAsset IAsset

func Asset() IAsset {
	if localAsset == nil {
		panic("implement not found for interface IAsset")
	}
	return localAsset
}

func RegisterAsset(i IAsset) {
	localAsset = i
}
```

当前项目 `go.mod` 的 module 为 `asset_management_svr`，示例导入路径已按该 module 编写。
