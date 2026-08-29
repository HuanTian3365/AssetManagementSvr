// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"asset_management_svr/internal/model"
	"context"
)

type (
	IAsset interface {
		// BuildingCreate 新建建筑
		BuildingCreate(ctx context.Context, req *model.BuildingCreateReq) (err error)
		// BuildingUpdate 更新建筑
		BuildingUpdate(ctx context.Context, req *model.BuildingUpdateReq) (err error)
		// BuildingDelete 删除建筑
		BuildingDelete(ctx context.Context, req *model.BuildingDeleteReq) (err error)
		// BuildingView 建筑详情
		BuildingView(ctx context.Context, req *model.BuildingViewReq) (res *model.BuildingViewRes, err error)
		// BuildingList 建筑列表
		BuildingList(ctx context.Context, req *model.BuildingListReq) (res []*model.BuildingListRes, total int, err error)
		// CategoryCreate 新建分类
		CategoryCreate(ctx context.Context, req *model.CategoryCreateReq) (err error)
		// CategoryUpdate 更新分类
		CategoryUpdate(ctx context.Context, req *model.CategoryUpdateReq) (err error)
		// CategoryDelete 删除分类
		CategoryDelete(ctx context.Context, req *model.CategoryDeleteReq) (err error)
		// CategoryView 分类详情
		CategoryView(ctx context.Context, req *model.CategoryViewReq) (res *model.CategoryViewRes, err error)
		// CategoryList 分类列表
		CategoryList(ctx context.Context, req *model.CategoryListReq) (res []*model.CategoryListRes, total int, err error)
		// FloorCreate 新建楼层
		FloorCreate(ctx context.Context, req *model.FloorCreateReq) (err error)
		// FloorUpdate 更新楼层
		FloorUpdate(ctx context.Context, req *model.FloorUpdateReq) (err error)
		// FloorDelete 删除楼层
		FloorDelete(ctx context.Context, req *model.FloorDeleteReq) (err error)
		// FloorView 楼层详情
		FloorView(ctx context.Context, req *model.FloorViewReq) (res *model.FloorViewRes, err error)
		// FloorList 楼层列表
		FloorList(ctx context.Context, req *model.FloorListReq) (res []*model.FloorListRes, total int, err error)
		// ProductCreate 新增产品
		ProductCreate(ctx context.Context, req *model.ProductCreateReq) (err error)
		// ProductUpdate 更新产品
		ProductUpdate(ctx context.Context, req *model.ProductUpdateReq) (err error)
		// ProductDelete 删除产品
		ProductDelete(ctx context.Context, req *model.ProductDeleteReq) (err error)
		// ProductView 产品详情
		ProductView(ctx context.Context, req *model.ProductViewReq) (res *model.ProductViewRes, err error)
		// ProductList 产品列表
		ProductList(ctx context.Context, req *model.ProductListReq) (res []*model.ProductListRes, total int, err error)
		RoomCreate(ctx context.Context, req *model.RoomCreateReq) (err error)
		RoomUpdate(ctx context.Context, req *model.RoomUpdateReq) (err error)
		RoomDelete(ctx context.Context, req *model.RoomDeleteReq) (err error)
		RoomView(ctx context.Context, req *model.RoomViewReq) (res *model.RoomViewRes, err error)
		RoomList(ctx context.Context, req *model.RoomListReq) (res []*model.RoomListRes, total int, err error)
	}
)

var (
	localAsset IAsset
)

func Asset() IAsset {
	if localAsset == nil {
		panic("implement not found for interface IAsset, forgot register?")
	}
	return localAsset
}

func RegisterAsset(i IAsset) {
	localAsset = i
}
