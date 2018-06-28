package auth

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type RoleController struct {
	beego.Controller
}

func (role *RoleController) CreateRoleCtl() {

	var sysRole SysRole
	json.Unmarshal(role.Ctx.Input.RequestBody, &sysRole)

	response := CreateRoleService(&sysRole)
	role.Data["json"] = response.Data
	role.Ctx.Output.Status = response.StatusCode
	role.ServeJSON()
}

func (role *RoleController) DeleteRoleCtl() {
	id, _ := role.GetInt64(":id")

	response := DeleteRoleService(id)
	role.Data["json"] = response.Data
	role.Ctx.Output.Status = response.StatusCode
	role.ServeJSON()
}

func (role *RoleController) UpdateRoleCtl() {
	id, _ := role.GetInt64(":id")
	var data = make(map[string]interface{})
	json.Unmarshal(role.Ctx.Input.RequestBody, &data)

	response := UpdateRoleService(id, data)
	role.Data["json"] = response.Data
	role.Ctx.Output.Status = response.StatusCode
	role.ServeJSON()
}
func (role *RoleController) FindRoleByCodeCtl() {
	code := role.GetString(":code")
	response := FindRoleByCodeService(code)
	role.Data["json"] = response.Data
	role.Ctx.Output.Status = response.StatusCode
	role.ServeJSON()

}

func (role *RoleController) FindRoleByIdCtl() {
	id, _ := role.GetInt64(":id")
	//beego.Trace("1111" + fmt.Sprintf("%d", id))
	response := FindRoleByIdService(id)
	role.Data["json"] = response.Data
	role.Ctx.Output.Status = response.StatusCode
	role.ServeJSON()

}

func (role *RoleController) FindRoleByPageCtl() {
	pageSize, _ := role.GetInt("perPage")
	counts := FindRoleCountByPageService()
	page := pagination.NewPaginator(role.Ctx.Request, pageSize, counts)
	response := FindRoleByPageService(page)
	role.Data["json"] = response.Data
	role.Ctx.Output.Status = response.StatusCode
	role.ServeJSON()

}
