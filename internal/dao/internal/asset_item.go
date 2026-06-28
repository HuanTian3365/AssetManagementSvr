// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AssetItemDao is the data access object for the table asset_item.
type AssetItemDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AssetItemColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AssetItemColumns defines and stores column names for the table asset_item.
type AssetItemColumns struct {
	Id                string // ID
	ProductId         string // 产品ID
	AssetCode         string // 资产编码
	Name              string // 资产名称
	Status            string // 资产状态：1闲置 2在用 3维修 4报废
	CurrentBuildingId string // 建筑ID
	CurrentFloorId    string // 楼层ID
	CurrentRoomId     string // 房间ID
	PurchaseDate      string // 购买日期
	PurchasePrice     string // 购买价格
	Remark            string // 备注
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
	DeletedAt         string // 删除时间
}

// assetItemColumns holds the columns for the table asset_item.
var assetItemColumns = AssetItemColumns{
	Id:                "id",
	ProductId:         "product_id",
	AssetCode:         "asset_code",
	Name:              "name",
	Status:            "status",
	CurrentBuildingId: "current_building_id",
	CurrentFloorId:    "current_floor_id",
	CurrentRoomId:     "current_room_id",
	PurchaseDate:      "purchase_date",
	PurchasePrice:     "purchase_price",
	Remark:            "remark",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
}

// NewAssetItemDao creates and returns a new DAO object for table data access.
func NewAssetItemDao(handlers ...gdb.ModelHandler) *AssetItemDao {
	return &AssetItemDao{
		group:    "default",
		table:    "asset_item",
		columns:  assetItemColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AssetItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AssetItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AssetItemDao) Columns() AssetItemColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AssetItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AssetItemDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AssetItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
