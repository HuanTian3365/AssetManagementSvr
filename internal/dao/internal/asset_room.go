// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AssetRoomDao is the data access object for the table asset_room.
type AssetRoomDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AssetRoomColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AssetRoomColumns defines and stores column names for the table asset_room.
type AssetRoomColumns struct {
	Id         string // ID
	BuildingId string // 建筑ID
	FloorId    string // 楼层ID
	Name       string // 房间名称
	Code       string // 房间编码
	RoomNo     string // 房间号
	Remark     string // 备注
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // 删除时间
}

// assetRoomColumns holds the columns for the table asset_room.
var assetRoomColumns = AssetRoomColumns{
	Id:         "id",
	BuildingId: "building_id",
	FloorId:    "floor_id",
	Name:       "name",
	Code:       "code",
	RoomNo:     "room_no",
	Remark:     "remark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewAssetRoomDao creates and returns a new DAO object for table data access.
func NewAssetRoomDao(handlers ...gdb.ModelHandler) *AssetRoomDao {
	return &AssetRoomDao{
		group:    "default",
		table:    "asset_room",
		columns:  assetRoomColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AssetRoomDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AssetRoomDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AssetRoomDao) Columns() AssetRoomColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AssetRoomDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AssetRoomDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AssetRoomDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
