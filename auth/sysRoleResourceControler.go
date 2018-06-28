package auth

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type RoleResourceController struct {
	beego.Controller
}

type RoleDistributorResource struct {
	ResourceId []int64
}

func (roleResource *RoleResourceController) RoleDistributorResourceCtl() {

	var roleDistributorResource RoleDistributorResource

	json.Unmarshal(roleResource.Ctx.Input.RequestBody, &roleDistributorResource)
	id, _ := roleResource.GetInt64(":id")
	response := RoleDistributorResourceService(id, &roleDistributorResource)
	roleResource.Data["json"] = response.Data
	roleResource.Ctx.Output.Status = response.StatusCode
	roleResource.ServeJSON()
}

func (roleResource *RoleResourceController) ResourceByRoleIdCtl() {

	id, _ := roleResource.GetInt64(":id")
	response := ResourceByRoleIdService(id)
	roleResource.Data["json"] = response.Data
	roleResource.Ctx.Output.Status = response.StatusCode
	roleResource.ServeJSON()
}
