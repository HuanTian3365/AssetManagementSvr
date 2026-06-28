# Task 3 / Step 1: 编写内部模型

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Service 和模型](README.md)

- [ ] **Step 1: 编写内部模型**

写入 `internal/model/asset.go`：

```go
package model

const (
	AssetStatusIdle    = 1
	AssetStatusInUse   = 2
	AssetStatusRepair  = 3
	AssetStatusScraped = 4

	AssetLocationActionInbound  = 0
	AssetLocationActionTransfer = 1
)

type PageInput struct {
	Page     int
	PageSize int
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

type AssetItemCreateInput struct {
	ProductId     uint64
	AssetCode     string
	Name          string
	Status        int
	RoomId        uint64
	PurchaseDate  string
	PurchasePrice float64
	Remark        string
}

type AssetItemCreateOutput struct {
	Id        uint64
	AssetCode string
}

type AssetItemUpdateInput struct {
	Id            uint64
	ProductId     uint64
	AssetCode     string
	Name          string
	Status        int
	PurchaseDate  string
	PurchasePrice float64
	Remark        string
}
```
