package auth

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type ResourceController struct {
	beego.Controller
}

func (resource *ResourceController) CreateResourceCtl() {

	var sysResource SysResource
	json.Unmarshal(resource.Ctx.Input.RequestBody, &sysResource)

	response := CreateResourceService(&sysResource)
	resource.Data["json"] = response.Data
	resource.Ctx.Output.Status = response.StatusCode
	resource.ServeJSON()
}

func (resource *ResourceController) DeleteResourceCtl() {
	id, _ := resource.GetInt64(":id")
	response := DeleteResourceService(id)
	resource.Data["json"] = response.Data
	resource.Ctx.Output.Status = response.StatusCode
	resource.ServeJSON()
}

func (resource *ResourceController) UpdateResourceCtl() {
	id, _ := resource.GetInt64(":id")
	var data = make(map[string]interface{})
	json.Unmarshal(resource.Ctx.Input.RequestBody, &data)

	response := UpdateResourceService(id, data)
	resource.Data["json"] = response.Data
	resource.Ctx.Output.Status = response.StatusCode
	resource.ServeJSON()
}

func (resource *ResourceController) FindResourceByCodeCtl() {
	code := resource.GetString(":code")
	response := FindResourceByCodeService(code)
	resource.Data["json"] = response.Data
	resource.Ctx.Output.Status = response.StatusCode
	resource.ServeJSON()

}

func (resource *ResourceController) FindResourceByPageCtl() {
	pageSize, _ := resource.GetInt("perPage")
	counts := FindResourceCountByPageService()
	page := pagination.NewPaginator(resource.Ctx.Request, pageSize, counts)
	response := FindResourceByPageService(page)
	resource.Data["json"] = response.Data
	resource.Ctx.Output.Status = response.StatusCode
	resource.ServeJSON()

}

func (resource *ResourceController) MenusByUserId() {
	userId, _ := resource.GetInt64(":userId")
	response := MenuByUserIdService(userId)
	resource.Data["json"] = response.Data
	resource.Ctx.Output.Status = response.StatusCode
	resource.ServeJSON()

}
