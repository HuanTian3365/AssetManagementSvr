# Task 4 / Step 3: 编写基础资料 CRUD

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Logic 实现](README.md)

- [ ] **Step 3: 编写基础资料 CRUD**

在同一文件追加：

```go
func (s *sAsset) CreateBuilding(ctx context.Context, data gdb.Map) (uint64, error) {
	return s.create(ctx, dao.AssetBuilding.Table(), data)
}

func (s *sAsset) UpdateBuilding(ctx context.Context, id uint64, data gdb.Map) error {
	return s.update(ctx, dao.AssetBuilding.Table(), id, data)
}

func (s *sAsset) DeleteBuilding(ctx context.Context, id uint64) error {
	count, err := dao.AssetFloor.Ctx(ctx).Where("building_id", id).WhereNull("deleted_at").Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("建筑下存在楼层，不能删除")
	}
	return s.softDelete(ctx, dao.AssetBuilding.Table(), id)
}

func (s *sAsset) DetailBuilding(ctx context.Context, id uint64) (gdb.Record, error) {
	return s.detail(ctx, dao.AssetBuilding.Table(), id)
}

func (s *sAsset) ListBuilding(ctx context.Context, page model.PageInput, filters gdb.Map) (gdb.Result, int, error) {
	return s.list(ctx, dao.AssetBuilding.Table(), page, filters)
}

func (s *sAsset) CreateFloor(ctx context.Context, data gdb.Map) (uint64, error) {
	return s.create(ctx, dao.AssetFloor.Table(), data)
}

func (s *sAsset) UpdateFloor(ctx context.Context, id uint64, data gdb.Map) error {
	return s.update(ctx, dao.AssetFloor.Table(), id, data)
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

func (s *sAsset) DetailFloor(ctx context.Context, id uint64) (gdb.Record, error) {
	return s.detail(ctx, dao.AssetFloor.Table(), id)
}

func (s *sAsset) ListFloor(ctx context.Context, page model.PageInput, filters gdb.Map) (gdb.Result, int, error) {
	return s.list(ctx, dao.AssetFloor.Table(), page, filters)
}

func (s *sAsset) CreateRoom(ctx context.Context, data gdb.Map) (uint64, error) {
	if err := s.checkFloorBelongsToBuilding(ctx, data["floor_id"].(uint64), data["building_id"].(uint64)); err != nil {
		return 0, err
	}
	return s.create(ctx, dao.AssetRoom.Table(), data)
}

func (s *sAsset) UpdateRoom(ctx context.Context, id uint64, data gdb.Map) error {
	if err := s.checkFloorBelongsToBuilding(ctx, data["floor_id"].(uint64), data["building_id"].(uint64)); err != nil {
		return err
	}
	return s.update(ctx, dao.AssetRoom.Table(), id, data)
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

func (s *sAsset) DetailRoom(ctx context.Context, id uint64) (gdb.Record, error) {
	return s.detail(ctx, dao.AssetRoom.Table(), id)
}

func (s *sAsset) ListRoom(ctx context.Context, page model.PageInput, filters gdb.Map) (gdb.Result, int, error) {
	return s.list(ctx, dao.AssetRoom.Table(), page, filters)
}
```
