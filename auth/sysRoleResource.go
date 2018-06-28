package auth

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
	"strings"
)

type SysRoleResource struct {
	Id           int64
	RoleId       int64
	ResourceId   int64
	DeleteStatus int8      `orm:"default(0)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
	Locked       int8      `orm:"default(0)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysRoleResource))
}

type FindRoleResource struct {
	Code       string
	ResourceId int64
	Action     string
	Method     string
	RoleId     int64
}

func FindResourceByMultiRole(roleIds []int64, resType int8) (resource []FindRoleResource, num int64, err error) {
	o := orm.NewOrm()
	var placeholders []string
	for i := 0; i < len(roleIds); i++ {
		placeholders = append(placeholders, "?")
	}

	ss := fmt.Sprintf("SELECT r.code ,r.id resourceId,r.action,r.method,rr.id roleId FROM sys_resource r "+
		" inner join sys_role_resource rr on rr.resource_id=r.id "+
		" WHERE rr.role_id in (%v) and r.res_type=%v  order by r.parent_id", strings.Join(placeholders, ","), resType)
	num, err = o.Raw(ss, roleIds).QueryRows(&resource)
	return resource, num, err
}

func FindResourceByRoleId(roleId int64) (resource []SysResource, num int64, err error) {
	o := orm.NewOrm()

	num, err = o.Raw("SELECT r.* FROM sys_resource r "+
		" inner join sys_role_resource rr on rr.resource_id=r.id "+
		" WHERE rr.role_id =? order by r.res_type", roleId).QueryRows(&resource)
	return resource, num, err
}

func CreateRoleResource(p *SysRoleResource) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(p)
	return id, err
}

func CreateMultiRoleResource(p []SysRoleResource) (int64, error) {
	o := orm.NewOrm()
	id, err := o.InsertMulti(len(p), p)
	return id, err
}

func DeleteRoleResource(roleId int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_role_resource").Filter("role_id",roleId).Delete()
	return num, err

}
