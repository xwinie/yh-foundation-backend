package cores

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"net/url"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/astaxie/beego/config"
)

var Conn *gorm.DB

func Connect(appConf config.Configer) (db *gorm.DB, err error) {
	if appConf.String("database::db_driver") == "mysql" {
		conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=%v&parseTime=true",
			appConf.String("database::db_user"),
			appConf.String("database::db_passwd"),
			appConf.String("database::db_host"),
			appConf.String("database::db_port"),
			appConf.String("database::db_name"),
			appConf.String("database::db_charset"),
			url.QueryEscape("Asia/Shanghai"))
		db, err = gorm.Open("mysql", conn)
	}
	Conn = db
	return
}
