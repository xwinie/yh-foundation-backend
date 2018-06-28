package auth

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type UserRoleController struct {
	beego.Controller
}

type UserDistributorRole struct {
	RoleId []int64
}

func (userRole *UserRoleController) UserDistributorRoleCtl() {

	var userDistributorRole UserDistributorRole
	json.Unmarshal(userRole.Ctx.Input.RequestBody, &userDistributorRole)
	id, _ := userRole.GetInt64(":id")

	response := UserDistributorRoleService(id, &userDistributorRole)
	userRole.Data["json"] = response.Data
	userRole.Ctx.Output.Status = response.StatusCode
	userRole.ServeJSON()
}

func (userRole *UserRoleController) RoleByUserIdCtl() {

	id, _ := userRole.GetInt64(":id")
	response := RoleByUserIdService(id)
	userRole.Data["json"] = response.Data
	userRole.Ctx.Output.Status = response.StatusCode
	userRole.ServeJSON()
}
