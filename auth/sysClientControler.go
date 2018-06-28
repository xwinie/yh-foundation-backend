package auth

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type ClientController struct {
	beego.Controller
}

func (client *ClientController) CreateClientCtl() {

	var sysClient SysClient
	json.Unmarshal(client.Ctx.Input.RequestBody, &sysClient)

	response := CreateClientService(&sysClient)
	client.Data["json"] = response.Data
	client.Ctx.Output.Status = response.StatusCode
	client.ServeJSON()
}

func (client *ClientController) DeleteClientCtl() {
	id, _ := client.GetInt64(":id")

	response := DeleteClientService(id)
	client.Data["json"] = response.Data
	client.Ctx.Output.Status = response.StatusCode
	client.ServeJSON()
}

func (client *ClientController) UpdateClientCtl() {
	id, _ := client.GetInt64(":id")
	var data = make(map[string]interface{})
	json.Unmarshal(client.Ctx.Input.RequestBody, &data)

	response := UpdateClientService(id, data)
	client.Data["json"] = response.Data
	client.Ctx.Output.Status = response.StatusCode
	client.ServeJSON()
}

func (client *ClientController) FindClientByClientIdCtl() {
	clientId := client.GetString(":clientId")
	response := FindClientByClientIdService(clientId)
	client.Data["json"] = response.Data
	client.Ctx.Output.Status = response.StatusCode
	client.ServeJSON()
}

func (client *ClientController) FindClientDefault() {
	clientID := client.Ctx.Input.Header("appid")
	response := FindClientByClientIdService(clientID)
	client.Data["json"] = response.Data
	client.Ctx.Output.Status = response.StatusCode
	client.ServeJSON()

}

func (client *ClientController) FindClientByPageCtl() {
	pageSize, _ := client.GetInt("perPage")
	counts := FindClientCountByPageService()
	page := pagination.NewPaginator(client.Ctx.Request, pageSize, counts)
	response := FindClientByPageService(page)
	client.Data["json"] = response.Data
	client.Ctx.Output.Status = response.StatusCode
	client.ServeJSON()

}

type ClientAuthData struct {
	ClientID string
	Secret   string
}

func (client *ClientController) Auth() {

	var clientAuthData ClientAuthData
	json.Unmarshal(client.Ctx.Input.RequestBody, &clientAuthData)

	response := SignAuthService(&clientAuthData)
	client.Data["json"] = response.Data
	client.Ctx.Output.Status = response.StatusCode
	client.ServeJSON()

}
