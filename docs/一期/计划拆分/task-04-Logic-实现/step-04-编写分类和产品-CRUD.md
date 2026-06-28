# Task 4 / Step 4: 编写分类和产品 CRUD

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Logic 实现](README.md)

- [ ] **Step 4: 编写分类和产品 CRUD**

在同一文件追加：

```go
func (s *sAsset) CreateCategory(ctx context.Context, in *model.AssetCategoryCreateInput) (uint64, error) {
	return s.create(ctx, dao.AssetCategory.Table(), gdb.Map{
		"parent_id": in.ParentId,
		"name":      in.Name,
		"code":      in.Code,
		"sort":      in.Sort,
		"remark":    in.Remark,
	})
}

func (s *sAsset) UpdateCategory(ctx context.Context, in *model.AssetCategoryUpdateInput) error {
	if in.ParentId == in.Id {
		return gerror.New("父级分类不能选择自己")
	}
	return s.update(ctx, dao.AssetCategory.Table(), in.Id, gdb.Map{
		"parent_id": in.ParentId,
		"name":      in.Name,
		"code":      in.Code,
		"sort":      in.Sort,
		"remark":    in.Remark,
	})
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

func (s *sAsset) DetailCategory(ctx context.Context, id uint64) (*model.AssetCategoryOutput, error) {
	var out *model.AssetCategoryOutput
	err := dao.AssetCategory.Ctx(ctx).Where("id", id).WhereNull("deleted_at").Scan(&out)
	return out, err
}

func (s *sAsset) ListCategory(ctx context.Context, in *model.AssetCategoryListInput) ([]*model.AssetCategoryOutput, int, error) {
	var list []*model.AssetCategoryOutput
	total, err := s.listModel(ctx, dao.AssetCategory.Table(), in.PageInput, gdb.Map{
		"parent_id": in.ParentId,
		"name":      in.Name,
		"code":      in.Code,
	}, &list)
	return list, total, err
}

func (s *sAsset) CreateProduct(ctx context.Context, in *model.AssetProductCreateInput) (uint64, error) {
	return s.create(ctx, dao.AssetProduct.Table(), gdb.Map{
		"category_id": in.CategoryId,
		"name":        in.Name,
		"code":        in.Code,
		"brand":       in.Brand,
		"model":       in.Model,
		"unit":        in.Unit,
		"remark":      in.Remark,
	})
}

func (s *sAsset) UpdateProduct(ctx context.Context, in *model.AssetProductUpdateInput) error {
	return s.update(ctx, dao.AssetProduct.Table(), in.Id, gdb.Map{
		"category_id": in.CategoryId,
		"name":        in.Name,
		"code":        in.Code,
		"brand":       in.Brand,
		"model":       in.Model,
		"unit":        in.Unit,
		"remark":      in.Remark,
	})
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

func (s *sAsset) DetailProduct(ctx context.Context, id uint64) (*model.AssetProductOutput, error) {
	var out *model.AssetProductOutput
	err := dao.AssetProduct.Ctx(ctx).Where("id", id).WhereNull("deleted_at").Scan(&out)
	return out, err
}

func (s *sAsset) ListProduct(ctx context.Context, in *model.AssetProductListInput) ([]*model.AssetProductOutput, int, error) {
	var list []*model.AssetProductOutput
	total, err := s.listModel(ctx, dao.AssetProduct.Table(), in.PageInput, gdb.Map{
		"category_id": in.CategoryId,
		"name":        in.Name,
		"code":        in.Code,
	}, &list)
	return list, total, err
}
```
