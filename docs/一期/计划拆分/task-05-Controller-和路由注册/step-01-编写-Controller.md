# Task 5 / Step 1: 编写 Controller

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Controller 和路由注册](README.md)

- [ ] **Step 1: 编写 Controller**

如果 `gf gen ctrl` 已生成同名文件，在生成文件中补齐方法体；如果没有生成，则写入 `internal/controller/asset/asset.go`。

Controller 只负责 HTTP 请求到 Service 的转发，不在这里重新拼接业务字段。API 请求结构已经嵌入 `internal/model` 的业务输入结构，因此这里直接把嵌入模型传给 Service。

```go
package asset

import (
	"context"

	v1 "asset_management_svr/api/asset/v1"
	"asset_management_svr/internal/service"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) BuildingCreate(ctx context.Context, req *v1.BuildingCreateReq) (res *v1.BuildingCreateRes, err error) {
	id, err := service.Asset().CreateBuilding(ctx, &req.AssetBuildingCreateInput)
	return &v1.BuildingCreateRes{Id: id}, err
}

func (c *ControllerV1) BuildingUpdate(ctx context.Context, req *v1.BuildingUpdateReq) (res *v1.BuildingUpdateRes, err error) {
	return &v1.BuildingUpdateRes{}, service.Asset().UpdateBuilding(ctx, &req.AssetBuildingUpdateInput)
}

func (c *ControllerV1) BuildingDelete(ctx context.Context, req *v1.BuildingDeleteReq) (res *v1.BuildingDeleteRes, err error) {
	return &v1.BuildingDeleteRes{}, service.Asset().DeleteBuilding(ctx, req.Id)
}

func (c *ControllerV1) BuildingView(ctx context.Context, req *v1.BuildingViewReq) (res *v1.BuildingViewRes, err error) {
	data, err := service.Asset().BuildingView(ctx, req.Id)
	return &v1.BuildingViewRes{Data: data}, err
}

func (c *ControllerV1) BuildingList(ctx context.Context, req *v1.BuildingListReq) (res *v1.BuildingListRes, err error) {
	list, total, err := service.Asset().ListBuilding(ctx, &req.AssetBuildingListInput)
	return &v1.BuildingListRes{List: list, Total: total}, err
}

func (c *ControllerV1) FloorCreate(ctx context.Context, req *v1.FloorCreateReq) (res *v1.FloorCreateRes, err error) {
	id, err := service.Asset().CreateFloor(ctx, &req.AssetFloorCreateInput)
	return &v1.FloorCreateRes{Id: id}, err
}

func (c *ControllerV1) FloorUpdate(ctx context.Context, req *v1.FloorUpdateReq) (res *v1.FloorUpdateRes, err error) {
	return &v1.FloorUpdateRes{}, service.Asset().UpdateFloor(ctx, &req.AssetFloorUpdateInput)
}

func (c *ControllerV1) FloorDelete(ctx context.Context, req *v1.FloorDeleteReq) (res *v1.FloorDeleteRes, err error) {
	return &v1.FloorDeleteRes{}, service.Asset().DeleteFloor(ctx, req.Id)
}

func (c *ControllerV1) FloorDetail(ctx context.Context, req *v1.FloorDetailReq) (res *v1.FloorDetailRes, err error) {
	data, err := service.Asset().DetailFloor(ctx, req.Id)
	return &v1.FloorDetailRes{Data: data}, err
}

func (c *ControllerV1) FloorList(ctx context.Context, req *v1.FloorListReq) (res *v1.FloorListRes, err error) {
	list, total, err := service.Asset().ListFloor(ctx, &req.AssetFloorListInput)
	return &v1.FloorListRes{List: list, Total: total}, err
}

func (c *ControllerV1) RoomCreate(ctx context.Context, req *v1.RoomCreateReq) (res *v1.RoomCreateRes, err error) {
	id, err := service.Asset().CreateRoom(ctx, &req.AssetRoomCreateInput)
	return &v1.RoomCreateRes{Id: id}, err
}

func (c *ControllerV1) RoomUpdate(ctx context.Context, req *v1.RoomUpdateReq) (res *v1.RoomUpdateRes, err error) {
	return &v1.RoomUpdateRes{}, service.Asset().UpdateRoom(ctx, &req.AssetRoomUpdateInput)
}

func (c *ControllerV1) RoomDelete(ctx context.Context, req *v1.RoomDeleteReq) (res *v1.RoomDeleteRes, err error) {
	return &v1.RoomDeleteRes{}, service.Asset().DeleteRoom(ctx, req.Id)
}

func (c *ControllerV1) RoomDetail(ctx context.Context, req *v1.RoomDetailReq) (res *v1.RoomDetailRes, err error) {
	data, err := service.Asset().DetailRoom(ctx, req.Id)
	return &v1.RoomDetailRes{Data: data}, err
}

func (c *ControllerV1) RoomList(ctx context.Context, req *v1.RoomListReq) (res *v1.RoomListRes, err error) {
	list, total, err := service.Asset().ListRoom(ctx, &req.AssetRoomListInput)
	return &v1.RoomListRes{List: list, Total: total}, err
}

func (c *ControllerV1) CategoryCreate(ctx context.Context, req *v1.CategoryCreateReq) (res *v1.CategoryCreateRes, err error) {
	id, err := service.Asset().CreateCategory(ctx, &req.AssetCategoryCreateInput)
	return &v1.CategoryCreateRes{Id: id}, err
}

func (c *ControllerV1) CategoryUpdate(ctx context.Context, req *v1.CategoryUpdateReq) (res *v1.CategoryUpdateRes, err error) {
	return &v1.CategoryUpdateRes{}, service.Asset().UpdateCategory(ctx, &req.AssetCategoryUpdateInput)
}

func (c *ControllerV1) CategoryDelete(ctx context.Context, req *v1.CategoryDeleteReq) (res *v1.CategoryDeleteRes, err error) {
	return &v1.CategoryDeleteRes{}, service.Asset().DeleteCategory(ctx, req.Id)
}

func (c *ControllerV1) CategoryDetail(ctx context.Context, req *v1.CategoryDetailReq) (res *v1.CategoryDetailRes, err error) {
	data, err := service.Asset().DetailCategory(ctx, req.Id)
	return &v1.CategoryDetailRes{Data: data}, err
}

func (c *ControllerV1) CategoryList(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error) {
	list, total, err := service.Asset().ListCategory(ctx, &req.AssetCategoryListInput)
	return &v1.CategoryListRes{List: list, Total: total}, err
}

func (c *ControllerV1) ProductCreate(ctx context.Context, req *v1.ProductCreateReq) (res *v1.ProductCreateRes, err error) {
	id, err := service.Asset().CreateProduct(ctx, &req.AssetProductCreateInput)
	return &v1.ProductCreateRes{Id: id}, err
}

func (c *ControllerV1) ProductUpdate(ctx context.Context, req *v1.ProductUpdateReq) (res *v1.ProductUpdateRes, err error) {
	return &v1.ProductUpdateRes{}, service.Asset().UpdateProduct(ctx, &req.AssetProductUpdateInput)
}

func (c *ControllerV1) ProductDelete(ctx context.Context, req *v1.ProductDeleteReq) (res *v1.ProductDeleteRes, err error) {
	return &v1.ProductDeleteRes{}, service.Asset().DeleteProduct(ctx, req.Id)
}

func (c *ControllerV1) ProductDetail(ctx context.Context, req *v1.ProductDetailReq) (res *v1.ProductDetailRes, err error) {
	data, err := service.Asset().DetailProduct(ctx, req.Id)
	return &v1.ProductDetailRes{Data: data}, err
}

func (c *ControllerV1) ProductList(ctx context.Context, req *v1.ProductListReq) (res *v1.ProductListRes, err error) {
	list, total, err := service.Asset().ListProduct(ctx, &req.AssetProductListInput)
	return &v1.ProductListRes{List: list, Total: total}, err
}

func (c *ControllerV1) ItemCreate(ctx context.Context, req *v1.ItemCreateReq) (res *v1.ItemCreateRes, err error) {
	out, err := service.Asset().CreateItem(ctx, &req.AssetItemCreateInput)
	if out == nil {
		return nil, err
	}
	return &v1.ItemCreateRes{Id: out.Id, AssetCode: out.AssetCode}, err
}

func (c *ControllerV1) ItemUpdate(ctx context.Context, req *v1.ItemUpdateReq) (res *v1.ItemUpdateRes, err error) {
	return &v1.ItemUpdateRes{}, service.Asset().UpdateItem(ctx, &req.AssetItemUpdateInput)
}

func (c *ControllerV1) ItemDelete(ctx context.Context, req *v1.ItemDeleteReq) (res *v1.ItemDeleteRes, err error) {
	return &v1.ItemDeleteRes{}, service.Asset().DeleteItem(ctx, req.Id)
}

func (c *ControllerV1) ItemDetail(ctx context.Context, req *v1.ItemDetailReq) (res *v1.ItemDetailRes, err error) {
	data, err := service.Asset().DetailItem(ctx, req.Id)
	return &v1.ItemDetailRes{Data: data}, err
}

func (c *ControllerV1) ItemList(ctx context.Context, req *v1.ItemListReq) (res *v1.ItemListRes, err error) {
	list, total, err := service.Asset().ListItem(ctx, &req.AssetItemListInput)
	return &v1.ItemListRes{List: list, Total: total}, err
}

func (c *ControllerV1) ItemInbound(ctx context.Context, req *v1.ItemInboundReq) (res *v1.ItemInboundRes, err error) {
	return &v1.ItemInboundRes{}, service.Asset().InboundItem(ctx, &req.AssetItemInboundInput)
}

func (c *ControllerV1) ItemTransfer(ctx context.Context, req *v1.ItemTransferReq) (res *v1.ItemTransferRes, err error) {
	return &v1.ItemTransferRes{}, service.Asset().TransferItem(ctx, &req.AssetItemTransferInput)
}

func (c *ControllerV1) LocationRecordList(ctx context.Context, req *v1.LocationRecordListReq) (res *v1.LocationRecordListRes, err error) {
	list, total, err := service.Asset().ListLocationRecord(ctx, &req.AssetLocationRecordListInput)
	return &v1.LocationRecordListRes{List: list, Total: total}, err
}
```
