package common

type Page struct {
	Page     int   `json:"page" dc:"当前页"`
	PageSize int   `json:"pageSize" dc:"每页大小"`
	Total    int64 `json:"total" dc:"总页数"`
}
