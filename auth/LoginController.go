package auth

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type LoginController struct {
	beego.Controller
}

type LoginData struct {
	UserName string
	Password string
}

func (lc *LoginController) Login() {
	var loginData LoginData
	json.Unmarshal(lc.Ctx.Input.RequestBody, &loginData)
	appID := lc.Ctx.Input.Header("appid")
	response := LoginService(&loginData, appID)
	lc.Data["json"] = response.Data
	lc.Ctx.Output.Status = response.StatusCode
	lc.ServeJSON()
}
