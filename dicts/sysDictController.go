package dicts

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/utils/pagination"
	"yh-foundation-backend/cores"
	"net/http"
)

type DictController struct {
	beego.Controller
}

func (dict *DictController) CreateDictCtl() {

	var sysDict SysDict
	json.Unmarshal(dict.Ctx.Input.RequestBody, &sysDict)
	response := CreateDictService(&sysDict)
	dict.Data["json"] = response.Data
	dict.Ctx.Output.Status = response.StatusCode
	dict.ServeJSON()
}

func (dict *DictController) DeleteDictCtl() {
	id, _ := dict.GetInt64(":id")
	response := DeleteDictService(id)
	dict.Data["json"] = response.Data
	dict.Ctx.Output.Status = response.StatusCode
	dict.ServeJSON()
}

func (dict *DictController) UpdateDictCtl() {
	id, _ := dict.GetInt64(":id")
	var data = make(map[string]interface{})
	json.Unmarshal(dict.Ctx.Input.RequestBody, &data)
	response := UpdateDictService(id, data)
	dict.Data["json"] = response.Data
	dict.Ctx.Output.Status = response.StatusCode
	dict.ServeJSON()
}

func (dict *DictController) FindDictByIdCtl() {
	id, _ := dict.GetInt64(":id")
	response := ReadDictByIdService(id)
	dict.Data["json"] = response.Data
	dict.Ctx.Output.Status = response.StatusCode
	dict.ServeJSON()

}

func (dict *DictController) FindDictByPageCtl() {

	var pageQueryDict PageQueryDict
	err := dict.ParseForm(&pageQueryDict)

	var response cores.ResponseEntity
	if err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Data = cores.BuildEntity(ParameterError, GetMsg(ParameterError)+err.Error())
	} else {
		counts := FindDictCountByPageService(pageQueryDict)
		page := pagination.NewPaginator(dict.Ctx.Request, pageQueryDict.PerPage, counts)
		response = FindDictByPageService(page, pageQueryDict)
	}
	dict.Data["json"] = response.Data
	dict.Ctx.Output.Status = response.StatusCode

	dict.ServeJSON()

}
