package dicts

type PageQueryDict struct {
	PerPage  int    `form:"perPage"`
	Id       int64  `form:"id"`
	Type     string `form:"type"`
	Code     string `form:"code"`
	ParentId int64  `form:"parentId"`
}
