package asset

import (
	"asset_management_svr/internal/dao"
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 根据建筑获取楼层
func (a *sAsset) getFloorByBuild(ctx context.Context, buildingId uint64) (res *entity.AssetFloor, err error) {
	err = dao.AssetFloor.Ctx(ctx).Where(dao.AssetFloor.Columns().BuildingId, buildingId).Scan(&res)
	return

}

// 根据编号获取楼层
func (a *sAsset) getFloorByCode(ctx context.Context, code string) (res *entity.AssetFloor, err error) {
	err = dao.AssetFloor.Ctx(ctx).Unscoped().Where(dao.AssetFloor.Columns().Code, code).Scan(&res)
	return
}

// 根据建筑和楼层序号获取楼层
func (a *sAsset) getFloorByNo(ctx context.Context, buildingId uint64, floorNo int) (res *entity.AssetFloor, err error) {
	err = dao.AssetFloor.Ctx(ctx).Unscoped().Where(dao.AssetFloor.Columns().BuildingId, buildingId).Where(dao.AssetFloor.Columns().FloorNo, floorNo).Scan(&res)
	return
}

// 根据ID获取楼层
func (a *sAsset) getFloorById(ctx context.Context, id uint64) (res *entity.AssetFloor, err error) {
	err = dao.AssetFloor.Ctx(ctx).WherePri(id).Scan(&res)
	return
}

// FloorCreate 新建楼层
func (a *sAsset) FloorCreate(ctx context.Context, req *model.FloorCreateReq) (err error) {
	if req.Name == "" {
		return gerror.New("请输入楼层名称")
	}
	if req.Code == "" {
		return gerror.New("请输入楼层编号")
	}
	if req.BuildingId == 0 {
		return gerror.New("请输入建筑编号")
	}
	if req.FloorNo == 0 {
		return gerror.New("请输入楼层序号")
	}

	build, err := a.getBuildingById(ctx, req.BuildingId)
	if err != nil {
		return err
	}
	if build == nil {
		return gerror.New("建筑编号不存在")
	}

	floor, err := a.getFloorByCode(ctx, req.Code)
	if err != nil {
		return err
	}
	if floor != nil {
		return gerror.New("楼层编号已存在")
	}

	floor, err = a.getFloorByNo(ctx, req.BuildingId, req.FloorNo)
	if err != nil {
		return err
	}
	if floor != nil {
		return gerror.New("楼层序号已存在")
	}

	_, err = dao.AssetFloor.Ctx(ctx).Fields(model.FloorCreateReq{}).Data(req).Insert()
	return
}

// FloorUpdate 更新楼层
func (a *sAsset) FloorUpdate(ctx context.Context, req *model.FloorUpdateReq) (err error) {
	if req.Id == 0 {
		return gerror.New("请输入ID")
	}
	if req.Name == "" {
		return gerror.New("请输入楼层名称")
	}
	if req.Code == "" {
		return gerror.New("请输入楼层编号")
	}
	if req.BuildingId == 0 {
		return gerror.New("请输入建筑编号")
	}
	if req.FloorNo == 0 {
		return gerror.New("请输入楼层序号")
	}

	floor, err := a.getFloorById(ctx, req.Id)
	if err != nil {
		return err
	}
	if floor == nil {
		return gerror.New("楼层不存在")
	}

	build, err := a.getBuildingById(ctx, req.BuildingId)
	if err != nil {
		return err
	}
	if build == nil {
		return gerror.New("建筑不存在")
	}

	floor, err = a.getFloorByCode(ctx, req.Code)
	if err != nil {
		return err
	}
	if floor != nil && floor.Id != req.Id {
		return gerror.New("楼层编号已存在")
	}

	floor, err = a.getFloorByNo(ctx, req.BuildingId, req.FloorNo)
	if err != nil {
		return err
	}
	if floor != nil && floor.Id != req.Id {
		return gerror.New("楼层序号已存在")
	}

	col := dao.AssetFloor.Columns()
	_, err = dao.AssetFloor.Ctx(ctx).WherePri(req.Id).Fields(col.BuildingId, col.Name, col.Code, col.FloorNo, col.Remark).Data(req).Update()
	return
}

// FloorDelete 删除楼层
func (a *sAsset) FloorDelete(ctx context.Context, req *model.FloorDeleteReq) (err error) {
	if req.Id == 0 {
		return gerror.New("请输入ID")
	}

	floor, err := a.getFloorById(ctx, req.Id)
	if err != nil {
		return err
	}
	if floor == nil {
		return gerror.New("楼层不存在")
	}
	room, err := a.getRoomByFloorId(ctx, req.Id)
	if err != nil {
		return err
	}
	if room != nil {
		return gerror.New("楼层下存在房间,无法删除")
	}

	_, err = dao.AssetFloor.Ctx(ctx).WherePri(req.Id).Delete()
	return
}

// FloorView 楼层详情
func (a *sAsset) FloorView(ctx context.Context, req *model.FloorViewReq) (res *model.FloorViewRes, err error) {
	floor, err := a.getFloorById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if floor == nil {
		return nil, gerror.New("楼层不存在")
	}
	return &model.FloorViewRes{
		Id:         floor.Id,
		BuildingId: floor.BuildingId,
		Name:       floor.Name,
		Code:       floor.Code,
		FloorNo:    floor.FloorNo,
		Remark:     floor.Remark,
		CreatedAt:  floor.CreatedAt,
		UpdatedAt:  floor.UpdatedAt,
	}, nil
}

// FloorList 楼层列表
func (a *sAsset) FloorList(ctx context.Context, req *model.FloorListReq) (res []*model.FloorListRes, total int, err error) {
	mod := dao.AssetFloor.Ctx(ctx)
	col := dao.AssetFloor.Columns()

	if req.Id != 0 {
		mod = mod.WherePri(req.Id)
	}
	if req.BuildingId != 0 {
		mod = mod.Where(col.BuildingId, req.BuildingId)
	}
	if req.Name != "" {
		mod = mod.WhereLike(col.Name, "%"+req.Name+"%")
	}
	if req.Code != "" {
		mod = mod.Where(col.Code, req.Code)
	}
	if len(req.CreatedAt) == 2 {
		mod = mod.WhereBetween(col.CreatedAt, req.CreatedAt[0], req.CreatedAt[1])
	}

	if req.FullSize == false {
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 20
		}
		mod = mod.Page(req.Page, req.PageSize)
	}

	mod = mod.OrderDesc(col.Id)

	err = mod.ScanAndCount(&res, &total, false)
	if err != nil {
		return nil, 0, err
	}
	return
}
