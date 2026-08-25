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

func (a *sAsset) CategoryCreate(ctx context.Context, req *model.CategoryCreateReq) (err error) {
	if req.Name == "" {
		return gerror.New("请输入分类名称")
	}
	if req.Code == "" {
		return gerror.New("请输入分类编码")
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

	category, err := a.getCategoryByCode(ctx, req.Code)
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

func (a *sAsset) CategoryView(ctx context.Context, req *model.CategoryViewReq) (res *model.CategoryViewRes, err error) {
	return
}

func (a *sAsset) CategoryList(ctx context.Context, req *model.CategoryListReq) (res []*model.CategoryListRes, total int, err error) {
	return
}
