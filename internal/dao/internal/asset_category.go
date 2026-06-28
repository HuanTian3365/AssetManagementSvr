// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AssetCategoryDao is the data access object for the table asset_category.
type AssetCategoryDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  AssetCategoryColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// AssetCategoryColumns defines and stores column names for the table asset_category.
type AssetCategoryColumns struct {
	Id        string // ID
	ParentId  string // 父分类ID
	Name      string // 分类名称
	Code      string // 分类编码
	Sort      string // 排序
	Remark    string // 备注
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// assetCategoryColumns holds the columns for the table asset_category.
var assetCategoryColumns = AssetCategoryColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Name:      "name",
	Code:      "code",
	Sort:      "sort",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewAssetCategoryDao creates and returns a new DAO object for table data access.
func NewAssetCategoryDao(handlers ...gdb.ModelHandler) *AssetCategoryDao {
	return &AssetCategoryDao{
		group:    "default",
		table:    "asset_category",
		columns:  assetCategoryColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AssetCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AssetCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AssetCategoryDao) Columns() AssetCategoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AssetCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AssetCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AssetCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
