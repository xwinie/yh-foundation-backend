package dicts

import (
	"yh-foundation-backend/conf"
)

const (
	Success         = conf.Dict
	QueryError      = conf.Dict + 1
	ParameterError  = conf.Dict + 2
	NotFoundDict    = conf.Dict + 3
	DictIsExist     = conf.Dict + 4
	CreateDictError = conf.Dict + 5
	DeleteDictError = conf.Dict + 6
	UpdateDictError = conf.Dict + 7
)

var Msg = map[int]string{
	QueryError:      "查询数据异常",
	ParameterError:  "条件异常",
	NotFoundDict:    "没有找到",
	DictIsExist:     "已存在",
	CreateDictError: "添加字典异常",
	DeleteDictError: "删除字典异常",
	UpdateDictError: "修改字典异常",
}

func GetMsg(code int) string {
	return Msg[code]
}
