// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package asset

import (
	"context"

	"asset_management_svr/api/asset/v1"
)

type IAssetV1 interface {
	BuildingCreate(ctx context.Context, req *v1.BuildingCreateReq) (res *v1.BuildingCreateRes, err error)
	BuildingUpdate(ctx context.Context, req *v1.BuildingUpdateReq) (res *v1.BuildingUpdateRes, err error)
	BuildingDelete(ctx context.Context, req *v1.BuildingDeleteReq) (res *v1.BuildingDeleteRes, err error)
	BuildingView(ctx context.Context, req *v1.BuildingViewReq) (res *v1.BuildingViewRes, err error)
	BuildingList(ctx context.Context, req *v1.BuildingListReq) (res *v1.BuildingListRes, err error)
	CategoryCreate(ctx context.Context, req *v1.CategoryCreateReq) (res *v1.CategoryCreateRes, err error)
	CategoryUpdate(ctx context.Context, req *v1.CategoryUpdateReq) (res *v1.CategoryUpdateRes, err error)
	CategoryDelete(ctx context.Context, req *v1.CategoryDeleteReq) (res *v1.CategoryDeleteRes, err error)
	CategoryView(ctx context.Context, req *v1.CategoryViewReq) (res *v1.CategoryViewRes, err error)
	CategoryList(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error)
	FloorCreate(ctx context.Context, req *v1.FloorCreateReq) (res *v1.FloorCreateRes, err error)
	FloorUpdate(ctx context.Context, req *v1.FloorUpdateReq) (res *v1.FloorUpdateRes, err error)
	FloorDelete(ctx context.Context, req *v1.FloorDeleteReq) (res *v1.FloorDeleteRes, err error)
	FloorView(ctx context.Context, req *v1.FloorViewReq) (res *v1.FloorViewRes, err error)
	FloorList(ctx context.Context, req *v1.FloorListReq) (res *v1.FloorListRes, err error)
	ItemCreate(ctx context.Context, req *v1.ItemCreateReq) (res *v1.ItemCreateRes, err error)
	ItemUpdate(ctx context.Context, req *v1.ItemUpdateReq) (res *v1.ItemUpdateRes, err error)
	ItemDelete(ctx context.Context, req *v1.ItemDeleteReq) (res *v1.ItemDeleteRes, err error)
	ItemView(ctx context.Context, req *v1.ItemViewReq) (res *v1.ItemViewRes, err error)
	ItemList(ctx context.Context, req *v1.ItemListReq) (res *v1.ItemListRes, err error)
	ProductCreate(ctx context.Context, req *v1.ProductCreateReq) (res *v1.ProductCreateRes, err error)
	ProductUpdate(ctx context.Context, req *v1.ProductUpdateReq) (res *v1.ProductUpdateRes, err error)
	ProductDelete(ctx context.Context, req *v1.ProductDeleteReq) (res *v1.ProductDeleteRes, err error)
	ProductView(ctx context.Context, req *v1.ProductViewReq) (res *v1.ProductViewRes, err error)
	ProductList(ctx context.Context, req *v1.ProductListReq) (res *v1.ProductListRes, err error)
	RoomCreate(ctx context.Context, req *v1.RoomCreateReq) (res *v1.RoomCreateRes, err error)
	RoomUpdate(ctx context.Context, req *v1.RoomUpdateReq) (res *v1.RoomUpdateRes, err error)
	RoomDelete(ctx context.Context, req *v1.RoomDeleteReq) (res *v1.RoomDeleteRes, err error)
	RoomView(ctx context.Context, req *v1.RoomViewReq) (res *v1.RoomViewRes, err error)
	RoomList(ctx context.Context, req *v1.RoomListReq) (res *v1.RoomListRes, err error)
}
