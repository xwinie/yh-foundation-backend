package auth

import (
	"yh-foundation-backend/conf"
	"net/http"
)

const (
	Unauthorized                 = http.StatusUnauthorized
	Success                      = conf.Auth
	NotFoundUser                 = conf.Auth + 1
	GenerateTokenError           = conf.Auth + 2
	UserIsExist                  = conf.Auth + 3
	CreateUserError              = conf.Auth + 4
	DeleteUserError              = conf.Auth + 5
	UpdateUserError              = conf.Auth + 6
	RoleIsExist                  = conf.Auth + 7
	CreateRoleError              = conf.Auth + 8
	DeleteRoleError              = conf.Auth + 9
	UpdateRoleError              = conf.Auth + 10
	ResourceIsExist              = conf.Auth + 11
	CreateResourceError          = conf.Auth + 12
	DeleteResourceError          = conf.Auth + 13
	UpdateResourceError          = conf.Auth + 14
	RoleDistributorResourceError = conf.Auth + 15
	DeleteRoleResourceError      = conf.Auth + 16
	UserDistributorRoleError     = conf.Auth + 17
	DeleteUserRoleError          = conf.Auth + 18
	ClientIsExist                = conf.Auth + 19
	CreateClientError            = conf.Auth + 20
	DeleteClientError            = conf.Auth + 21
	UpdateClientError            = conf.Auth + 22
	NotFoundUserRole             = conf.Auth + 23
	ParameterError               = conf.Auth + 24
	QueryError                   = conf.Auth + 25
)

var Msg = map[int]string{
	Unauthorized:                 "没有权限",
	NotFoundUser:                 "无此用户",
	NotFoundUserRole:             "获取数据异常",
	GenerateTokenError:           "生成token异常",
	UserIsExist:                  "用户已存在",
	CreateUserError:              "创建用户失败",
	DeleteUserError:              "删除用户异常，用户不存在",
	UpdateUserError:              "更新用户异常，用户不存在",
	UserDistributorRoleError:     "用户分配角色异常",
	DeleteUserRoleError:          "删除用户角色异常",
	Success:                      "成功",
	ParameterError:               "参数异常",
	QueryError:                   "查询异常",
	RoleIsExist:                  "角色已存在",
	CreateRoleError:              "创建角色失败",
	DeleteRoleError:              "删除角色异常，角色不存在",
	UpdateRoleError:              "更新角色异常，角色不存在",
	RoleDistributorResourceError: "角色分配资源异常",
	DeleteRoleResourceError:      "删除角色资源异常",
	ResourceIsExist:              "资源已存在",
	CreateResourceError:          "创建资源失败",
	DeleteResourceError:          "删除资源异常，资源不存在",
	UpdateResourceError:          "更新资源异常，资源不存在",
	ClientIsExist:                "应用已存在",
	CreateClientError:            "创建应用失败",
	DeleteClientError:            "删除应用异常，应用不存在",
	UpdateClientError:            "更新应用异常，应用不存在",
}

func GetMsg(code int) string {
	return Msg[code]
}
