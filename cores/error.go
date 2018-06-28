package cores

import (
	"github.com/astaxie/beego"
	"net/http"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = Entity{
		Code: http.StatusNotFound,
		Msg:  "Not Found",
	}
	c.ServeJSON()
}
func (c *ErrorController) Error401() {
	c.Data["json"] = Entity{
		Code: http.StatusUnauthorized,
		Msg:  "Permission denied",
	}
	c.ServeJSON()
}
func (c *ErrorController) Error403() {
	c.Data["json"] = Entity{
		Code: http.StatusForbidden,
		Msg:  "Forbidden",
	}
	c.ServeJSON()
}
