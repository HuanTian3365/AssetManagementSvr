package asset

import (
	"asset_management_svr/internal/dao"
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 根据Code获取建筑 为确保唯一,不忽略软删除数据
func (a *sAsset) getBuildingByCode(ctx context.Context, code string) (res *entity.AssetBuilding, err error) {
	err = dao.AssetBuilding.Ctx(ctx).Unscoped().
		Where(dao.AssetBuilding.Columns().Code, code).
		Scan(&res)
	return
}

// 根据Id获取建筑
func (a *sAsset) getBuildingById(ctx context.Context, id uint64) (res *entity.AssetBuilding, err error) {
	err = dao.AssetBuilding.Ctx(ctx).
		WherePri(id).
		Scan(&res)
	return
}

// BuildingCreate 新建建筑
func (a *sAsset) BuildingCreate(ctx context.Context, req *model.BuildingCreateReq) (err error) {
	// 非空判断
	if req.Name == "" {
		return gerror.New("请输入建筑名称")
	}
	if req.Code == "" {
		return gerror.New("请输入建筑编号")
	}

	// 判断建筑编号是否已经存在
	data, err := a.getBuildingByCode(ctx, req.Code)
	if err != nil {
		return err
	}
	if data != nil {
		err = gerror.New("建筑编号已存在")
		return
	}

	// 新增数据
	_, err = dao.AssetBuilding.Ctx(ctx).Fields(model.BuildingCreateReq{}).Data(req).Insert()
	return err
}

// BuildingUpdate 更新建筑
func (a *sAsset) BuildingUpdate(ctx context.Context, req *model.BuildingUpdateReq) (err error) {
	// 非空判断
	if req.Id == 0 {
		return gerror.New("请输入ID")
	}
	if req.Name == "" {
		return gerror.New("请输入建筑名称")
	}
	if req.Code == "" {
		return gerror.New("请输入建筑编号")
	}

	// 判断建筑ID是否存在
	data, err := a.getBuildingById(ctx, req.Id)
	if err != nil {
		return err
	}
	if data == nil {
		return gerror.New("建筑不存在")
	}
	// 判断建筑编号是否存在
	data, err = a.getBuildingByCode(ctx, req.Code)
	if err != nil {
		return err
	}
	if data != nil && req.Id != data.Id {
		return gerror.New("重复的建筑编号")
	}

	// 更新数据
	col := dao.AssetBuilding.Columns()
	_, err = dao.AssetBuilding.Ctx(ctx).WherePri(req.Id).Fields(col.Name, col.Code, col.Remark, col.Address).Data(req).Update()

	return err
}

// BuildingDelete 删除建筑
func (a *sAsset) BuildingDelete(ctx context.Context, req *model.BuildingDeleteReq) (err error) {
	// 非空判断
	if req.Id == 0 {
		return gerror.New("请输入ID")
	}

	// 判断建筑是否存在
	data, err := a.getBuildingById(ctx, req.Id)
	if err != nil {
		return err
	}
	if data == nil {
		return gerror.New("建筑不存在")
	}

	// 判断建筑下是否存在楼层
	floor, err := a.getFloorByBuild(ctx, req.Id)
	if err != nil {
		return err
	}
	if floor != nil {
		return gerror.New("建筑下存在楼层,无法删除")
	}

	// 删除数据
	_, err = dao.AssetBuilding.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return err
	}

	return nil
}

// BuildingView 建筑详情
func (a *sAsset) BuildingView(ctx context.Context, req *model.BuildingViewReq) (res *model.BuildingViewRes, err error) {
	data, err := a.getBuildingById(ctx, req.Id)
	if err != nil {
		return
	}
	if data == nil {
		err = gerror.New("建筑不存在")
		return
	}
	return &model.BuildingViewRes{
		Id:       data.Id,
		Name:     data.Name,
		Code:     data.Code,
		Address:  data.Address,
		Remark:   data.Remark,
		CreateAt: data.CreatedAt,
		UpdateAt: data.UpdatedAt,
	}, nil
}

// BuildingList 建筑列表
func (a *sAsset) BuildingList(ctx context.Context, req *model.BuildingListReq) (res []*model.BuildingListRes, total int, err error) {
	mod := dao.AssetBuilding.Ctx(ctx)
	col := dao.AssetBuilding.Columns()

	// 条件查询
	if req.Id != 0 {
		mod = mod.WherePri(req.Id)
	}
	if req.Name != "" {
		mod = mod.WhereLike(col.Name, "%"+req.Name+"%")
	}
	if req.Code != "" {
		mod = mod.Where(col.Code, req.Code)
	}
	if req.Address != "" {
		mod = mod.WhereLike(col.Address, "%"+req.Address+"%")
	}
	if len(req.CreatedAt) == 2 {
		mod = mod.WhereBetween(col.CreatedAt, req.CreatedAt[0], req.CreatedAt[1])
	}

	// 分页
	if req.FullSize == false {
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 20
		}
		mod = mod.Page(req.Page, req.PageSize)
	}

	// 排序
	mod = mod.OrderDesc(col.Id)

	err = mod.ScanAndCount(&res, &total, false)
	if err != nil {
		return nil, 0, err
	}
	return
}
