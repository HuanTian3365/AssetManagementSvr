# Task 5 / Step 1: 编写 Controller

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Controller 和路由注册](README.md)

- [ ] **Step 1: 编写 Controller**

如果 `gf gen ctrl` 已生成同名文件，在生成文件中补齐方法体；如果没有生成，则写入 `internal/controller/asset/asset.go`：

```go
package asset

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	v1 "your_project/api/asset/v1"
	"your_project/internal/model"
	"your_project/internal/service"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

func page(page int, pageSize int) model.PageInput {
	return model.PageInput{Page: page, PageSize: pageSize}
}

func (c *ControllerV1) BuildingCreate(ctx context.Context, req *v1.BuildingCreateReq) (res *v1.BuildingCreateRes, err error) {
	id, err := service.Asset().CreateBuilding(ctx, gdb.Map{"name": req.Name, "code": req.Code, "address": req.Address, "remark": req.Remark})
	return &v1.BuildingCreateRes{Id: id}, err
}

func (c *ControllerV1) BuildingUpdate(ctx context.Context, req *v1.BuildingUpdateReq) (res *v1.BuildingUpdateRes, err error) {
	return &v1.BuildingUpdateRes{}, service.Asset().UpdateBuilding(ctx, req.Id, gdb.Map{"name": req.Name, "code": req.Code, "address": req.Address, "remark": req.Remark})
}

func (c *ControllerV1) BuildingDelete(ctx context.Context, req *v1.BuildingDeleteReq) (res *v1.BuildingDeleteRes, err error) {
	return &v1.BuildingDeleteRes{}, service.Asset().DeleteBuilding(ctx, req.Id)
}

func (c *ControllerV1) BuildingDetail(ctx context.Context, req *v1.BuildingDetailReq) (res *v1.BuildingDetailRes, err error) {
	data, err := service.Asset().DetailBuilding(ctx, req.Id)
	return &v1.BuildingDetailRes{Data: data}, err
}

func (c *ControllerV1) BuildingList(ctx context.Context, req *v1.BuildingListReq) (res *v1.BuildingListRes, err error) {
	list, total, err := service.Asset().ListBuilding(ctx, page(req.Page, req.PageSize), gdb.Map{"name": req.Name, "code": req.Code})
	return &v1.BuildingListRes{List: list, Total: total}, err
}

func (c *ControllerV1) FloorCreate(ctx context.Context, req *v1.FloorCreateReq) (res *v1.FloorCreateRes, err error) {
	id, err := service.Asset().CreateFloor(ctx, gdb.Map{"building_id": req.BuildingId, "name": req.Name, "code": req.Code, "floor_no": req.FloorNo, "remark": req.Remark})
	return &v1.FloorCreateRes{Id: id}, err
}

func (c *ControllerV1) FloorUpdate(ctx context.Context, req *v1.FloorUpdateReq) (res *v1.FloorUpdateRes, err error) {
	return &v1.FloorUpdateRes{}, service.Asset().UpdateFloor(ctx, req.Id, gdb.Map{"building_id": req.BuildingId, "name": req.Name, "code": req.Code, "floor_no": req.FloorNo, "remark": req.Remark})
}

func (c *ControllerV1) FloorDelete(ctx context.Context, req *v1.FloorDeleteReq) (res *v1.FloorDeleteRes, err error) {
	return &v1.FloorDeleteRes{}, service.Asset().DeleteFloor(ctx, req.Id)
}

func (c *ControllerV1) FloorDetail(ctx context.Context, req *v1.FloorDetailReq) (res *v1.FloorDetailRes, err error) {
	data, err := service.Asset().DetailFloor(ctx, req.Id)
	return &v1.FloorDetailRes{Data: data}, err
}

func (c *ControllerV1) FloorList(ctx context.Context, req *v1.FloorListReq) (res *v1.FloorListRes, err error) {
	list, total, err := service.Asset().ListFloor(ctx, page(req.Page, req.PageSize), gdb.Map{"building_id": req.BuildingId, "name": req.Name, "code": req.Code})
	return &v1.FloorListRes{List: list, Total: total}, err
}

func (c *ControllerV1) RoomCreate(ctx context.Context, req *v1.RoomCreateReq) (res *v1.RoomCreateRes, err error) {
	id, err := service.Asset().CreateRoom(ctx, gdb.Map{"building_id": req.BuildingId, "floor_id": req.FloorId, "name": req.Name, "code": req.Code, "room_no": req.RoomNo, "remark": req.Remark})
	return &v1.RoomCreateRes{Id: id}, err
}

func (c *ControllerV1) RoomUpdate(ctx context.Context, req *v1.RoomUpdateReq) (res *v1.RoomUpdateRes, err error) {
	return &v1.RoomUpdateRes{}, service.Asset().UpdateRoom(ctx, req.Id, gdb.Map{"building_id": req.BuildingId, "floor_id": req.FloorId, "name": req.Name, "code": req.Code, "room_no": req.RoomNo, "remark": req.Remark})
}

func (c *ControllerV1) RoomDelete(ctx context.Context, req *v1.RoomDeleteReq) (res *v1.RoomDeleteRes, err error) {
	return &v1.RoomDeleteRes{}, service.Asset().DeleteRoom(ctx, req.Id)
}

func (c *ControllerV1) RoomDetail(ctx context.Context, req *v1.RoomDetailReq) (res *v1.RoomDetailRes, err error) {
	data, err := service.Asset().DetailRoom(ctx, req.Id)
	return &v1.RoomDetailRes{Data: data}, err
}

func (c *ControllerV1) RoomList(ctx context.Context, req *v1.RoomListReq) (res *v1.RoomListRes, err error) {
	list, total, err := service.Asset().ListRoom(ctx, page(req.Page, req.PageSize), gdb.Map{"building_id": req.BuildingId, "floor_id": req.FloorId, "name": req.Name, "code": req.Code})
	return &v1.RoomListRes{List: list, Total: total}, err
}

func (c *ControllerV1) CategoryCreate(ctx context.Context, req *v1.CategoryCreateReq) (res *v1.CategoryCreateRes, err error) {
	id, err := service.Asset().CreateCategory(ctx, gdb.Map{"parent_id": req.ParentId, "name": req.Name, "code": req.Code, "sort": req.Sort, "remark": req.Remark})
	return &v1.CategoryCreateRes{Id: id}, err
}

func (c *ControllerV1) CategoryUpdate(ctx context.Context, req *v1.CategoryUpdateReq) (res *v1.CategoryUpdateRes, err error) {
	return &v1.CategoryUpdateRes{}, service.Asset().UpdateCategory(ctx, req.Id, gdb.Map{"parent_id": req.ParentId, "name": req.Name, "code": req.Code, "sort": req.Sort, "remark": req.Remark})
}

func (c *ControllerV1) CategoryDelete(ctx context.Context, req *v1.CategoryDeleteReq) (res *v1.CategoryDeleteRes, err error) {
	return &v1.CategoryDeleteRes{}, service.Asset().DeleteCategory(ctx, req.Id)
}

func (c *ControllerV1) CategoryDetail(ctx context.Context, req *v1.CategoryDetailReq) (res *v1.CategoryDetailRes, err error) {
	data, err := service.Asset().DetailCategory(ctx, req.Id)
	return &v1.CategoryDetailRes{Data: data}, err
}

func (c *ControllerV1) CategoryList(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error) {
	list, total, err := service.Asset().ListCategory(ctx, page(req.Page, req.PageSize), gdb.Map{"parent_id": req.ParentId, "name": req.Name, "code": req.Code})
	return &v1.CategoryListRes{List: list, Total: total}, err
}

func (c *ControllerV1) ProductCreate(ctx context.Context, req *v1.ProductCreateReq) (res *v1.ProductCreateRes, err error) {
	id, err := service.Asset().CreateProduct(ctx, gdb.Map{"category_id": req.CategoryId, "name": req.Name, "code": req.Code, "brand": req.Brand, "model": req.Model, "unit": req.Unit, "remark": req.Remark})
	return &v1.ProductCreateRes{Id: id}, err
}

func (c *ControllerV1) ProductUpdate(ctx context.Context, req *v1.ProductUpdateReq) (res *v1.ProductUpdateRes, err error) {
	return &v1.ProductUpdateRes{}, service.Asset().UpdateProduct(ctx, req.Id, gdb.Map{"category_id": req.CategoryId, "name": req.Name, "code": req.Code, "brand": req.Brand, "model": req.Model, "unit": req.Unit, "remark": req.Remark})
}

func (c *ControllerV1) ProductDelete(ctx context.Context, req *v1.ProductDeleteReq) (res *v1.ProductDeleteRes, err error) {
	return &v1.ProductDeleteRes{}, service.Asset().DeleteProduct(ctx, req.Id)
}

func (c *ControllerV1) ProductDetail(ctx context.Context, req *v1.ProductDetailReq) (res *v1.ProductDetailRes, err error) {
	data, err := service.Asset().DetailProduct(ctx, req.Id)
	return &v1.ProductDetailRes{Data: data}, err
}

func (c *ControllerV1) ProductList(ctx context.Context, req *v1.ProductListReq) (res *v1.ProductListRes, err error) {
	list, total, err := service.Asset().ListProduct(ctx, page(req.Page, req.PageSize), gdb.Map{"category_id": req.CategoryId, "name": req.Name, "code": req.Code})
	return &v1.ProductListRes{List: list, Total: total}, err
}

func (c *ControllerV1) ItemCreate(ctx context.Context, req *v1.ItemCreateReq) (res *v1.ItemCreateRes, err error) {
	out, err := service.Asset().CreateItem(ctx, model.AssetItemCreateInput{ProductId: req.ProductId, AssetCode: req.AssetCode, Name: req.Name, Status: req.Status, RoomId: req.RoomId, PurchaseDate: req.PurchaseDate, PurchasePrice: req.PurchasePrice, Remark: req.Remark})
	return &v1.ItemCreateRes{Id: out.Id, AssetCode: out.AssetCode}, err
}

func (c *ControllerV1) ItemUpdate(ctx context.Context, req *v1.ItemUpdateReq) (res *v1.ItemUpdateRes, err error) {
	return &v1.ItemUpdateRes{}, service.Asset().UpdateItem(ctx, model.AssetItemUpdateInput{Id: req.Id, ProductId: req.ProductId, AssetCode: req.AssetCode, Name: req.Name, Status: req.Status, PurchaseDate: req.PurchaseDate, PurchasePrice: req.PurchasePrice, Remark: req.Remark})
}

func (c *ControllerV1) ItemDelete(ctx context.Context, req *v1.ItemDeleteReq) (res *v1.ItemDeleteRes, err error) {
	return &v1.ItemDeleteRes{}, service.Asset().DeleteItem(ctx, req.Id)
}

func (c *ControllerV1) ItemDetail(ctx context.Context, req *v1.ItemDetailReq) (res *v1.ItemDetailRes, err error) {
	data, err := service.Asset().DetailItem(ctx, req.Id)
	return &v1.ItemDetailRes{Data: data}, err
}

func (c *ControllerV1) ItemList(ctx context.Context, req *v1.ItemListReq) (res *v1.ItemListRes, err error) {
	list, total, err := service.Asset().ListItem(ctx, page(req.Page, req.PageSize), gdb.Map{"product_id": req.ProductId, "current_room_id": req.RoomId, "name": req.Name, "asset_code": req.AssetCode, "status": req.Status})
	return &v1.ItemListRes{List: list, Total: total}, err
}

func (c *ControllerV1) ItemInbound(ctx context.Context, req *v1.ItemInboundReq) (res *v1.ItemInboundRes, err error) {
	return &v1.ItemInboundRes{}, service.Asset().InboundItem(ctx, req.AssetId, req.RoomId, req.Remark)
}

func (c *ControllerV1) ItemTransfer(ctx context.Context, req *v1.ItemTransferReq) (res *v1.ItemTransferRes, err error) {
	return &v1.ItemTransferRes{}, service.Asset().TransferItem(ctx, req.AssetId, req.RoomId, req.Remark)
}

func (c *ControllerV1) LocationRecordList(ctx context.Context, req *v1.LocationRecordListReq) (res *v1.LocationRecordListRes, err error) {
	list, total, err := service.Asset().ListLocationRecord(ctx, req.AssetId, page(req.Page, req.PageSize))
	return &v1.LocationRecordListRes{List: list, Total: total}, err
}
```

将 `your_project/...` 替换为项目实际 module 路径对应的导入路径。
