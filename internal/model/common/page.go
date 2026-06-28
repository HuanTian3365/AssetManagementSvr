package common

type PageRequest struct {
	Page     int  `json:"page" dc:"当前页"`
	PageSize int  `json:"pageSize" dc:"每页大小"`
	FullSize bool `json:"fullSize" dc:"全量拉取"`
}

type PageResult struct {
	Page     int   `json:"page" dc:"当前页"`
	PageSize int   `json:"pageSize" dc:"每页大小"`
	Total    int64 `json:"total" dc:"总页数"`
}
