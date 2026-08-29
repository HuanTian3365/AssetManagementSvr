package asset

import (
	"asset_management_svr/internal/dao"
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (a *sAsset) getProductById(ctx context.Context, id uint64) (res *entity.AssetProduct, err error) {
	err = dao.AssetProduct.Ctx(ctx).WherePri(id).Scan(&res)
	return
}

// ProductCreate 新增产品
func (a *sAsset) ProductCreate(ctx context.Context, req *model.ProductCreateReq) (err error) {
	if req.Name == "" {
		return gerror.New("请输入产品名称")
	}
	if req.Code == "" {
		return gerror.New("请输入产品编码")
	}
	if req.CategoryId == 0 {
		return gerror.New("请选择分类")
	}
	category, err := a.getCategoryById(ctx, req.CategoryId)
	if err != nil {
		return err
	}
	if category == nil {
		return gerror.New("分类不存在")
	}

	_, err = dao.AssetProduct.Ctx(ctx).Fields(&model.ProductCreateReq{}).Data(req).Insert()

	return
}

// ProductUpdate 更新产品
func (a *sAsset) ProductUpdate(ctx context.Context, req *model.ProductUpdateReq) (err error) {
	if req.Id == 0 {
		return gerror.New("请输入产品ID")
	}
	if req.Name == "" {
		return gerror.New("请输入产品名称")
	}
	if req.Code == "" {
		return gerror.New("请输入产品编码")
	}
	if req.CategoryId == 0 {
		return gerror.New("请选择分类")
	}

	product, err := a.getProductById(ctx, req.Id)
	if err != nil {
		return err
	}
	if product == nil {
		return gerror.New("产品不存在")
	}

	category, err := a.getCategoryById(ctx, req.CategoryId)
	if err != nil {
		return err
	}
	if category == nil {
		return gerror.New("分类不存在")
	}

	_, err = dao.AssetProduct.Ctx(ctx).WherePri(req.Id).Fields(&model.ProductUpdateReq{}).Data(req).Update()
	return
}

// ProductDelete 删除产品
func (a *sAsset) ProductDelete(ctx context.Context, req *model.ProductDeleteReq) (err error) {
	if req.Id == 0 {
		return gerror.New("请输入ID")
	}
	product, err := a.getProductById(ctx, req.Id)
	if err != nil {
		return err
	}
	if product == nil {
		return gerror.New("产品不存在")
	}
	_, err = dao.AssetProduct.Ctx(ctx).Where(dao.AssetProduct.Columns().Id, req.Id).Delete()
	return
}

// ProductView 产品详情
func (a *sAsset) ProductView(ctx context.Context, req *model.ProductViewReq) (res *model.ProductViewRes, err error) {
	data, err := a.getProductById(ctx, req.Id)
	if err != nil {
		return
	}
	if data == nil {
		return nil, gerror.New("产品不存在")
	}
	return &model.ProductViewRes{
		Id:         data.Id,
		CategoryId: data.CategoryId,
		Name:       data.Name,
		Code:       data.Code,
		Brand:      data.Brand,
		Model:      data.Model,
		Unit:       data.Unit,
		Remark:     data.Remark,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}, nil
}

// ProductList 产品列表
func (a *sAsset) ProductList(ctx context.Context, req *model.ProductListReq) (res []*model.ProductListRes, total int, err error) {
	mod := dao.AssetProduct.Ctx(ctx)
	col := dao.AssetProduct.Columns()

	if req.Id != 0 {
		mod = mod.WherePri(req.Id)
	}
	if req.CategoryId != 0 {
		mod = mod.Where(col.CategoryId, req.CategoryId)
	}
	if req.Name != "" {
		mod = mod.Where(col.Name, req.Name)
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
