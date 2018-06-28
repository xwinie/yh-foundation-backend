package filter

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

func CorsHandler() beego.FilterFunc {
	return func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Methods", "*")
	}
}