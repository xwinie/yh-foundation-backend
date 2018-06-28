package cores

import (
	"testing"
	"github.com/astaxie/beego"
)

func TestModel_GenerateSelectSql(t *testing.T) {
	orm := NewSqlBuild()
	orm.Table("user")
	i := 1
	if i == 1 {
		orm.Where(getWhere())
	}
	beego.Trace(111, orm.GenerateSelectSql())
	orm.Where("1=2").Limit(10,2)
	beego.Trace(111, orm.GenerateSelectSql())
}
func getWhere() string {
	return "1<>2"
}
