# Task 4 / Step 3: 编写基础资料 CRUD

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Logic 实现](README.md)

- [ ] **Step 3: 编写基础资料 CRUD**

在同一文件追加：

```go
func (s *sAsset) CreateBuilding(ctx context.Context, in *model.AssetBuildingCreateInput) (uint64, error) {
	return s.create(ctx, dao.AssetBuilding.Table(), gdb.Map{
		"name":    in.Name,
		"code":    in.Code,
		"address": in.Address,
		"remark":  in.Remark,
	})
}

func (s *sAsset) UpdateBuilding(ctx context.Context, in *model.AssetBuildingUpdateInput) error {
	return s.update(ctx, dao.AssetBuilding.Table(), in.Id, gdb.Map{
		"name":    in.Name,
		"code":    in.Code,
		"address": in.Address,
		"remark":  in.Remark,
	})
}

func (s *sAsset) DeleteBuilding(ctx context.Context, id uint64) error {
	return s.softDelete(ctx, dao.AssetBuilding.Table(), id)
}

func (s *sAsset) BuildingView(ctx context.Context, id uint64) (*model.AssetBuildingOutput, error) {
	var out *model.AssetBuildingOutput
	err := dao.AssetBuilding.Ctx(ctx).Where("id", id).WhereNull("deleted_at").Scan(&out)
	return out, err
}

func (s *sAsset) ListBuilding(ctx context.Context, in *model.AssetBuildingListInput) ([]*model.AssetBuildingOutput, int, error) {
	var list []*model.AssetBuildingOutput
	total, err := s.listModel(ctx, dao.AssetBuilding.Table(), in.PageInput, gdb.Map{
		"name": in.Name,
		"code": in.Code,
	}, &list)
	return list, total, err
}

func (s *sAsset) CreateFloor(ctx context.Context, in *model.AssetFloorCreateInput) (uint64, error) {
	return s.create(ctx, dao.AssetFloor.Table(), gdb.Map{
		"building_id": in.BuildingId,
		"name":        in.Name,
		"code":        in.Code,
		"floor_no":    in.FloorNo,
		"remark":      in.Remark,
	})
}

func (s *sAsset) UpdateFloor(ctx context.Context, in *model.AssetFloorUpdateInput) error {
	return s.update(ctx, dao.AssetFloor.Table(), in.Id, gdb.Map{
		"building_id": in.BuildingId,
		"name":        in.Name,
		"code":        in.Code,
		"floor_no":    in.FloorNo,
		"remark":      in.Remark,
	})
}

func (s *sAsset) DeleteFloor(ctx context.Context, id uint64) error {
	count, err := dao.AssetRoom.Ctx(ctx).Where("floor_id", id).WhereNull("deleted_at").Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("楼层下存在房间，不能删除")
	}
	return s.softDelete(ctx, dao.AssetFloor.Table(), id)
}

func (s *sAsset) DetailFloor(ctx context.Context, id uint64) (*model.AssetFloorOutput, error) {
	var out *model.AssetFloorOutput
	err := dao.AssetFloor.Ctx(ctx).Where("id", id).WhereNull("deleted_at").Scan(&out)
	return out, err
}

func (s *sAsset) ListFloor(ctx context.Context, in *model.AssetFloorListInput) ([]*model.AssetFloorOutput, int, error) {
	var list []*model.AssetFloorOutput
	total, err := s.listModel(ctx, dao.AssetFloor.Table(), in.PageInput, gdb.Map{
		"building_id": in.BuildingId,
		"name":        in.Name,
		"code":        in.Code,
	}, &list)
	return list, total, err
}

func (s *sAsset) CreateRoom(ctx context.Context, in *model.AssetRoomCreateInput) (uint64, error) {
	if err := s.checkFloorBelongsToBuilding(ctx, in.FloorId, in.BuildingId); err != nil {
		return 0, err
	}
	return s.create(ctx, dao.AssetRoom.Table(), gdb.Map{
		"building_id": in.BuildingId,
		"floor_id":    in.FloorId,
		"name":        in.Name,
		"code":        in.Code,
		"room_no":     in.RoomNo,
		"remark":      in.Remark,
	})
}

func (s *sAsset) UpdateRoom(ctx context.Context, in *model.AssetRoomUpdateInput) error {
	if err := s.checkFloorBelongsToBuilding(ctx, in.FloorId, in.BuildingId); err != nil {
		return err
	}
	return s.update(ctx, dao.AssetRoom.Table(), in.Id, gdb.Map{
		"building_id": in.BuildingId,
		"floor_id":    in.FloorId,
		"name":        in.Name,
		"code":        in.Code,
		"room_no":     in.RoomNo,
		"remark":      in.Remark,
	})
}

func (s *sAsset) DeleteRoom(ctx context.Context, id uint64) error {
	count, err := dao.AssetItem.Ctx(ctx).Where("current_room_id", id).WhereNull("deleted_at").Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("房间已被资产引用，不能删除")
	}
	return s.softDelete(ctx, dao.AssetRoom.Table(), id)
}

func (s *sAsset) DetailRoom(ctx context.Context, id uint64) (*model.AssetRoomOutput, error) {
	var out *model.AssetRoomOutput
	err := dao.AssetRoom.Ctx(ctx).Where("id", id).WhereNull("deleted_at").Scan(&out)
	return out, err
}

func (s *sAsset) ListRoom(ctx context.Context, in *model.AssetRoomListInput) ([]*model.AssetRoomOutput, int, error) {
	var list []*model.AssetRoomOutput
	total, err := s.listModel(ctx, dao.AssetRoom.Table(), in.PageInput, gdb.Map{
		"building_id": in.BuildingId,
		"floor_id":    in.FloorId,
		"name":        in.Name,
		"code":        in.Code,
	}, &list)
	return list, total, err
}
```
