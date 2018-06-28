package cores

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"net/url"
	"time"
)

func InitDB(appConf config.Configer) {

	if appConf.String("database::db_driver") == "mysql" {
		conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=%v&parseTime=true",
			appConf.String("database::db_user"),
			appConf.String("database::db_passwd"),
			appConf.String("database::db_host"),
			appConf.String("database::db_port"),
			appConf.String("database::db_name"),
			appConf.String("database::db_charset"),
			url.QueryEscape("Asia/Shanghai"))
		orm.RegisterDataBase("default", "mysql", conn)
		orm.DefaultTimeLoc = time.UTC
	}

	if db_Auto_Ddl, err := strconv.ParseBool(appConf.String("database::db_autoddl")); db_Auto_Ddl && err == nil {
		//自动建表
		name := "default"
		force := false
		verbose := true
		err = orm.RunSyncdb(name, force, verbose)
		if err != nil {
			beego.Error("init db table error:", err)
		}
		orm.RunCommand()
	}

}
