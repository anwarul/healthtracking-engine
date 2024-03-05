package entity

type PagedQueryEntity struct {
	PageIndex   int    `json:"pageIndex"`
	PageSize    int    `json:"pageSize"`
	OrderBy     string `json:"orderBy"`
	IsAscending bool   `json:"isAscending"`
}
