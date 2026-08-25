package v1

import (
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/model/common"

	"github.com/gogf/gf/v2/frame/g"
)

type CategoryCreateReq struct {
	g.Meta `path:"/category/create" method:"post" tags:"资产分类" summary:"新增资产"`
	model.CategoryCreateReq
}
type CategoryCreateRes struct{}
type CategoryUpdateReq struct {
	g.Meta `path:"/category/update" method:"put" tags:"资产分类" summary:"修改资产"`
	model.CategoryUpdateReq
}
type CategoryUpdateRes struct{}
type CategoryDeleteReq struct {
	g.Meta `path:"/category/delete" method:"delete" tags:"资产分类" summary:"删除资产"`
	model.CategoryDeleteReq
}
type CategoryDeleteRes struct{}
type CategoryViewReq struct {
	g.Meta `path:"/category/view" method:"get" tags:"资产分类" summary:"资产详情"`
	model.CategoryViewReq
}
type CategoryViewRes struct {
	*model.CategoryViewRes
}
type CategoryListReq struct {
	g.Meta `path:"/category/list" method:"get" tags:"资产分类" summary:"资产列表"`
	model.CategoryListReq
}
type CategoryListRes struct {
	common.PageResult
	Items []*model.CategoryListRes `json:"items"`
}
