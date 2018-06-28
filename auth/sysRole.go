package auth

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SysRole struct {
	Id           int64
	Code         string    `orm:"size(100);unique"`
	Name         string    `orm:"size(100)"`
	Description  string    `orm:"size(255);null"`
	DeleteStatus int8      `orm:"default(0)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
	Locked       int8      `orm:"default(0)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysRole))
}

func CreateRole(p *SysRole) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(p)
	return id, err
}

func CheckRoleCodeIsExist(code string) (bool, error) {
	o := orm.NewOrm()
	var role SysRole
	err := o.QueryTable("sys_role").Filter("code", code).One(&role)
	if &role.Id != nil && err == nil {
		return true, err
	} else {
		return false, err
	}
}

func DeleteRole(p *SysRole) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(p)
	return num, err

}

func UpdateRole(p map[string]interface{}, id int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_role").Filter("id", id).Update(p)
	return num, err

}

func ReadRoleByCode(code string) (*SysRole, error) {
	o := orm.NewOrm()
	var role SysRole
	err := o.QueryTable("sys_role").Filter("code", code).One(&role)
	return &role, err
}

func ReadRoleById(id int64) (*SysRole, error) {
	o := orm.NewOrm()
	var role SysRole
	err := o.QueryTable("sys_role").Filter("id", id).One(&role)
	return &role, err
}

func ReadRoleByPage(pageSize int, offset int) (roles []*SysRole, num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_role").Limit(pageSize, offset).All(&roles)
	return roles, num, err
}

func ReadRoleCountByPage() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_role").Count()
	return num, err
}
