# Task 4 / Step 5: 编写位置校验、自动编码、资产入库和转移

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Logic 实现](README.md)

- [ ] **Step 5: 编写位置校验、自动编码、资产入库和转移**

在同一文件追加：

```go
func (s *sAsset) checkFloorBelongsToBuilding(ctx context.Context, floorId uint64, buildingId uint64) error {
	count, err := dao.AssetFloor.Ctx(ctx).
		Where("id", floorId).
		Where("building_id", buildingId).
		WhereNull("deleted_at").
		Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("楼层不属于所选建筑")
	}
	return nil
}

func (s *sAsset) getRoomLocation(ctx context.Context, roomId uint64) (gdb.Record, error) {
	room, err := dao.AssetRoom.Ctx(ctx).Where("id", roomId).WhereNull("deleted_at").One()
	if err != nil {
		return nil, err
	}
	if room.IsEmpty() {
		return nil, gerror.New("房间不存在")
	}
	return room, nil
}

func (s *sAsset) generateAssetCode(ctx context.Context, productId uint64) (string, error) {
	product, err := dao.AssetProduct.Ctx(ctx).Where("id", productId).WhereNull("deleted_at").One()
	if err != nil {
		return "", err
	}
	if product.IsEmpty() {
		return "", gerror.New("产品不存在")
	}
	productCode := product["code"].String()
	dateText := time.Now().Format("20060102")
	prefix := fmt.Sprintf("ASSET-%s-%s-", productCode, dateText)
	count, err := dao.AssetItem.Ctx(ctx).WhereLike("asset_code", prefix+"%").Count()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%03d", prefix, count+1), nil
}

func (s *sAsset) CreateItem(ctx context.Context, in *model.AssetItemCreateInput) (*model.AssetItemCreateOutput, error) {
	out := &model.AssetItemCreateOutput{}
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		assetCode := strings.TrimSpace(in.AssetCode)
		var err error
		if assetCode == "" {
			assetCode, err = s.generateAssetCode(ctx, in.ProductId)
			if err != nil {
				return err
			}
		}
		data := gdb.Map{
			"product_id":      in.ProductId,
			"asset_code":      assetCode,
			"name":            in.Name,
			"status":          in.Status,
			"purchase_date":   in.PurchaseDate,
			"purchase_price":  in.PurchasePrice,
			"remark":          in.Remark,
		}
		if in.RoomId > 0 {
			room, err := s.getRoomLocation(ctx, in.RoomId)
			if err != nil {
				return err
			}
			data["current_building_id"] = room["building_id"].Uint64()
			data["current_floor_id"] = room["floor_id"].Uint64()
			data["current_room_id"] = room["id"].Uint64()
		}
		id, err := tx.Model(dao.AssetItem.Table()).Data(data).InsertAndGetId()
		if err != nil {
			return err
		}
		out.Id = uint64(id)
		out.AssetCode = assetCode
		if in.RoomId > 0 {
			room, err := s.getRoomLocation(ctx, in.RoomId)
			if err != nil {
				return err
			}
			_, err = tx.Model(dao.AssetLocationRecord.Table()).Data(gdb.Map{
				"asset_id":       out.Id,
				"action_type":    model.AssetLocationActionInbound,
				"to_building_id": room["building_id"].Uint64(),
				"to_floor_id":    room["floor_id"].Uint64(),
				"to_room_id":     room["id"].Uint64(),
				"operated_at":    gtime.Now(),
				"remark":         in.Remark,
			}).Insert()
			return err
		}
		return nil
	})
	return out, err
}

func (s *sAsset) UpdateItem(ctx context.Context, in *model.AssetItemUpdateInput) error {
	_, err := dao.AssetItem.Ctx(ctx).Where("id", in.Id).Data(gdb.Map{
		"product_id":      in.ProductId,
		"asset_code":      in.AssetCode,
		"name":            in.Name,
		"status":          in.Status,
		"purchase_date":   in.PurchaseDate,
		"purchase_price":  in.PurchasePrice,
		"remark":          in.Remark,
	}).Update()
	return err
}

func (s *sAsset) DeleteItem(ctx context.Context, id uint64) error {
	return s.softDelete(ctx, dao.AssetItem.Table(), id)
}

func (s *sAsset) DetailItem(ctx context.Context, id uint64) (*model.AssetItemOutput, error) {
	var out *model.AssetItemOutput
	err := dao.AssetItem.Ctx(ctx).Where("id", id).WhereNull("deleted_at").Scan(&out)
	return out, err
}

func (s *sAsset) ListItem(ctx context.Context, in *model.AssetItemListInput) ([]*model.AssetItemOutput, int, error) {
	var list []*model.AssetItemOutput
	total, err := s.listModel(ctx, dao.AssetItem.Table(), in.PageInput, gdb.Map{
		"product_id":      in.ProductId,
		"current_room_id": in.RoomId,
		"name":            in.Name,
		"asset_code":      in.AssetCode,
		"status":          in.Status,
	}, &list)
	return list, total, err
}

func (s *sAsset) InboundItem(ctx context.Context, in *model.AssetItemInboundInput) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		room, err := s.getRoomLocation(ctx, in.RoomId)
		if err != nil {
			return err
		}
		_, err = tx.Model(dao.AssetItem.Table()).Where("id", in.AssetId).Data(gdb.Map{
			"current_building_id": room["building_id"].Uint64(),
			"current_floor_id":    room["floor_id"].Uint64(),
			"current_room_id":     room["id"].Uint64(),
		}).Update()
		if err != nil {
			return err
		}
		_, err = tx.Model(dao.AssetLocationRecord.Table()).Data(gdb.Map{
			"asset_id":       in.AssetId,
			"action_type":    model.AssetLocationActionInbound,
			"to_building_id": room["building_id"].Uint64(),
			"to_floor_id":    room["floor_id"].Uint64(),
			"to_room_id":     room["id"].Uint64(),
			"operated_at":    gtime.Now(),
			"remark":         in.Remark,
		}).Insert()
		return err
	})
}

func (s *sAsset) TransferItem(ctx context.Context, in *model.AssetItemTransferInput) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		item, err := dao.AssetItem.Ctx(ctx).Where("id", in.AssetId).WhereNull("deleted_at").One()
		if err != nil {
			return err
		}
		if item.IsEmpty() {
			return gerror.New("资产不存在")
		}
		if item["current_room_id"].Uint64() == 0 {
			return gerror.New("资产还没有当前位置，不能转移")
		}
		if item["current_room_id"].Uint64() == in.RoomId {
			return gerror.New("目标房间不能等于当前房间")
		}
		room, err := s.getRoomLocation(ctx, in.RoomId)
		if err != nil {
			return err
		}
		_, err = tx.Model(dao.AssetItem.Table()).Where("id", in.AssetId).Data(gdb.Map{
			"current_building_id": room["building_id"].Uint64(),
			"current_floor_id":    room["floor_id"].Uint64(),
			"current_room_id":     room["id"].Uint64(),
		}).Update()
		if err != nil {
			return err
		}
		_, err = tx.Model(dao.AssetLocationRecord.Table()).Data(gdb.Map{
			"asset_id":          in.AssetId,
			"action_type":       model.AssetLocationActionTransfer,
			"from_building_id":  item["current_building_id"].Uint64(),
			"from_floor_id":     item["current_floor_id"].Uint64(),
			"from_room_id":      item["current_room_id"].Uint64(),
			"to_building_id":    room["building_id"].Uint64(),
			"to_floor_id":       room["floor_id"].Uint64(),
			"to_room_id":        room["id"].Uint64(),
			"operated_at":       gtime.Now(),
			"remark":            in.Remark,
		}).Insert()
		return err
	})
}

func (s *sAsset) ListLocationRecord(ctx context.Context, in *model.AssetLocationRecordListInput) ([]*model.AssetLocationRecordOutput, int, error) {
	var list []*model.AssetLocationRecordOutput
	total, err := s.listModel(ctx, dao.AssetLocationRecord.Table(), in.PageInput, gdb.Map{"asset_id": in.AssetId}, &list)
	return list, total, err
}
```
