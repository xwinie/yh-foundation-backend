package auth

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SysUserRole struct {
	Id           int64
	RoleId       int64
	UserId       int64
	DeleteStatus int8      `orm:"default(0)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
	Locked       int8      `orm:"default(0)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysUserRole))
}

func CreateUserRole(p *SysUserRole) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(p)
	return id, err
}

func CreateMultiUserRole(p []SysUserRole) (int64, error) {
	o := orm.NewOrm()
	id, err := o.InsertMulti(len(p), p)
	return id, err
}

func DeleteUserRole(userId int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sys_user_role").Filter("user_id",userId).Delete()
	return num, err

}

func FindRoleIdByUserId(userId int64) ([]int64, int64, error) {
	o := orm.NewOrm()
	var roleIds []int64
	num, err := o.Raw("SELECT role_id FROM sys_user_role WHERE user_id =?", userId).QueryRows(&roleIds)
	return roleIds, num, err

}

func FindRoleByUserId(userId int64) (roles []SysRole, num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("SELECT r.* FROM sys_role r "+
		" inner join sys_user_role  ur on ur.role_id=r.id "+
		" WHERE ur.user_id =?", userId).QueryRows(&roles)
	return roles, num, err
}
