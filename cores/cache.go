package cores

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/config"
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
	"bytes"
	_ "github.com/astaxie/beego/cache/redis"
)

var GlobalCaches cache.Cache

func InitCache(appConf config.Configer) {
	values := make(map[string]interface{})
	var err error
	if appConf.String("cache::type") == "file" {
		file := fmt.Sprintf(`{"CachePath":"%s","FileSuffix":".bin","DirectoryLevel":2,"EmbedExpiry":0}`,
			appConf.String("cache::file_path"))
		GlobalCaches, _ = cache.NewCache("file", file)
	} else if appConf.String("cache::type") == "redis" {
		values["conn"] = appConf.String("cache::db_host") + ":" + appConf.String("cache::db_port")
		jsonValue, _ := json.Marshal(values)
		GlobalCaches, err = cache.NewCache("redis", bytes.NewBuffer(jsonValue).String())
		if err != nil {
			beego.Debug(err.Error())
		}

	} else {
		GlobalCaches, _ = cache.NewCache("memory", `{"interval":60}`)
	}
}
