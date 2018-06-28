package main

import (
	"yh-foundation-backend/cores"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"yh-foundation-backend/routers"
	"time"
	"yh-foundation-backend/filter"
)

func init() {

	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	cores.InitDB(appConf)
	cores.InitCache(appConf)
	cores.InitLog(appConf)

}
func main() {
	routers.AuthRouter()
	beego.InsertFilter("*", beego.BeforeRouter, filter.CorsHandler())
	beego.InsertFilter("*", beego.BeforeRouter, filter.APISecretAuth(100))
	beego.InsertFilter("*", beego.BeforeRouter, filter.Authorizer())
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		orm.Debug = true
	}
	orm.DefaultTimeLoc = time.UTC
	beego.ErrorController(&cores.ErrorController{})
	beego.BConfig.ServerName = "yh api server 1.0"
	beego.Run()
}
