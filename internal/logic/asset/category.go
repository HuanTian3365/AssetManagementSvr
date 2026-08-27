package asset

import (
	"asset_management_svr/internal/dao"
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (a *sAsset) getCategoryByCode(ctx context.Context, code string) (res *entity.AssetCategory, err error) {
	err = dao.AssetCategory.Ctx(ctx).Where(dao.AssetCategory.Columns().Code, code).Scan(&res)
	return
}

func (a *sAsset) getCategoryById(ctx context.Context, id uint64) (res *entity.AssetCategory, err error) {
	err = dao.AssetCategory.Ctx(ctx).WherePri(id).Scan(&res)
	return
}
func (a *sAsset) getChildCategoryById(ctx context.Context, id uint64) (res *entity.AssetCategory, err error) {
	err = dao.AssetCategory.Ctx(ctx).Where(dao.AssetCategory.Columns().ParentId, id).Scan(&res)
	return
}

// CategoryCreate 新建分类
func (a *sAsset) CategoryCreate(ctx context.Context, req *model.CategoryCreateReq) (err error) {
	if req.Name == "" {
		return gerror.New("请输入分类名称")
	}
	if req.Code == "" {
		return gerror.New("请输入分类编码")
	}

	// 判断父分类是否存在
	if req.ParentId != 0 {
		category, err := a.getCategoryById(ctx, req.ParentId)
		if err != nil {
			return err
		}
		if category == nil {
			return gerror.New("父分类不存在")
		}
	}

	category, err := a.getCategoryByCode(ctx, req.Code)
	if err != nil {
		return err
	}
	if category != nil {
		return gerror.New("分类编码已存在")
	}

	_, err = dao.AssetCategory.Ctx(ctx).Fields(&model.CategoryCreateReq{}).Data(req).Insert()
	return
}

// CategoryUpdate 更新分类
func (a *sAsset) CategoryUpdate(ctx context.Context, req *model.CategoryUpdateReq) (err error) {
	if req.Id == 0 {
		return gerror.New("请输入ID")
	}
	if req.Name == "" {
		return gerror.New("请输入分类名称")
	}
	if req.Code == "" {
		return gerror.New("请输入分类编码")
	}

	// 检查修改分类是否存在
	category, err := a.getCategoryById(ctx, req.Id)
	if err != nil {
		return err
	}
	if category == nil {
		return gerror.New("分类不存在")
	}

	// 判断父分类是否存在
	if req.ParentId != 0 {
		category, err := a.getCategoryById(ctx, req.ParentId)
		if err != nil {
			return err
		}
		if category == nil {
			return gerror.New("父分类不存在")
		}
	}

	if req.Id == req.ParentId {
		return gerror.New("父分类不能为该分类")
	}

	category, err = a.getCategoryByCode(ctx, req.Code)
	if err != nil {
		return err
	}
	if category != nil && category.Id != req.Id {
		return gerror.New("分类编码已存在")
	}

	col := dao.AssetCategory.Columns()
	_, err = dao.AssetCategory.Ctx(ctx).WherePri(req.Id).Fields(col.ParentId, col.Name, col.Code, col.Sort, col.Remark).Data(req).Update()
	return
}

// CategoryDelete 删除分类
func (a *sAsset) CategoryDelete(ctx context.Context, req *model.CategoryDeleteReq) (err error) {
	if req.Id == 0 {
		return gerror.New("请输入ID")
	}

	category, err := a.getCategoryById(ctx, req.Id)
	if err != nil {
		return err
	}
	if category == nil {
		return gerror.New("分类不存在")
	}
	category, err = a.getChildCategoryById(ctx, req.Id)
	if err != nil {
		return err
	}
	if category != nil {
		return gerror.New("该分类存在子分类,无法删除")
	}

	_, err = dao.AssetCategory.Ctx(ctx).Where(dao.AssetCategory.Columns().Id, req.Id).Delete()
	return
}

// CategoryView 分类详情
func (a *sAsset) CategoryView(ctx context.Context, req *model.CategoryViewReq) (res *model.CategoryViewRes, err error) {
	data, err := a.getCategoryById(ctx, req.Id)
	if err != nil {
		return
	}
	if data == nil {
		return nil, gerror.New("分类不存在")
	}
	return &model.CategoryViewRes{
		Id:        data.Id,
		ParentId:  data.ParentId,
		Name:      data.Name,
		Code:      data.Code,
		Sort:      data.Sort,
		Remark:    data.Remark,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}, nil
}

// CategoryList 分类列表
func (a *sAsset) CategoryList(ctx context.Context, req *model.CategoryListReq) (res []*model.CategoryListRes, total int, err error) {
	mod := dao.AssetCategory.Ctx(ctx)
	col := dao.AssetCategory.Columns()

	if req.Id != 0 {
		mod = mod.WherePri(req.Id)
	}
	if req.ParentId != 0 {
		mod = mod.Where(col.ParentId, req.ParentId)
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
