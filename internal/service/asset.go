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
