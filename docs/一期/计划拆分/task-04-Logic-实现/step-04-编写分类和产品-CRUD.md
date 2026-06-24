# Task 4 / Step 4: 编写分类和产品 CRUD

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Logic 实现](README.md)

- [ ] **Step 4: 编写分类和产品 CRUD**

在同一文件追加：

```go
func (s *sAsset) CreateCategory(ctx context.Context, data gdb.Map) (uint64, error) {
	return s.create(ctx, dao.AssetCategory.Table(), data)
}

func (s *sAsset) UpdateCategory(ctx context.Context, id uint64, data gdb.Map) error {
	if parentId, ok := data["parent_id"].(uint64); ok && parentId == id {
		return gerror.New("父级分类不能选择自己")
	}
	return s.update(ctx, dao.AssetCategory.Table(), id, data)
}

func (s *sAsset) DeleteCategory(ctx context.Context, id uint64) error {
	childCount, err := dao.AssetCategory.Ctx(ctx).Where("parent_id", id).WhereNull("deleted_at").Count()
	if err != nil {
		return err
	}
	if childCount > 0 {
		return gerror.New("分类下存在子分类，不能删除")
	}
	productCount, err := dao.AssetProduct.Ctx(ctx).Where("category_id", id).WhereNull("deleted_at").Count()
	if err != nil {
		return err
	}
	if productCount > 0 {
		return gerror.New("分类下存在产品，不能删除")
	}
	return s.softDelete(ctx, dao.AssetCategory.Table(), id)
}

func (s *sAsset) DetailCategory(ctx context.Context, id uint64) (gdb.Record, error) {
	return s.detail(ctx, dao.AssetCategory.Table(), id)
}

func (s *sAsset) ListCategory(ctx context.Context, page model.PageInput, filters gdb.Map) (gdb.Result, int, error) {
	return s.list(ctx, dao.AssetCategory.Table(), page, filters)
}

func (s *sAsset) CreateProduct(ctx context.Context, data gdb.Map) (uint64, error) {
	return s.create(ctx, dao.AssetProduct.Table(), data)
}

func (s *sAsset) UpdateProduct(ctx context.Context, id uint64, data gdb.Map) error {
	return s.update(ctx, dao.AssetProduct.Table(), id, data)
}

func (s *sAsset) DeleteProduct(ctx context.Context, id uint64) error {
	count, err := dao.AssetItem.Ctx(ctx).Where("product_id", id).WhereNull("deleted_at").Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("产品已被资产引用，不能删除")
	}
	return s.softDelete(ctx, dao.AssetProduct.Table(), id)
}

func (s *sAsset) DetailProduct(ctx context.Context, id uint64) (gdb.Record, error) {
	return s.detail(ctx, dao.AssetProduct.Table(), id)
}

func (s *sAsset) ListProduct(ctx context.Context, page model.PageInput, filters gdb.Map) (gdb.Result, int, error) {
	return s.list(ctx, dao.AssetProduct.Table(), page, filters)
}
```
