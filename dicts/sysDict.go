package dicts

import (
	"github.com/astaxie/beego/orm"
	"time"
	"bytes"
	"fmt"
	"text/template"
)

type SysDict struct {
	Id         int64
	Code       string    `orm:"size(100)"`
	City       string    `orm:"size(8);null"`
	Name       string    `orm:"size(100);null"`
	Level      int       `orm:"size(3)"`
	ParentId   int64
	ParentCode string    `orm:"size(100);"`
	Type       string    `orm:"size(6)"`
	Status     int8      `orm:"default(0)"`
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysDict))
}

func CreateDict(p *SysDict) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(p)
	return id, err
}

func CheckCodeIsExist(code string) (bool, error) {
	o := orm.NewOrm()
	var dict SysDict
	err := o.QueryTable("sys_dict").Filter("code", code).One(&dict)
	if &dict.Id != nil && err == nil {
		return true, err
	} else {
		return false, err
	}
}

func DeleteDict(p *SysDict) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(p)
	return num, err

}

func UpdateDict(p map[string]interface{}, id int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_dict").Filter("id", id).Update(p)
	return num, err

}

func ReadDict(code string) (*SysDict, error) {
	o := orm.NewOrm()
	var dict SysDict
	err := o.QueryTable("sys_dict").Filter("code", code).One(&dict)
	return &dict, err
}

func ReadDictById(id int64) (*SysDict, error) {

	o := orm.NewOrm()
	var dict SysDict
	err := o.QueryTable("sys_dict").Filter("id", id).One(&dict)
	return &dict, err
}

func ReadDictByPage(pg PageQueryDict, pageSize int, offset int) (dicts []*SysDict, num int64, err error) {

	o := orm.NewOrm()
	qs := o.QueryTable("sys_dict")

	cond := orm.NewCondition()
	if pg.Id > 0 {
		qs = qs.SetCond(cond.And("id", pg.Id))
	}
	if pg.ParentId > 0 {
		qs = qs.SetCond(cond.And("parent_id", pg.ParentId))
	}
	if pg.Type != "" {
		qs = qs.SetCond(cond.And("type", pg.Type))

	}
	if pg.Code != "" && pg.Type != "" {
		qs = qs.SetCond(cond.And("code", pg.Code).And("type", pg.Type))
	}
	num, err = qs.Limit(pageSize, offset).All(&dicts)

	return dicts, num, err
}

func ReadDictCountByPage(pg PageQueryDict) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_dict")

	cond := orm.NewCondition()
	if pg.Id > 0 {
		qs = qs.SetCond(cond.And("id", pg.Id))
	}
	if pg.ParentId > 0 {
		qs = qs.SetCond(cond.And("parent_id", pg.ParentId))
	}
	if pg.Type != "" {
		qs = qs.SetCond(cond.And("type", pg.Type))

	}
	if pg.Code != "" && pg.Type != "" {
		qs = qs.SetCond(cond.And("code", pg.Code).And("type", pg.Type))
	}
	num, err = qs.Count()
	return num, err
}

func QueryDictFilter(pg PageQueryDict) string {
	var buffer bytes.Buffer
	buffer.WriteString(" 1=1")
	if pg.Id > 0 {
		buffer.WriteString(" and id=" + fmt.Sprintf("%d", pg.Id))
	}
	if pg.ParentId > 0 {
		buffer.WriteString(" and parent_id=" + fmt.Sprintf("%d", pg.ParentId))
	}
	if pg.Type != "" {
		buffer.WriteString(" and type=" + fmt.Sprintf("%d", pg.Type))
	}
	if pg.Code != "" {
		buffer.WriteString(" and code=" + fmt.Sprintf("%d", pg.Code))
	}
	return buffer.String()
}

func QueryDictFilterByTemplate(pg PageQueryDict) (string, error) {

	sql := ` 1=1
	  {{if .Id  $le 0}}
	    and id = {{.Id}}
	  {{end}}

    `
	var buffer bytes.Buffer
	tmpl, err := template.New("sql").Parse(sql)
	err = tmpl.Execute(&buffer, pg)
	return buffer.String(), err
}
