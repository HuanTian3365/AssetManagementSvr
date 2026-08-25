package asset

import (
	"asset_management_svr/internal/dao"
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (a *sAsset) getRoomByCode(ctx context.Context, code string) (res *entity.AssetRoom, err error) {
	err = dao.AssetRoom.Ctx(ctx).Where(dao.AssetRoom.Columns().Code, code).Scan(&res)
	return
}
func (a *sAsset) getRoomByNo(ctx context.Context, buildingId uint64, floorId uint64, roomNo string) (res *entity.AssetRoom, err error) {
	err = dao.AssetRoom.Ctx(ctx).Unscoped().
		Where(dao.AssetRoom.Columns().BuildingId, buildingId).
		Where(dao.AssetRoom.Columns().FloorId, floorId).
		Where(dao.AssetRoom.Columns().RoomNo, roomNo).Scan(&res)
	return
}
func (a *sAsset) getRoomById(ctx context.Context, roomId uint64) (res *entity.AssetRoom, err error) {
	err = dao.AssetRoom.Ctx(ctx).WherePri(roomId).Scan(&res)
	return
}
func (a *sAsset) getRoomByFloorId(ctx context.Context, floorId uint64) (res *entity.AssetRoom, err error) {
	err = dao.AssetRoom.Ctx(ctx).Where(dao.AssetRoom.Columns().FloorId, floorId).Scan(&res)
	return
}

func (a *sAsset) RoomCreate(ctx context.Context, req *model.RoomCreateReq) (err error) {
	if req.BuildingId == 0 {
		return gerror.New("请输入建筑编码")
	}
	if req.FloorId == 0 {
		return gerror.New("请输入楼层编码")
	}
	if req.Name == "" {
		return gerror.New("请输入房间名称")
	}
	if req.Code == "" {
		return gerror.New("请输入房间编号")
	}
	if req.RoomNo == "" {
		return gerror.New("请输入房间号")
	}

	room, err := a.getRoomByCode(ctx, req.Code)
	if err != nil {
		return err
	}
	if room != nil {
		return gerror.New("房间编号已存在")
	}
	build, err := a.getBuildingById(ctx, req.BuildingId)
	if err != nil {
		return err
	}
	if build == nil {
		return gerror.New("建筑不存在")
	}
	floor, err := a.getFloorById(ctx, req.FloorId)
	if err != nil {
		return err
	}
	if floor == nil {
		return gerror.New("楼层不存在")
	}
	if floor.BuildingId != req.BuildingId {
		return gerror.New("楼层不在该建筑下")
	}
	room, err = a.getRoomByNo(ctx, req.BuildingId, req.FloorId, req.RoomNo)
	if err != nil {
		return err
	}
	if room != nil {
		return gerror.New("房间号已存在")
	}

	_, err = dao.AssetRoom.Ctx(ctx).Fields(&model.RoomCreateReq{}).Data(req).Insert()
	return
}

func (a *sAsset) RoomUpdate(ctx context.Context, req *model.RoomUpdateReq) (err error) {
	if req.Id == 0 {
		return gerror.New("请输入房间ID")
	}
	if req.BuildingId == 0 {
		return gerror.New("请输入建筑编码")
	}
	if req.FloorId == 0 {
		return gerror.New("请输入楼层编码")
	}
	if req.Name == "" {
		return gerror.New("请输入房间名称")
	}
	if req.Code == "" {
		return gerror.New("请输入房间编号")
	}
	if req.RoomNo == "" {
		return gerror.New("请输入房间号")
	}

	room, err := a.getRoomById(ctx, req.Id)
	if err != nil {
		return err
	}
	if room == nil {
		return gerror.New("房间不存在")
	}
	room, err = a.getRoomByCode(ctx, req.Code)
	if err != nil {
		return err
	}
	if room != nil && room.Id != req.Id {
		return gerror.New("房间编号已存在")
	}
	build, err := a.getBuildingById(ctx, req.BuildingId)
	if err != nil {
		return err
	}
	if build == nil {
		return gerror.New("建筑不存在")
	}
	floor, err := a.getFloorById(ctx, req.FloorId)
	if err != nil {
		return err
	}
	if floor == nil {
		return gerror.New("楼层不存在")
	}
	if floor.BuildingId != req.BuildingId {
		return gerror.New("楼层不在该建筑下")
	}
	room, err = a.getRoomByNo(ctx, req.BuildingId, req.FloorId, req.RoomNo)
	if err != nil {
		return err
	}
	if room != nil && req.Id != room.Id {
		return gerror.New("房间号已存在")
	}

	col := dao.AssetRoom.Columns()
	_, err = dao.AssetRoom.Ctx(ctx).WherePri(req.Id).Fields(col.BuildingId, col.FloorId, col.Name, col.Code, col.RoomNo, col.Remark).Data(req).Update()
	return
}

func (a *sAsset) RoomDelete(ctx context.Context, req *model.RoomDeleteReq) (err error) {
	if req.Id == 0 {
		return gerror.New("请输入房间ID")
	}
	room, err := a.getRoomById(ctx, req.Id)
	if err != nil {
		return err
	}
	if room == nil {
		return gerror.New("房间不存在")
	}
	// TODO: 等关联表
	_, err = dao.AssetRoom.Ctx(ctx).WherePri(req.Id).Delete()
	return
}

func (a *sAsset) RoomView(ctx context.Context, req *model.RoomViewReq) (res *model.RoomViewRes, err error) {
	if req.Id == 0 {
		return nil, gerror.New("请输入房间ID")
	}
	room, err := a.getRoomById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, gerror.New("房间不存在")
	}
	return &model.RoomViewRes{
		Id:         room.Id,
		BuildingId: room.BuildingId,
		FloorId:    room.FloorId,
		Name:       room.Name,
		Code:       room.Code,
		RoomNo:     room.RoomNo,
		Remark:     room.Remark,
		CreatedAt:  room.CreatedAt,
		UpdatedAt:  room.UpdatedAt,
	}, nil
}

func (a *sAsset) RoomList(ctx context.Context, req *model.RoomListReq) (res []*model.RoomListRes, total int, err error) {
	col := dao.AssetRoom.Columns()
	mod := dao.AssetRoom.Ctx(ctx)

	if req.Id != 0 {
		mod = mod.WherePri(req.Id)
	}
	if req.BuildingId != 0 {
		mod = mod.Where(col.BuildingId, req.BuildingId)
	}
	if req.FloorId != 0 {
		mod = mod.Where(col.FloorId, req.FloorId)
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
