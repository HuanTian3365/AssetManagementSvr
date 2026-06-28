package model

import "github.com/gogf/gf/v2/os/gtime"

// BuildingCreateReq 创建建筑
type BuildingCreateReq struct {
	Name    string `json:"name" dc:"建筑名称"`
	Code    string `json:"code" dc:"建筑编码"`
	Address string `json:"address" dc:"建筑地址"`
	Remark  string `json:"remark" dc:"备注"`
}

// 更新建筑

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
	Id       uint64        `json:"id" dc:"ID"`
	Name     string        `json:"name" dc:"建筑名称"`
	Code     string        `json:"code" dc:"建筑编码"`
	Address  string        `json:"address" dc:"建筑地址"`
	CreateAt []*gtime.Time `json:"createAt" dc:"创建时间"`
}

type BuildingListRes struct {
	Id       uint64      `json:"id" dc:"ID"`
	Name     string      `json:"name" dc:"建筑名称"`
	Code     string      `json:"code" dc:"建筑编码"`
	Address  string      `json:"address" dc:"建筑地址"`
	Remark   string      `json:"remark" dc:"备注"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}
