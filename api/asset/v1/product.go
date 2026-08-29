package v1

import (
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/model/common"

	"github.com/gogf/gf/v2/frame/g"
)

// ProductCreateReq 新增产品
type ProductCreateReq struct {
	g.Meta `path:"/product/create" method:"post" tags:"产品管理" summary:"新增产品"`
	model.ProductCreateReq
}
type ProductCreateRes struct{}

// ProductUpdateReq 更新产品
type ProductUpdateReq struct {
	g.Meta `path:"/product/update" method:"put" tags:"产品管理" summary:"修改产品"`
	model.ProductUpdateReq
}
type ProductUpdateRes struct{}

// ProductDeleteReq 删除产品
type ProductDeleteReq struct {
	g.Meta `path:"/product/delete" method:"delete" tags:"产品管理" summary:"删除产品"`
	model.ProductDeleteReq
}
type ProductDeleteRes struct{}

// ProductViewReq 产品详情
type ProductViewReq struct {
	g.Meta `path:"/product/view" method:"get" tags:"产品管理" summary:"产品详情"`
	model.ProductViewReq
}
type ProductViewRes struct {
	*model.ProductViewRes
}

// ProductListReq 产品列表
type ProductListReq struct {
	g.Meta `path:"/product/list" method:"get" tags:"产品管理" summary:"产品列表"`
	model.ProductListReq
}
type ProductListRes struct {
	common.PageResult
	Items []*model.ProductListRes `json:"items"`
}
