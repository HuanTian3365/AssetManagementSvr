# Task 3 / Step 2: 编写 Service 接口

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Service 和模型](README.md)

- [ ] **Step 2: 编写 Service 接口**

写入 `internal/service/asset.go`：

```go
package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"your_project/internal/model"
)

type IAsset interface {
	CreateBuilding(ctx context.Context, data gdb.Map) (uint64, error)
	UpdateBuilding(ctx context.Context, id uint64, data gdb.Map) error
	DeleteBuilding(ctx context.Context, id uint64) error
	DetailBuilding(ctx context.Context, id uint64) (gdb.Record, error)
	ListBuilding(ctx context.Context, page model.PageInput, filters gdb.Map) (gdb.Result, int, error)

	CreateFloor(ctx context.Context, data gdb.Map) (uint64, error)
	UpdateFloor(ctx context.Context, id uint64, data gdb.Map) error
	DeleteFloor(ctx context.Context, id uint64) error
	DetailFloor(ctx context.Context, id uint64) (gdb.Record, error)
	ListFloor(ctx context.Context, page model.PageInput, filters gdb.Map) (gdb.Result, int, error)

	CreateRoom(ctx context.Context, data gdb.Map) (uint64, error)
	UpdateRoom(ctx context.Context, id uint64, data gdb.Map) error
	DeleteRoom(ctx context.Context, id uint64) error
	DetailRoom(ctx context.Context, id uint64) (gdb.Record, error)
	ListRoom(ctx context.Context, page model.PageInput, filters gdb.Map) (gdb.Result, int, error)

	CreateCategory(ctx context.Context, data gdb.Map) (uint64, error)
	UpdateCategory(ctx context.Context, id uint64, data gdb.Map) error
	DeleteCategory(ctx context.Context, id uint64) error
	DetailCategory(ctx context.Context, id uint64) (gdb.Record, error)
	ListCategory(ctx context.Context, page model.PageInput, filters gdb.Map) (gdb.Result, int, error)

	CreateProduct(ctx context.Context, data gdb.Map) (uint64, error)
	UpdateProduct(ctx context.Context, id uint64, data gdb.Map) error
	DeleteProduct(ctx context.Context, id uint64) error
	DetailProduct(ctx context.Context, id uint64) (gdb.Record, error)
	ListProduct(ctx context.Context, page model.PageInput, filters gdb.Map) (gdb.Result, int, error)

	CreateItem(ctx context.Context, in model.AssetItemCreateInput) (model.AssetItemCreateOutput, error)
	UpdateItem(ctx context.Context, in model.AssetItemUpdateInput) error
	DeleteItem(ctx context.Context, id uint64) error
	DetailItem(ctx context.Context, id uint64) (gdb.Record, error)
	ListItem(ctx context.Context, page model.PageInput, filters gdb.Map) (gdb.Result, int, error)
	InboundItem(ctx context.Context, assetId uint64, roomId uint64, remark string) error
	TransferItem(ctx context.Context, assetId uint64, roomId uint64, remark string) error
	ListLocationRecord(ctx context.Context, assetId uint64, page model.PageInput) (gdb.Result, int, error)
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

将 `your_project/internal/model` 替换为项目实际 module 路径对应的导入路径。这个导入路径来自 `go.mod` 的 `module` 行。
