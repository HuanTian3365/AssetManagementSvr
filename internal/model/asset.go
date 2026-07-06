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
	Id       uint64      `json:"id" dc:"ID"`
	Name     string      `json:"name" dc:"建筑名称"`
	Code     string      `json:"code" dc:"建筑编码"`
	Address  string      `json:"address" dc:"建筑地址"`
	Remark   string      `json:"remark" dc:"备注"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
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
	BuildingId uint64 `json:"building_id" dc:"建筑ID"`
	Name       string `json:"name" dc:"楼层名称"`
	Code       string `json:"code" dc:"楼层编号"`
	FloorNo    int    `json:"floorNo" dc:"楼层序号"`
	Remark     string `json:"remark" dc:"备注"`
}

// FloorUpdateReq 修改楼层
type FloorUpdateReq struct {
	Id         uint64 `json:"id" dc:"ID"`
	BuildingId uint64 `json:"building_id" dc:"建筑ID"`
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
	BuildingId uint64      `json:"building_id" dc:"建筑ID"`
	Name       string      `json:"name" dc:"楼层名称"`
	Code       string      `json:"code" dc:"楼层编号"`
	FloorNo    int         `json:"floorNo" dc:"楼层序号"`
	Remark     string      `json:"remark" dc:"备注"`
	CreateAt   *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt   *gtime.Time `json:"updateAt" dc:"更新时间"`
}

// FloorListReq 楼层列表
type FloorListReq struct {
	common.PageRequest
	Id         uint64        `json:"id" dc:"ID"`
	BuildingId uint64        `json:"building_id" dc:"建筑ID"`
	Name       string        `json:"name" dc:"楼层名称"`
	Code       string        `json:"code" dc:"楼层编号"`
	CreatedAt  []*gtime.Time `json:"createdAt" dc:"创建时间"`
}
type FloorListRes struct {
	Id         uint64      `json:"id" dc:"ID"`
	BuildingId uint64      `json:"building_id" dc:"建筑ID"`
	Name       string      `json:"name" dc:"楼层名称"`
	Code       string      `json:"code" dc:"楼层编号"`
	FloorNo    int         `json:"floorNo" dc:"楼层序号"`
	Remark     string      `json:"remark" dc:"备注"`
	CreateAt   *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt   *gtime.Time `json:"updateAt" dc:"更新时间"`
}
