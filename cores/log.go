package cores

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"encoding/json"
	"bytes"
	"github.com/astaxie/beego/logs"
)

func InitLog(appConf config.Configer) {
	values := make(map[string]interface{})
	values["filename"] = appConf.String("log::filename")
	values["level"], _ = appConf.Int("log::level")
	jsonValue, _ := json.Marshal(values)

	logs.SetLogger(logs.AdapterMultiFile, bytes.NewBuffer(jsonValue).String())
	beego.SetLogFuncCall(true)
}
