# Task 4 / Step 2: 编写通用 CRUD 辅助方法

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Logic 实现](README.md)

- [ ] **Step 2: 编写通用 CRUD 辅助方法**

在同一文件追加：

```go
func (s *sAsset) create(ctx context.Context, table string, data gdb.Map) (uint64, error) {
	id, err := g.DB().Model(table).Ctx(ctx).Data(data).InsertAndGetId()
	return uint64(id), err
}

func (s *sAsset) update(ctx context.Context, table string, id uint64, data gdb.Map) error {
	_, err := g.DB().Model(table).Ctx(ctx).Where("id", id).Data(data).Update()
	return err
}

func (s *sAsset) detail(ctx context.Context, table string, id uint64) (gdb.Record, error) {
	return g.DB().Model(table).Ctx(ctx).Where("id", id).WhereNull("deleted_at").One()
}

func (s *sAsset) listModel(ctx context.Context, table string, page model.PageInput, filters gdb.Map, out interface{}) (int, error) {
	m := g.DB().Model(table).Ctx(ctx).WhereNull("deleted_at")
	for k, v := range filters {
		switch value := v.(type) {
		case string:
			if strings.TrimSpace(value) != "" {
				m = m.WhereLike(k, "%"+strings.TrimSpace(value)+"%")
			}
		case uint64:
			if value > 0 {
				m = m.Where(k, value)
			}
		case int:
			if value > 0 {
				m = m.Where(k, value)
			}
		default:
			if value != nil {
				m = m.Where(k, value)
			}
		}
	}
	total, err := m.Count()
	if err != nil {
		return 0, err
	}
	err = m.OrderDesc("id").Limit(page.Offset(), page.Limit()).Scan(out)
	return total, err
}

func (s *sAsset) softDelete(ctx context.Context, table string, id uint64) error {
	_, err := g.DB().Model(table).Ctx(ctx).Where("id", id).Data(gdb.Map{
		"deleted_at": gtime.Now(),
	}).Update()
	return err
}
```
