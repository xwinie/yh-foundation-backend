package auth

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type UserController struct {
	beego.Controller
}

func (user *UserController) CreateUserCtl() {

	var sysUser SysUser
	json.Unmarshal(user.Ctx.Input.RequestBody, &sysUser)
	response := CreateUserService(&sysUser)
	user.Data["json"] = response.Data
	user.Ctx.Output.Status = response.StatusCode
	user.ServeJSON()
}

func (user *UserController) DeleteUserCtl() {
	id, _ := user.GetInt64(":id")
	response := DeleteUserService(id)
	user.Data["json"] = response.Data
	user.Ctx.Output.Status = response.StatusCode
	user.ServeJSON()
}

func (user *UserController) UpdateUserCtl() {
	id, _ := user.GetInt64(":id")
	var data = make(map[string]interface{})
	json.Unmarshal(user.Ctx.Input.RequestBody, &data)
	response := UpdateUserService(id, data)
	user.Data["json"] = response.Data
	user.Ctx.Output.Status = response.StatusCode
	user.ServeJSON()
}

func (user *UserController) FindUserByAccountCtl() {
	account := user.GetString(":account")
	response := FindUserByAccountService(account)
	user.Data["json"] = response.Data
	user.Ctx.Output.Status = response.StatusCode
	user.ServeJSON()

}

func (user *UserController) FindUserByPageCtl() {
	pageSize, _ := user.GetInt("perPage")
	counts := FindUserCountByPageService()
	page := pagination.NewPaginator(user.Ctx.Request, pageSize, counts)
	response := FindUserByPageService(page)
	user.Data["json"] = response.Data
	user.Ctx.Output.Status = response.StatusCode
	user.ServeJSON()

}
