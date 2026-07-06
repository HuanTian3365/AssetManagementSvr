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
		// FloorCreate 新建楼层
		FloorCreate(ctx context.Context, req *model.FloorCreateReq) (err error)
		// FloorUpdate 更新楼层
		FloorUpdate(ctx context.Context, req *model.FloorUpdateReq) (err error)
		// FloorDelete 删除楼层
		FloorDelete(ctx context.Context, req *model.FloorDeleteReq) (err error)
		// FloorView 楼层详情
		FloorView(ctx context.Context, req *model.FloorViewReq) (res *model.FloorViewRes, err error)
		FloorList(ctx context.Context, req *model.FloorListReq) (res []*model.FloorListRes, total int, err error)
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
