package model

import (
	"asset_management_svr/internal/model/common"

	"github.com/gogf/gf/v2/os/gtime"
)

// BuildingCreateReq 创建建筑
type BuildingCreateReq struct {
	Name    string `json:"name" dc:"建筑名称"`
	Code    string `json:"code" dc:"建筑编码"`
	Address string `json:"address" dc:"建筑地址"`
	Remark  string `json:"remark" dc:"备注"`
}

// BuildingUpdateReq 更新建筑
type BuildingUpdateReq struct {
	Id      uint64 `json:"id" dc:"ID"`
	Name    string `json:"name" dc:"建筑名称"`
	Code    string `json:"code" dc:"建筑编码"`
	Address string `json:"address" dc:"建筑地址"`
	Remark  string `json:"remark" dc:"备注"`
}

// BuildingDeleteReq 删除建筑
type BuildingDeleteReq struct {
	Id uint64 `json:"id" dc:"ID"`
}

// BuildingViewReq 建筑详情
type BuildingViewReq struct {
	Id uint64 `json:"id" dc:"ID"`
}

type BuildingViewRes struct {
	Id        uint64      `json:"id" dc:"ID"`
	Name      string      `json:"name" dc:"建筑名称"`
	Code      string      `json:"code" dc:"建筑编码"`
	Address   string      `json:"address" dc:"建筑地址"`
	Remark    string      `json:"remark" dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// BuildingListReq 建筑列表
type BuildingListReq struct {
	common.PageRequest
	Id        uint64        `json:"id" dc:"ID"`
	Name      string        `json:"name" dc:"建筑名称"`
	Code      string        `json:"code" dc:"建筑编码"`
	Address   string        `json:"address" dc:"建筑地址"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

type BuildingListRes struct {
	Id        uint64      `json:"id" dc:"ID"`
	Name      string      `json:"name" dc:"建筑名称"`
	Code      string      `json:"code" dc:"建筑编码"`
	Address   string      `json:"address" dc:"建筑地址"`
	Remark    string      `json:"remark" dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// FloorCreateReq 创建楼层
type FloorCreateReq struct {
	BuildingId uint64 `json:"buildingId" dc:"建筑ID"`
	Name       string `json:"name" dc:"楼层名称"`
	Code       string `json:"code" dc:"楼层编号"`
	FloorNo    int    `json:"floorNo" dc:"楼层序号"`
	Remark     string `json:"remark" dc:"备注"`
}

// FloorUpdateReq 修改楼层
type FloorUpdateReq struct {
	Id         uint64 `json:"id" dc:"ID"`
	BuildingId uint64 `json:"buildingId" dc:"建筑ID"`
	Name       string `json:"name" dc:"楼层名称"`
	Code       string `json:"code" dc:"楼层编号"`
	FloorNo    int    `json:"floorNo" dc:"楼层序号"`
	Remark     string `json:"remark" dc:"备注"`
}

// FloorDeleteReq 删除楼层
type FloorDeleteReq struct {
	Id uint64 `json:"id" dc:"ID"`
}

// FloorViewReq 建筑详情
type FloorViewReq struct {
	Id uint64 `json:"id" dc:"ID"`
}
type FloorViewRes struct {
	Id         uint64      `json:"id" dc:"ID"`
	BuildingId uint64      `json:"buildingId" dc:"建筑ID"`
	Name       string      `json:"name" dc:"楼层名称"`
	Code       string      `json:"code" dc:"楼层编号"`
	FloorNo    int         `json:"floorNo" dc:"楼层序号"`
	Remark     string      `json:"remark" dc:"备注"`
	CreatedAt  *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// FloorListReq 楼层列表
type FloorListReq struct {
	common.PageRequest
	Id         uint64        `json:"id" dc:"ID"`
	BuildingId uint64        `json:"buildingId" dc:"建筑ID"`
	Name       string        `json:"name" dc:"楼层名称"`
	Code       string        `json:"code" dc:"楼层编号"`
	CreatedAt  []*gtime.Time `json:"createdAt" dc:"创建时间"`
}
type FloorListRes struct {
	Id         uint64      `json:"id" dc:"ID"`
	BuildingId uint64      `json:"buildingId" dc:"建筑ID"`
	Name       string      `json:"name" dc:"楼层名称"`
	Code       string      `json:"code" dc:"楼层编号"`
	FloorNo    int         `json:"floorNo" dc:"楼层序号"`
	Remark     string      `json:"remark" dc:"备注"`
	CreatedAt  *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// RoomCreateReq 创建房间
type RoomCreateReq struct {
	BuildingId uint64 `json:"buildingId" dc:"建筑ID"`
	FloorId    uint64 `json:"floorId" dc:"楼层ID"`
	Name       string `json:"name" dc:"房间名称"`
	Code       string `json:"code" dc:"房间编码"`
	RoomNo     string `json:"roomNo" dc:"房间号"`
	Remark     string `json:"remark" dc:"备注"`
}

// RoomUpdateReq 修改房间
type RoomUpdateReq struct {
	Id         uint64 `json:"id" dc:"ID"`
	BuildingId uint64 `json:"buildingId" dc:"建筑ID"`
	FloorId    uint64 `json:"floorId" dc:"楼层ID"`
	Name       string `json:"name" dc:"房间名称"`
	Code       string `json:"code" dc:"房间编码"`
	RoomNo     string `json:"roomNo" dc:"房间号"`
	Remark     string `json:"remark" dc:"备注"`
}

// RoomDeleteReq 删除房间
type RoomDeleteReq struct {
	Id uint64 `json:"id" dc:"ID"`
}

// RoomViewReq 房间详情
type RoomViewReq struct {
	Id uint64 `json:"id" dc:"ID"`
}
type RoomViewRes struct {
	Id         uint64      `json:"id" dc:"ID"`
	BuildingId uint64      `json:"buildingId" dc:"建筑ID"`
	FloorId    uint64      `json:"floorId" dc:"楼层ID"`
	Name       string      `json:"name" dc:"房间名称"`
	Code       string      `json:"code" dc:"房间编码"`
	RoomNo     string      `json:"roomNo" dc:"房间号"`
	Remark     string      `json:"remark" dc:"备注"`
	CreatedAt  *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// RoomListReq 房间列表
type RoomListReq struct {
	common.PageRequest
	Id         uint64        `json:"id" dc:"ID"`
	BuildingId uint64        `json:"buildingId" dc:"建筑ID"`
	FloorId    uint64        `json:"floorId" dc:"楼层ID"`
	Name       string        `json:"name" dc:"房间名称"`
	Code       string        `json:"code" dc:"房间编码"`
	CreatedAt  []*gtime.Time `json:"createdAt" dc:"创建时间"`
}
type RoomListRes struct {
	Id         uint64      `json:"id" dc:"ID"`
	BuildingId uint64      `json:"buildingId" dc:"建筑ID"`
	FloorId    uint64      `json:"floorId" dc:"楼层ID"`
	Name       string      `json:"name" dc:"房间名称"`
	Code       string      `json:"code" dc:"房间编码"`
	RoomNo     string      `json:"roomNo" dc:"房间号"`
	Remark     string      `json:"remark" dc:"备注"`
	CreatedAt  *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// CategoryCreateReq 新增分类
type CategoryCreateReq struct {
	ParentId uint64 `json:"parentId" dc:"父分类ID"`
	Name     string `json:"name" dc:"分类名称"`
	Code     string `json:"code" dc:"分类编码"`
	Sort     int    `json:"sort" dc:"分类排序"`
	Remark   string `json:"remark" dc:"备注"`
}

// CategoryUpdateReq 更新分类
type CategoryUpdateReq struct {
	Id       uint64 `json:"id" dc:"ID"`
	ParentId uint64 `json:"parentId" dc:"父分类ID"`
	Name     string `json:"name" dc:"分类名称"`
	Code     string `json:"code" dc:"分类编码"`
	Sort     int    `json:"sort" dc:"分类排序"`
	Remark   string `json:"remark" dc:"备注"`
}

// CategoryDeleteReq 删除分类
type CategoryDeleteReq struct {
	Id uint64 `json:"id" dc:"ID"`
}

// CategoryViewReq 分类详情
type CategoryViewReq struct {
	Id uint64 `json:"id" dc:"ID"`
}
type CategoryViewRes struct {
	Id        uint64      `json:"id" dc:"ID"`
	ParentId  uint64      `json:"parentId" dc:"父分类ID"`
	Name      string      `json:"name" dc:"分类名称"`
	Code      string      `json:"code" dc:"分类编码"`
	Sort      int         `json:"sort" dc:"分类排序"`
	Remark    string      `json:"remark" dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// CategoryListReq 分类列表
type CategoryListReq struct {
	common.PageRequest
	Id        uint64        `json:"id" dc:"ID"`
	ParentId  uint64        `json:"parentId" dc:"父分类ID"`
	Name      string        `json:"name" dc:"分类名称"`
	Code      string        `json:"code" dc:"分类编码"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}
type CategoryListRes struct {
	Id        uint64      `json:"id" dc:"ID"`
	ParentId  uint64      `json:"parentId" dc:"父分类ID"`
	Name      string      `json:"name" dc:"分类名称"`
	Code      string      `json:"code" dc:"分类编码"`
	Sort      int         `json:"sort" dc:"分类排序"`
	Remark    string      `json:"remark" dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// ProductCreateReq 新增产品
type ProductCreateReq struct {
	CategoryId uint64 `json:"categoryId" dc:"所属分类"`
	Name       string `json:"name" dc:"产品名称"`
	Code       string `json:"code" dc:"产品编码"`
	Brand      string `json:"brand" dc:"品牌"`
	Model      string `json:"model" dc:"型号"`
	Unit       string `json:"unit" dc:"计量单位"`
	Remark     string `json:"remark" dc:"备注"`
}

// ProductUpdateReq 更新产品
type ProductUpdateReq struct {
	Id         uint64 `json:"id" dc:"ID"`
	CategoryId uint64 `json:"categoryId" dc:"所属分类"`
	Name       string `json:"name" dc:"产品名称"`
	Code       string `json:"code" dc:"产品编码"`
	Brand      string `json:"brand" dc:"品牌"`
	Model      string `json:"model" dc:"型号"`
	Unit       string `json:"unit" dc:"计量单位"`
	Remark     string `json:"remark" dc:"备注"`
}

// ProductDeleteReq 删除产品
type ProductDeleteReq struct {
	Id uint64 `json:"id" dc:"ID"`
}

// ProductViewReq 产品详情
type ProductViewReq struct {
	Id uint64 `json:"id" dc:"ID"`
}
type ProductViewRes struct {
	Id         uint64      `json:"id" dc:"ID"`
	CategoryId uint64      `json:"categoryId" dc:"所属分类"`
	Name       string      `json:"name" dc:"产品名称"`
	Code       string      `json:"code" dc:"产品编码"`
	Brand      string      `json:"brand" dc:"品牌"`
	Model      string      `json:"model" dc:"型号"`
	Unit       string      `json:"unit" dc:"计量单位"`
	Remark     string      `json:"remark" dc:"备注"`
	CreatedAt  *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// ProductListReq 产品列表
type ProductListReq struct {
	common.PageRequest
	Id         uint64        `json:"id" dc:"ID"`
	CategoryId uint64        `json:"categoryId" dc:"所属分类"`
	Name       string        `json:"name" dc:"产品名称"`
	Code       string        `json:"code" dc:"产品编码"`
	CreatedAt  []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

type ProductListRes struct {
	Id         uint64      `json:"id" dc:"ID"`
	CategoryId uint64      `json:"categoryId" dc:"所属分类"`
	Name       string      `json:"name" dc:"产品名称"`
	Code       string      `json:"code" dc:"产品编码"`
	Brand      string      `json:"brand" dc:"品牌"`
	Model      string      `json:"model" dc:"型号"`
	Unit       string      `json:"unit" dc:"计量单位"`
	Remark     string      `json:"remark" dc:"备注"`
	CreatedAt  *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// ItemCreateReq 新增物品
type ItemCreateReq struct {
	ProductId         uint64      `json:"productId" dc:"产品ID"`
	Name              string      `json:"name" dc:"资产名称"`
	AssetCode         string      `json:"assetCode" dc:"资产编码"`
	Status            int8        `json:"status" dc:"资产状态:1闲置 2在用 3维修 4报废"`
	CurrentBuildingId uint64      `json:"currentBuildingId" dc:"建筑ID"`
	CurrentFloorId    uint64      `json:"currentFloorId" dc:"楼层ID"`
	CurrentRoomId     uint64      `json:"currentRoomId" dc:"房间ID"`
	PurchaseDate      *gtime.Time `json:"purchaseDate" dc:"购买日期"`
	PurchasePrice     float32     `json:"purchasePrice" dc:"购买价格"`
	Remark            string      `json:"remark" dc:"备注"`
}

// ItemUpdateReq 更新物品
type ItemUpdateReq struct {
	Id                uint64      `json:"id" dc:"ID"`
	ProductId         uint64      `json:"productId" dc:"产品ID"`
	Name              string      `json:"name" dc:"资产名称"`
	AssetCode         string      `json:"assetCode" dc:"资产编码"`
	Status            int8        `json:"status" dc:"资产状态:1闲置 2在用 3维修 4报废"`
	CurrentBuildingId uint64      `json:"currentBuildingId" dc:"建筑ID"`
	CurrentFloorId    uint64      `json:"currentFloorId" dc:"楼层ID"`
	CurrentRoomId     uint64      `json:"currentRoomId" dc:"房间ID"`
	PurchaseDate      *gtime.Time `json:"purchaseDate" dc:"购买日期"`
	PurchasePrice     float32     `json:"purchasePrice" dc:"购买价格"`
	Remark            string      `json:"remark" dc:"备注"`
}

// ItemDeleteReq 删除物品
type ItemDeleteReq struct {
	Id uint64 `json:"id" dc:"ID"`
}

// ItemViewReq 物品详情
type ItemViewReq struct {
	Id uint64 `json:"id" dc:"ID"`
}
type ItemViewRes struct {
	Id                uint64      `json:"id" dc:"ID"`
	ProductId         uint64      `json:"productId" dc:"产品ID"`
	Name              string      `json:"name" dc:"资产名称"`
	AssetCode         string      `json:"assetCode" dc:"资产编码"`
	Status            int8        `json:"status" dc:"资产状态:1闲置 2在用 3维修 4报废"`
	CurrentBuildingId uint64      `json:"currentBuildingId" dc:"建筑ID"`
	CurrentFloorId    uint64      `json:"currentFloorId" dc:"楼层ID"`
	CurrentRoomId     uint64      `json:"currentRoomId" dc:"房间ID"`
	PurchaseDate      *gtime.Time `json:"purchaseDate" dc:"购买日期"`
	PurchasePrice     float32     `json:"purchasePrice" dc:"购买价格"`
	Remark            string      `json:"remark" dc:"备注"`
	CreatedAt         *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// ItemListReq 物品列表
type ItemListReq struct {
	common.PageRequest
	Id        uint64        `json:"id" dc:"ID"`
	ProductId uint64        `json:"productId" dc:"产品ID"`
	Name      string        `json:"name" dc:"资产名称"`
	AssetCode string        `json:"assetCode" dc:"资产编码"`
	Status    int8          `json:"status" dc:"资产状态:1闲置 2在用 3维修 4报废"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}
type ItemListRes struct {
	Id                uint64      `json:"id" dc:"ID"`
	ProductId         uint64      `json:"productId" dc:"产品ID"`
	Name              string      `json:"name" dc:"资产名称"`
	AssetCode         string      `json:"assetCode" dc:"资产编码"`
	Status            int8        `json:"status" dc:"资产状态:1闲置 2在用 3维修 4报废"`
	CurrentBuildingId uint64      `json:"currentBuildingId" dc:"建筑ID"`
	CurrentFloorId    uint64      `json:"currentFloorId" dc:"楼层ID"`
	CurrentRoomId     uint64      `json:"currentRoomId" dc:"房间ID"`
	PurchaseDate      *gtime.Time `json:"purchaseDate" dc:"购买日期"`
	PurchasePrice     float64     `json:"purchasePrice" dc:"购买价格"`
	Remark            string      `json:"remark" dc:"备注"`
	CreatedAt         *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt" dc:"更新时间"`
}
